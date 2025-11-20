# tollo

## プロジェクト概要
本プロジェクトは、Q&Aプラットフォームをマイクロサービスアーキテクチャで構築しています。
AIとプロフェッショナルの両方が回答できる仕組みで、RAGによる学習とリワードシステムを備えています。

## システム概要

### コア機能
1. **質問 → AI回答 → プロ回答** のフロー
   - ユーザーが質問を投稿
   - AI (RAG) が即座に回答 + 精度スコア表示
   - ユーザーがプロに聞くか選択
   - プロとマッチングしてチャット形式で対話

2. **RAG学習サイクル**
   - プロの回答がVector DBに自動インデックス化
   - AI回答の精度が継続的に向上

3. **リワードシステム**
   - プロの回答に対して需給と評価に基づくリワード
   - リアルタイムで報酬計算

### アーキテクチャ

```
Client (React + KMP)
    ↓ GraphQL
Gateway
    ↓ gRPC
┌──────────────┬─────────────┬──────────────┐
│ QuestionSvc  │  ChatSvc    │ AIAnswerSvc  │
│ (ロジック)   │ (データ)    │ (RAG)        │
└──────────────┴─────────────┴──────────────┘
                    ↓ Pub/Sub
            Message Broker (イベント駆動)
```

**サービスの責任分離:**
- **QuestionService**: 質問管理、マッチングロジック、Question↔Chat紐付け
- **ChatService**: メッセージ保存・取得のみ (Thin Layer)
- **AIAnswerService**: RAGによるAI回答生成、プロ回答から学習
- **RewardService**: リワード計算・支払い

## ドキュメント
- **要求仕様**: [docs/01_requirement.md](docs/01_requirement.md)
- **アーキテクチャ**: [docs/02_architecture.md](docs/02_architecture.md)
- **開発ガイド**: [CLAUDE.md](CLAUDE.md) - サービス間連携の詳細

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
