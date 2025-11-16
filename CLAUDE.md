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

### サービス追加時の手順

1. **Protocol Buffers定義**
   ```bash
   # protos/servicename/v1/servicename.proto を作成
   buf generate
   ```

2. **サービス実装**
   ```bash
   # services/servicename/ に実装
   # - main.go
   # - handler.go
   ```

3. **Gateway統合**
   ```bash
   # internal/gateway/graph/schema.graphqls にスキーマ追加
   # internal/gateway/graph/resolver.go にリゾルバ追加
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

## 参考リンク

- [要求仕様](docs/01_requirement.md)
- [アーキテクチャ詳細](docs/02_architecture.md)
- [WebApp README](client/webApp/README.md)