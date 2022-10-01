# go_todo_app

[go_todo_app](https://github.com/budougumi0617/go_todo_app)を実装してみた．

## API

| HTTPメソッド | パス         | 概要                       |
|----------|------------|--------------------------|
| POST     | `/regiser` | 新しいユーザーを登録する             |
| POST     | `/login`   | 登録済みユーザー情報でアクセストークンを取得する |
| POST     | `/tasks`   | アクセストークンを使ってタスクを登録する     |
| GET      | `/tasks`   | アクセストークンを使ってタスクを一覧する     |

## 起動方法

Dockerイメージ作成．

```bash
make build-local
```

サービス起動．

```bash
make up
```

MySQLマイグレーション実行．

```bash
make migrate
```

以降のコマンドは，`test/run.sh`を参照．
