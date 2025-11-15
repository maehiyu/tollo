# Tollo Web App

TolloのWebクライアントアプリケーション。React + TypeScript + Viteで構築されています。

## 技術スタック

- **React**: UI フレームワーク
- **TypeScript**: 型安全な開発
- **Vite**: 高速なビルドツール
- **Kotlin Multiplatform (KMP)**: 共通ロジック層（`client/shared`）
  - API呼び出し（Apollo Kotlin）
  - バリデーション
  - ビジネスロジック

## フォルダ構成

```
src/
├── features/                    # 機能ごとのモジュール
│   ├── user/                    # User機能
│   │   ├── pages/              # ページコンポーネント
│   │   │   ├── UserCreatePage.tsx
│   │   │   ├── UserDetailPage.tsx
│   │   │   └── UserListPage.tsx
│   │   ├── components/         # UIコンポーネント
│   │   │   ├── UserForm.tsx
│   │   │   ├── UserProfile.tsx
│   │   │   ├── UserSearchForm.tsx
│   │   │   ├── GeneralProfileForm.tsx
│   │   │   └── ProfessionalProfileForm.tsx
│   │   ├── hooks/              # カスタムフック（KMP serviceをラップ）
│   │   │   └── useUser.ts
│   │   └── types/              # UI層の型定義
│   │       └── index.ts
│   │
│   └── chat/                    # Chat機能
│       ├── pages/
│       │   └── ChatRoomPage.tsx
│       ├── components/
│       │   ├── ChatMessageList.tsx
│       │   ├── ChatInput.tsx
│       │   └── ChatRoom.tsx
│       ├── hooks/
│       │   └── useChat.ts
│       └── types/
│           └── index.ts
│
├── components/common/           # 全機能共通のUIコンポーネント
│   ├── Button.tsx
│   ├── Input.tsx
│   ├── ErrorMessage.tsx
│   └── LoadingSpinner.tsx
│
├── routes/                      # ルーティング設定
│   └── index.tsx
│
├── App.tsx                      # アプリケーションルート
└── main.tsx                     # エントリーポイント
```

## 設計方針

### 責任の分離

**KMP側（`client/shared`）の責任:**
- GraphQL API呼び出し
- データのバリデーション
- ビジネスロジック（UserService、ChatServiceなど）
- データモデル・型定義

**webApp側（React）の責任:**
- UIコンポーネント（プレゼンテーション）
- ルーティング
- React固有の状態管理
- KMPサービスの呼び出しとReact hooksへのラップ

### features/ 構造

各機能（user, chat, questionなど）は独立したモジュールとして管理：
- **pages/**: ルーティングに対応するページコンポーネント
- **components/**: その機能専用のUIコンポーネント
- **hooks/**: KMPのサービスを呼び出すカスタムフック
- **types/**: UI層で必要な型定義（KMP生成以外）

この構造により：
- 機能の追加・削除が容易
- ドメインロジックが1箇所にまとまる
- KMP側の構造（`domain/service/`）と対応
- チーム分担がしやすい

## 開発

### インストール
```bash
npm install
```

### 開発サーバー起動
```bash
npm run dev
```

### ビルド
```bash
npm run build
```

### プレビュー
```bash
npm run preview
```

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type-aware lint rules:

```js
export default defineConfig([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      // Other configs...

      // Remove tseslint.configs.recommended and replace with this
      tseslint.configs.recommendedTypeChecked,
      // Alternatively, use this for stricter rules
      tseslint.configs.strictTypeChecked,
      // Optionally, add this for stylistic rules
      tseslint.configs.stylisticTypeChecked,

      // Other configs...
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
      // other options...
    },
  },
])
```

You can also install [eslint-plugin-react-x](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-x) and [eslint-plugin-react-dom](https://github.com/Rel1cx/eslint-react/tree/main/packages/plugins/eslint-plugin-react-dom) for React-specific lint rules:

```js
// eslint.config.js
import reactX from 'eslint-plugin-react-x'
import reactDom from 'eslint-plugin-react-dom'

export default defineConfig([
  globalIgnores(['dist']),
  {
    files: ['**/*.{ts,tsx}'],
    extends: [
      // Other configs...
      // Enable lint rules for React
      reactX.configs['recommended-typescript'],
      // Enable lint rules for React DOM
      reactDom.configs.recommended,
    ],
    languageOptions: {
      parserOptions: {
        project: ['./tsconfig.node.json', './tsconfig.app.json'],
        tsconfigRootDir: import.meta.dirname,
      },
      // other options...
    },
  },
])
```
