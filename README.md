# tolla

## プロジェクト概要
本プロジェクトは、Q&Aプラットフォームをマイクロサービスアーキテクチャで構築しています。

## ドキュメント
- **要求仕様**: [docs/01_requirement.md](docs/01_requirement.md)
- **アーキテクチャ**: [docs/02_architecture.md](docs/02_architecture.md)
- **API仕様**: [docs/openapi.yaml](docs/openapi.yaml)

## 開発環境
- Go: 1.24.4

## Protobufコードの生成
`protos`ディレクトリ内の`.proto`ファイルからGoのコードを生成するには、以下のコマンドを実行します。

```bash
make
```
このコマンドにより、`gen/go`ディレクトリにGoのProtobufコードが生成されます。

## 補足
現時点では、`go install`で直接インストールする実行可能バイナリは提供されていません。
