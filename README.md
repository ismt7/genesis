# genesis

A CLI tool built with [Cobra](https://github.com/spf13/cobra).

## 前提条件

- Go 1.21 以上

## インストール

### `go install` を使う方法（推奨）

```bash
go install github.com/ismt7/genesis@latest
```

`$GOPATH/bin`（デフォルトは `~/go/bin`）に `genesis` バイナリが配置されます。  
パスが通っていない場合は以下を `.zshrc` / `.bashrc` に追加してください。

```bash
export PATH="$HOME/go/bin:$PATH"
```

### ソースからビルドする方法

```bash
git clone https://github.com/ismt7/genesis.git
cd genesis
go build -o genesis .
```

ビルドされた `genesis` バイナリをパスの通ったディレクトリに移動してください。

```bash
mv genesis /usr/local/bin/
```

## アップデート

### `go install` でインストールした場合

最新バージョンに更新するには、インストール時と同じコマンドを再実行してください。

```bash
go install github.com/ismt7/genesis@latest
```

> **注意:** `go install` ではビルド時の `-ldflags` が渡されないため、`genesis version` を実行すると `genesis vdev (commit: none, built: unknown)` と表示されます。正確なバージョン情報が必要な場合は、[GitHubリリースページ](https://github.com/ismt7/genesis/releases)から該当バージョンのバイナリを直接ダウンロードしてください。

### ソースからビルドした場合

リポジトリを最新の状態に更新してからビルドしてください。

```bash
cd genesis
git pull origin main
go build -o genesis .
mv genesis /usr/local/bin/
```

## 使い方

```bash
# ヘルプを表示
genesis --help

# バージョンを確認
genesis version
```

## 動作確認

ソースから直接実行して動作を確認できます。

```bash
# ヘルプを表示
go run . --help

# バージョンを確認
go run . version
```

ビルド後に確認する場合は以下の通りです。

```bash
go build -o genesis .

# ヘルプを表示
./genesis --help

# バージョンを確認
./genesis version
```

## ライセンス

See [LICENSE](LICENSE).