# tollo

## プロジェクト概要
本プロジェクトは、Q&Aプラットフォームをマイクロサービスアーキテクチャで構築しています。
クライアントサイドではKotlin Multiplatform (KMP) を採用し、サービス定義にはApollo Kotlinを使用しています。共通ロジックはcommonMainで、プラットフォーム固有の実装はjsMainで管理されます。

## ドキュメント
- **要求仕様**: [docs/01_requirement.md](docs/01_requirement.md)
- **アーキテクチャ**: [docs/02_architecture.md](docs/02_architecture.md)

## 開発環境
- Go: 1.24.4
- [Buf CLI](https://buf.build/docs/installation)

## Protobufコードの生成
`protos`ディレクトリ内の`.proto`ファイルからGoのコードを生成するには、[Buf CLI](https://buf.build/docs/installation)が必要です。

以下のコマンドを実行します。

```bash
buf generate
```

このコマンドにより、`buf.gen.yaml`の設定に従って`gen/go`ディレクトリにコードが生成されます。

## ビルドとテスト
本プロジェクトのクライアントサイドはGradleで管理されています。

### ビルド
すべてのサブプロジェクトをビルドします。
```bash
./gradlew build
```

### クリーン
ビルドディレクトリを削除します。
```bash
./gradlew clean
```

### テストとチェック
`client:shared`モジュールのテストと静的解析を実行します。
```bash
./gradlew :client:shared:check
```

## 補足
- 現時点では、`go install`で直接インストールする実行可能バイナリは提供されていません。
- コードの修正はせず、修正後のコードを見せるだけにしてください。
