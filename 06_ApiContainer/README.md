# 事前準備
任意のフォルダーに以下のファイルがダウンロードされていること
```
任意のフォルダ
├── main.go
├── go.mod
├── go.sum
└── Dockerfile
```

## ハンズオン手順

## 1. イメージの作成
`api-image`を作成する
```
docker image build -t api-image .
```

`api-image`が作成されていることを確認する
```
docker image ls
```

## 2. コンテナーの起動
`api-container`を起動する
```
docker container run -p 8080:8080 \
  --name api-container \
  -e DB_USERNAME=root \
  -e DB_PASSWORD=cloudtech \
  -e DB_SERVERNAME=db-container \
  api-image
```

`api-container`が実行中（STATUSが`Up xxxx`で）あることを確認する
```
docker image ls
```

## 3. 動作確認
下記の`curl`コマンドにより「API接続テストが成功しました」のメッセージが表示されることを確認する
```
curl http://localhost:8080
```