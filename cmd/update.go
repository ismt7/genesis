package cmd

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

const repoAPI = "https://api.github.com/repos/ismt7/genesis/releases/latest"

type githubRelease struct {
	TagName string        `json:"tag_name"`
	Assets  []githubAsset `json:"assets"`
}

type githubAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update genesis to the latest version",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version == "dev" {
			fmt.Println("skipping update: running dev build")
			return nil
		}

		fmt.Println("checking for updates...")

		release, err := fetchLatestRelease()
		if err != nil {
			return fmt.Errorf("failed to fetch latest release: %w", err)
		}

		latest := strings.TrimPrefix(release.TagName, "v")
		current := strings.TrimPrefix(version, "v")

		if latest == current {
			fmt.Printf("already up to date (v%s)\n", current)
			return nil
		}

		fmt.Printf("updating from v%s to v%s...\n", current, latest)

		assetName := fmt.Sprintf("genesis_%s_%s_%s.tar.gz", latest, runtime.GOOS, runtime.GOARCH)
		var downloadURL string
		for _, a := range release.Assets {
			if a.Name == assetName {
				downloadURL = a.BrowserDownloadURL
				break
			}
		}
		if downloadURL == "" {
			return fmt.Errorf("no binary found for %s/%s in release %s", runtime.GOOS, runtime.GOARCH, release.TagName)
		}

		exePath, err := os.Executable()
		if err != nil {
			return fmt.Errorf("could not determine executable path: %w", err)
		}
		exePath, err = filepath.EvalSymlinks(exePath)
		if err != nil {
			return fmt.Errorf("could not resolve symlinks: %w", err)
		}

		newBinary, err := downloadBinary(downloadURL)
		if err != nil {
			return fmt.Errorf("download failed: %w", err)
		}
		defer os.Remove(newBinary)

		if err := replaceBinary(exePath, newBinary); err != nil {
			return fmt.Errorf("failed to replace binary: %w", err)
		}

		fmt.Printf("updated to v%s successfully\n", latest)
		return nil
	},
}

func fetchLatestRelease() (*githubRelease, error) {
	req, err := http.NewRequest(http.MethodGet, repoAPI, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var release githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}
	return &release, nil
}

func downloadBinary(url string) (string, error) {
	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %s", resp.Status)
	}

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return "", err
	}
	defer gz.Close()

	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if hdr.Name == "genesis" || hdr.Name == "genesis.exe" {
			tmp, err := os.CreateTemp("", "genesis-update-*")
			if err != nil {
				return "", err
			}
			if _, err := io.Copy(tmp, tr); err != nil { //nolint:gosec
				tmp.Close()
				os.Remove(tmp.Name())
				return "", err
			}
			tmp.Close()
			if err := os.Chmod(tmp.Name(), 0755); err != nil {
				os.Remove(tmp.Name())
				return "", err
			}
			return tmp.Name(), nil
		}
	}
	return "", fmt.Errorf("genesis binary not found in archive")
}

func replaceBinary(dest, src string) error {
	// rename is atomic on same filesystem; use a backup for cross-fs safety
	backup := dest + ".bak"
	if err := os.Rename(dest, backup); err != nil {
		return err
	}
	if err := os.Rename(src, dest); err != nil {
		// restore backup
		_ = os.Rename(backup, dest)
		return err
	}
	os.Remove(backup)
	return nil
}

func init() {
	rootCmd.AddCommand(updateCmd)
}