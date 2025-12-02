# Copilot Instructions

このリポジトリは、React + TypeScript フロントエンドと、  
Go + chi バックエンドによる Todo アプリの学習用プロジェクト。

Copilot は以下のルールに従ってコードを提案すること。

---

## General

- シンプル、小さいコード実装を優先し、過剰な設計をしない
- 現時点のベストプラクティスに沿った、モダン、クリーン、ロバスト、シンプルなコードを書く
- OCP原則をはじめとした、SOLID原則を意識し、クリーンな設計にする。ただし、過剰にしない。
- 小さな変更を優先する（既存テストは必ず維持する）
- 過剰な機能追加や不要な抽象化はしない
- 型定義・API仕様・テストをセットで整合させる
- 新しい依存ライブラリを増やしすぎない

---

## Frontend (React + TypeScript)

- API 型定義を `Todo` として統一する

```ts
export interface Todo {
  id: number;
  text: string;
  done: boolean;
}
````

- API 呼び出しは `src/api/` に集約する
  例: `getTodos / createTodo / updateTodo / deleteTodo`
- UI はシンプルな状態管理でよい（まずはローカル state）
- 未実装のフィールドを勝手に追加しない
  例: `createdAt` や `updatedAt` を自動提案しない
- CORS 設定に依存するため、`Origin: http://localhost:3000` 前提で通信

---

## Backend (Go + chi)

### 基本ルール

- ルータ構成は `setupRouter()` に集約し、テストでも共通利用
- JSON の Decode 時は `DisallowUnknownFields()` を使用
- スライス操作と ID 採番は `mu.Lock()`〜`defer mu.Unlock()` で排他する
- 既存の API 仕様を尊重し、破壊的変更は避ける

### エンドポイント

| Method | Path | 説明 |
| - | - | - |
| GET | /todos | 一覧取得 |
| POST | /todos | 新規追加 |
| PATCH  | /todos/{id} | 部分更新（title/done のみ） |
| DELETE | /todos/{id} | 削除 |

- PATCH の仕様
  - 不正 JSON → 400
  - 未知フィールド → 400
  - 空パッチ `{}` → 400
  - 存在しない ID → 404
- CORS は `go-chi/cors` によりミドルウェアで一元管理
  → 各ハンドラで個別設定しない

### テスト方針

- 正常系 + 異常系を明示
- CORS のプリフライトは 200/204 のどちらも許容
- 競合テストあり
  → ID がユニークであることを検証

---

以上の方針に従い、既存機能と整合するコード提案のみを行うこと。
