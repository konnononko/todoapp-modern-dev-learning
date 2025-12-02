# Todo App (Go + React)

個人学習用の Todo アプリです。  
フロントエンドは **React + TypeScript**、  
バックエンドは **Go + chi** により実装しています。

## Features

- Todo の追加・更新（タイトル text／完了状態 done）
- Todo の削除
- 完了チェック
- REST API によるフロント連携
- CORS 対応
- 並行処理に対する排他制御
- API の単体テストあり（Go）

## Quick Start

Backend:

```bash
cd backend
go mod tidy
go run .
```

Frontend:

```bash
cd frontend
npm install
npm run dev
# http://localhost:3000
```

## Project Structure

```
backend/            # Go API サーバー
  main.go           # サーバーエントリ（setupRouter, handlers）
  main_test.go      # バックエンドの単体テスト

frontend/           # React + TypeScript SPA
  src/
    api/             # API 呼び出し実装およびモック (todoApi.ts, todoApi.mock.ts)
      todoApi.ts
      todoApi.mock.ts
    components/      # React コンポーネント
      TodoInput.tsx
      TodoList.tsx
      TodoItem.tsx
    types.ts         # Todo 型定義など
    App.tsx
    App.test.tsx     # フロントエンドのユニットテスト
```

## Backend

### API 仕様

| Method | Path | Description |
| - | - | - |
| GET | /todos | 一覧取得 |
| POST | /todos | 新規追加 |
| PATCH | /todos/{id} | 部分更新（更新対象フィールド:text, done） |
| DELETE | /todos/{id} | 削除 |

- PATCH の仕様
  - 更新対象フィールドは `text`（Todo本文）と `done`（完了フラグ）のみ
  - 不正 JSON → 400
  - 未知フィールド → 400
  - 空パッチ `{}` → 400
  - 存在しない ID → 404

### CORS

開発オリジン `http://localhost:3000` を許可（go-chi/cors ミドルウェアで一元管理）。

### Run Backend Tests

```bash
cd backend
go test ./...
```

## Frontend

### Components

- App
  - TodoInput（入力フォーム）
  - TodoList（リスト表示）
    - TodoItem（1つ1つのTodo）

## Todo 型（参照）

```ts
interface Todo { id: number; text: string; done: boolean }
```

### Run Frontend Tests

```bash
cd frontend
npm test        # run unit tests (vitest)
npm run test:ui # interactive test UI
npm run typecheck # TypeScript 型チェック (no emit)
```

## License

MIT

## Author

Personal Learning Project
