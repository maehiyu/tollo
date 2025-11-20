# Claude Code 作業ガイド

このドキュメントは、Claude Codeでの作業時に参照する、プロジェクト全体の設計方針とアーキテクチャのガイドです。

## プロジェクト概要

Tolloは、Q&Aプラットフォームをマイクロサービスアーキテクチャで構築するプロジェクトです。

### 技術スタック

**バックエンド:**
- Go 1.24.4+
- gRPC (サービス間通信)
- GraphQL (クライアント向けAPI)
- Protocol Buffers

**フロントエンド:**
- React + TypeScript + Vite
- Kotlin Multiplatform (KMP) - 共通ロジック層
- Apollo Kotlin (GraphQL client)

## アーキテクチャ概要

```
┌─────────────────────────────────────────────────┐
│              Client (React + KMP)               │
│  ┌───────────────────────────────────────────┐  │
│  │  webApp (React/TypeScript)                │  │
│  │  - UI/プレゼンテーション層                  │  │
│  │  - features/ 構造                          │  │
│  └───────────────────────────────────────────┘  │
│  ┌───────────────────────────────────────────┐  │
│  │  shared (KMP)                             │  │
│  │  - ビジネスロジック                         │  │
│  │  - API呼び出し (Apollo Kotlin)             │  │
│  │  - バリデーション                           │  │
│  └───────────────────────────────────────────┘  │
└─────────────────────────────────────────────────┘
                        │
                        ↓ GraphQL
┌─────────────────────────────────────────────────┐
│              Gateway (GraphQL)                  │
│  - GraphQL API                                  │
│  - gRPC クライアント                            │
└─────────────────────────────────────────────────┘
                        │
                        ↓ gRPC
┌──────────────────┬──────────────────┬───────────┐
│  UserService     │  ChatService     │  ...      │
│  (gRPC Server)   │  (gRPC Server)   │           │
└──────────────────┴──────────────────┴───────────┘
```

## フォルダ構成

### ルート構造

```
tollo/
├── cmd/                    # 実行可能ファイル
│   └── gateway/           # GraphQL Gateway
├── internal/              # 内部パッケージ
│   ├── gateway/          # Gateway実装
│   └── auth/             # 認証ロジック
├── services/              # マイクロサービス
│   ├── userservice/      # ユーザーサービス
│   └── chatservice/      # チャットサービス
├── protos/                # Protocol Buffers定義
├── gen/                   # 生成コード
│   └── go/               # Go生成コード
├── client/                # クライアント
│   ├── shared/           # KMP共通ロジック
│   └── webApp/           # Reactアプリ
└── docs/                  # ドキュメント
```

### client/webApp の構造 (features/)

```
client/webApp/src/
├── features/              # 機能ごとのモジュール
│   ├── user/
│   │   ├── pages/        # ページコンポーネント
│   │   ├── components/   # UIコンポーネント
│   │   ├── hooks/        # カスタムフック (KMP serviceラップ)
│   │   └── types/        # UI層の型定義
│   └── chat/
│       ├── pages/
│       ├── components/
│       ├── hooks/
│       └── types/
├── components/common/     # 全機能共通UI
├── routes/                # ルーティング設定
├── App.tsx
└── main.tsx
```

## 責任の分離

### KMP (client/shared) の責任

**KMP側で扱うもの:**
- ✅ GraphQL API呼び出し (Apollo Kotlin)
- ✅ ビジネスロジック (UserService, ChatService など)
- ✅ データのバリデーション
- ✅ データモデル・型定義
- ✅ リポジトリパターンの実装

**KMP構造:**
```
client/shared/src/commonMain/
├── graphql/                      # GraphQLクエリ/ミューテーション
├── kotlin/.../client/shared/
│   ├── data/
│   │   └── repository/          # リポジトリ実装
│   └── domain/
│       ├── repository/          # リポジトリインターフェース
│       └── service/             # ビジネスロジック
└── jsMain/                       # JS固有実装
```

### React (client/webApp) の責任

**React側で扱うもの:**
- ✅ UIコンポーネント (プレゼンテーション)
- ✅ ルーティング
- ✅ React固有の状態管理
- ✅ KMPサービスの呼び出しとhooksへのラップ

**典型的なフロー:**
```typescript
// hooks/useUser.ts - KMPサービスをReact hooksでラップ
import { createUser } from 'shared'; // KMP生成

export const useUser = () => {
  const [loading, setLoading] = useState(false);
  const create = async (params) => {
    setLoading(true);
    try {
      return await createUser(...params);
    } finally {
      setLoading(false);
    }
  };
  return { create, loading };
};

// pages/UserCreatePage.tsx - ページはhooksを使う
const UserCreatePage = () => {
  const { create, loading } = useUser();
  return <UserForm onSubmit={create} loading={loading} />;
};
```

## バックエンドサービス

### サービス間連携とデータフロー

本プロジェクトでは、各サービスが明確な責務を持ち、イベント駆動で連携します。

#### QuestionService (質問管理・ビジネスロジック)
**責務:**
- 質問の永続化・管理
- AI回答の要否判定
- プロマッチングロジック
- Question ↔ Chat の紐付け管理
- リワード計算のトリガー

**データモデル:**
```go
type Question struct {
    ID              string
    UserID          string
    Content         string
    RequireAI       bool         // AI経由するか
    AIAnswerID      *string      // AI回答への参照
    ChatID          *string      // Chat への参照
    Status          QuestionStatus
    ProfessionalID  *string      // マッチしたプロ
}

type QuestionStatus int
const (
    Open      QuestionStatus = iota  // AI回答済み、プロ待ち
    Matched                           // プロとマッチング済み
    Answered                          // 回答完了
    Closed                            // クローズ
)
```

#### ChatService (データストア層 - Thin Layer)
**責務:**
- メッセージの保存と取得のみ
- Chat の CRUD
- Message の CRUD
- ユーザーのチャット一覧取得
- ❌ ビジネスロジックは持たない

**データモデル:**
```go
type Chat struct {
    ID              string
    QuestionID      *string      // Question から来た場合の参照
    GeneralUserID   string
    ProfessionalID  *string      // マッチング前は null
}

type Message struct {
    ID          string
    ChatID      string
    SenderID    string
    Type        MessageType  // STANDARD, AI_ANSWER, PRO_ANSWER, etc
    Content     string
}
```

#### AIAnswerService (RAG - AI回答生成)
**責務:**
- RAGによるAI回答生成
- プロの回答からナレッジベース構築
- Vector DBへのインデックス化

**データソース:**
- ChatService内のプロの回答 (type=PRO_ANSWER) をRAGのソースとする
- イベント駆動でリアルタイムに学習

**イベントフロー:**
```
ChatService: プロが回答
    ↓ イベント発行
MessageCreated (type=PRO_ANSWER)
    ↓ Pub/Sub
AIAnswerService: イベント購読
    ↓
Vector DBにインデックス化
    ↓
次回のAI回答生成時に活用
```

#### ユーザーフロー全体像

```
ユーザーが質問入力
    ↓
┌────────────────────────────────────────┐
│ ユーザーの選択                          │
│ 1. AIに聞く → AI回答後、プロ検討       │
│ 2. 最初からプロに聞く                   │
└────────────────────────────────────────┘
    ↓
QuestionService で Question 作成
    ↓
    ├─ AIに聞く場合
    │    ↓ AIAnswerService (RAG)
    │    AI回答 + 精度スコア
    │    ↓ ユーザー判断
    │    プロにも聞く? → Yes
    │         ↓
    └─ 最初からプロの場合
         ↓
    プロマッチング (QuestionService)
         ↓
    Chat作成 (ChatService)
         ↓
    Chat一覧に表示 (AI回答もChat UI、プロ回答もChat UI)
         ↓
    プロが回答 → MessageCreated イベント
         ↓
    AIAnswerService が学習 (Vector DB)
```

### サービス追加時の手順

1. **Protocol Buffers定義**
   ```bash
   # protos/servicename/v1/servicename.proto を作成
   buf generate
   ```

2. **サービス実装**
   ```bash
   # internal/servicename/ に実装
   # - server.go (gRPC Server)
   # - usecase.go (ビジネスロジック)
   # - domain/ (ドメインモデル)
   ```

3. **Gateway統合**
   ```bash
   # internal/gateway/graph/schema.graphqls にスキーマ追加
   # internal/gateway/graph/resolver.go にリゾルバ追加
   ```

4. **イベント駆動連携 (必要な場合)**
   ```go
   // イベント発行
   eventBus.Publish(EventName{...})

   // イベント購読
   eventBus.Subscribe("EventName", handlerFunc)
   ```

## 開発ワークフロー

### フロントエンド開発

```bash
# KMP側のビルド（GraphQLコード生成含む）
cd client
./gradlew :shared:build

# Reactアプリ起動
cd webApp
npm install
npm run dev
```

### バックエンド開発

```bash
# Protocol Buffersからコード生成
buf generate

# Gatewayの起動
go run cmd/gateway/main.go

# 各サービスの起動
go run services/userservice/main.go
go run services/chatservice/main.go
```

## features/ 構造の原則

### 新機能追加時

例: Question機能を追加する場合

1. **KMP側**: `client/shared/src/commonMain/kotlin/.../domain/service/QuestionService.kt` を作成
2. **React側**: `client/webApp/src/features/question/` ディレクトリを作成
   ```
   features/question/
   ├── pages/
   │   ├── QuestionListPage.tsx
   │   ├── QuestionDetailPage.tsx
   │   └── QuestionCreatePage.tsx
   ├── components/
   │   ├── QuestionCard.tsx
   │   └── QuestionForm.tsx
   ├── hooks/
   │   └── useQuestion.ts
   └── types/
       └── index.ts
   ```

3. **ルーティング**: `routes/index.tsx` に追加

### コンポーネント分割の方針

- **pages/**: ルーティング単位のページ全体
- **components/**: 再利用可能なUIパーツ（propsを受け取り表示のみ）
- **hooks/**: 状態管理とKMPサービスの呼び出し
- **types/**: UI層固有の型（FormState, UIのenumなど）

### 共通コンポーネント

全機能で使う場合は `components/common/` に配置：
- Button, Input, Select などの基本UI
- ErrorMessage, LoadingSpinner などのフィードバックUI
- Layout, Header, Footer などの構造UI

## Git ワークフロー

- **メインブランチ**: `main`
- **機能ブランチ**: `feature/<機能名>` (例: `feature/client/user-api`)
- コミット前に必ず `go mod tidy` と `./gradlew build` を実行

## トラブルシューティング

### IntelliJ IDEA で import エラー

```bash
# プロジェクトを閉じてから
rm -rf .idea
# IntelliJで再度開く
```

### KMP の JS 生成がうまくいかない

```bash
cd client
./gradlew clean
./gradlew :shared:build
```

### GraphQL スキーマ変更後

```bash
# KMP側で再生成
cd client
./gradlew :shared:build

# React側で反映
cd webApp
npm run dev  # ホットリロードで自動反映
```

## イベント駆動実装戦略

### Message Broker: NATS JetStream (推奨)

**選定理由:**
- Go製で軽量・高速
- gRPCマイクロサービスとの相性抜群
- セットアップが簡単
- 永続化サポート (JetStream)
- CNCF プロジェクト

**代替案:**
- RabbitMQ: 成熟、管理UIが優秀
- Kafka: 超高スループット、大規模向け (オーバーキル)
- Redis Streams: 既存Redis活用、シンプル

### 実装の段階的アプローチ

```
Phase 1: 同期的gRPC実装
  - 各サービスをgRPCで実装
  - 動作確認・テスト完了
  ↓
Phase 2: NATS導入
  - Docker Composeに追加
  - イベントバス抽象層を作成
  ↓
Phase 3: 段階的にイベント駆動化
  - QuestionCreated → AIAnswerService
  - MessageCreated (PRO_ANSWER) → AIAnswerService
  - AnswerCreated → RewardService
```

**NATS使用例:**
```go
// Publish
nc, _ := nats.Connect(nats.DefaultURL)
js, _ := nc.JetStream()
js.Publish("question.created", questionData)

// Subscribe
js.Subscribe("question.created", func(msg *nats.Msg) {
    // イベント処理
})
```

## 参考リンク

- [要求仕様](docs/01_requirement.md)
- [アーキテクチャ詳細](docs/02_architecture.md)
- [システム全体像 (Mermaid図)](docs/03_system_overview.md)
- [WebApp README](client/webApp/README.md)