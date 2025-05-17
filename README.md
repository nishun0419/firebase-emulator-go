# Go + Firebase Emulator with Docker

このリポジトリは、Docker 上で Firebase Emulator を立ち上げて、Go (Golang) から Firestore にデータを読み書きする最小構成のサンプルです。

## 📦 構成

```
.
├── backend/              # Goアプリケーション（Firestoreに接続）
│   ├── cmd/main.go       # テストデータをFirestoreに追加するメインスクリプト
│   ├── go.mod, go.sum
│   └── Dockerfile
├── firebase/             # Firebase Emulator用設定
│   ├── .firebaserc
│   ├── firebase.json
│   └── Dockerfile
├── docker-compose.yml
└── README.md
```

## 🚀 起動方法

```bash
docker compose up -d
```

Emulator UI が立ち上がったら、ブラウザで以下を開く：

```
http://localhost:4000
```

## 🔨 Firestore にテストデータを追加

```bash
docker compose exec backend go run cmd/main.go
```

追加されたデータは以下から確認可能：

```
http://localhost:4000/firestore
```

## 🧵 補足

- `FIRESTORE_EMULATOR_HOST` を設定することで、実プロジェクト不要でFirestoreに接続できます
- Firebase CLIやJavaをローカルにインストールせずに開発できます

## 📝 参考リンク

- [Firebase Emulator Suite](https://firebase.google.com/docs/emulator-suite)
- [Firestore Go SDK](https://pkg.go.dev/cloud.google.com/go/firestore)

