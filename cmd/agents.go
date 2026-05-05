package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const agentsTemplate = `# AGENTS.md

## プロジェクト概要
<!-- プロジェクトの目的・背景を記述してください -->

## アーキテクチャ
<!-- システムの全体構成を記述してください -->

## 開発ガイドライン

### コードスタイル
- リポジトリの既存のコード規約に従ってください

### テスト
- 新機能にはテストを作成してください
- 変更前にテストが通ることを確認してください

### コミットメッセージ
- 変更内容が明確に伝わるメッセージを書いてください

## 主要コマンド
<!-- 開発で使用する主要なコマンドを記述してください -->

## AIエージェントへの注意事項

英語で考えて日本語で回答してください。

<!-- その他、AIコーディングエージェント向けの特記事項・コンテキストを記述してください -->
`

var agentsCmd = &cobra.Command{
	Use:   "agents",
	Short: "Create an AGENTS.md file in the current directory",
	Long:  `Create an AGENTS.md template file in the current directory for AI coding agents.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		const filename = "AGENTS.md"

		if _, err := os.Stat(filename); err == nil {
			return fmt.Errorf("%s already exists", filename)
		}

		if err := os.WriteFile(filename, []byte(agentsTemplate), 0644); err != nil {
			return fmt.Errorf("failed to create %s: %w", filename, err)
		}

		fmt.Printf("Created %s\n", filename)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(agentsCmd)
}