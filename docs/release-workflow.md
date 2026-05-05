# リリース・バージョン管理 運用手順

## 概要

genesis では **GoReleaser** を使ったリリースフローを採用しています。  
バイナリにはビルド時に `version` / `commit` / `date` が埋め込まれ、`genesis version` で確認できます。

---

## バージョン番号の仕組み

| 場面 | version の値 |
|------|-------------|
| `go run` / `go build` (タグなし) | `dev` |
| `make build` (タグあり) | `git describe` の出力 (例: `1.2.3`) |
| GoReleaser でビルド | タグ名から `v` を除いた値 (例: `1.2.3`) |

バージョンは `-ldflags` で次の変数に注入されます。

```go
// cmd/version.go
var (
    version = "dev"   // -X github.com/ismt7/genesis/cmd.version=...
    commit  = "none"  // -X github.com/ismt7/genesis/cmd.commit=...
    date    = "unknown" // -X github.com/ismt7/genesis/cmd.date=...
)
```

---

## ローカルでのビルド・確認

```bash
# バイナリをビルド
make build

# バージョン情報を表示
make version
# → genesis v1.2.3 (commit: abc1234, built: 2026-05-05T04:41:00Z)

# ビルド成果物を削除
make clean
```

> **注意**: `make build` はカレントディレクトリの最新 git タグを参照します。  
> タグが存在しない場合は `dev` になります。

---

## リリース手順

### 1. 変更をコミットする

```bash
git add .
git commit -m "feat: <変更内容>"
```

### 2. git tag を打つ

バージョンは **セマンティックバージョニング** (`vMAJOR.MINOR.PATCH`) に従います。

```bash
# 例: v1.3.0 をリリース
git tag v1.3.0

# タグをリモートに push
git push origin v1.3.0
```

> `git push origin main` だけではリリースは行われません。**タグの push が必須**です。

### 3. GoReleaser でリリースする

GitHub Actions が設定されている場合はタグ push で自動実行されます。  
手動で実行する場合:

```bash
# 本番リリース (GitHub へ成果物をアップロード)
goreleaser release --clean

# ローカルでのテストビルド (GitHub へはアップロードしない)
goreleaser release --snapshot --clean
```

GoReleaser は `.goreleaser.yaml` の設定に従い、次のターゲット向けバイナリを生成します:

| OS | アーキテクチャ |
|----|--------------|
| Linux | amd64 / arm64 |
| macOS (darwin) | amd64 / arm64 |
| Windows | amd64 / arm64 |

成果物のファイル名形式: `genesis_<version>_<os>_<arch>.tar.gz`  
(Windows のみ `.zip`)

### 4. リリース結果の確認

```bash
# 最新タグを確認
git tag --sort=-version:refname | head -5

# リモートの最新リリースを確認 (gh CLI)
gh release list --limit 5

# リリースの詳細を確認
gh release view v1.3.0
```

---

## タグの修正・削除

```bash
# タグをローカルで削除
git tag -d v1.3.0

# タグをリモートから削除
git push origin --delete v1.3.0

# タグを付け直してリモートに push
git tag v1.3.0
git push origin v1.3.0
```

> リリース済みのタグを削除・付け直す場合は、GitHub 上のリリースも手動で削除してください。

---

## 自動アップデート (`genesis update`)

`genesis update` コマンドは GitHub Releases API から最新バージョンを取得し、バイナリを自動置換します。

- **dev ビルド** (`version == "dev"`) の場合はスキップされます。
- GoReleaser でリリースされた正規バイナリでのみ動作します。

```bash
genesis update
# → checking for updates...
# → updating from v1.2.0 to v1.3.0...
# → updated to v1.3.0 successfully
```

---

## Changelog の自動生成

GoReleaser は git のコミットメッセージから Changelog を自動生成します。  
以下のプレフィックスを持つコミットは Changelog から**除外**されます:

| プレフィックス | 用途 |
|--------------|------|
| `docs:` | ドキュメント変更 |
| `test:` | テスト追加・修正 |
| `chore:` | 雑務・依存更新など |

Changelog に載せたい変更は `feat:` / `fix:` / `refactor:` などを使ってください。