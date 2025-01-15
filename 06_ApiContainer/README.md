# 事前準備
下記のGitHubディレクトリから関連資材がダウンロード（`git clone`）されていること<br>
https://github.com/CloudTechOrg/course-docker.git

以下のように必要なファイルが構成されていること
```
course-container/
└── 05_ApiContainer/
    ├── index.html
    └── Dockerfile
```

cdコマンドにより、カレントディレクトリを`05_ApiContainer`に移動する
```
cd 05_ApiContainer
```

## ハンズオン手順

## 1. イメージの作成
下記の`docker image build`コマンドにより、Dockerイメージを生成する
```
docker image build -t api-image .
```

下記コマンドにより、`api-image`のイメージが出力されることを確認する
```
docker image ls
```

## 2. Containerの起動
下記の`docker container run`コマンドにより、Containerを起動する
```
docker container run -p 8080:8080 \
  --name api-container \
  -e DB_USERNAME=root \
  -e DB_PASSWORD=cloudtech \
  -e DB_SERVERNAME=db-container \
  api-image
```

下記のコマンドにより`api-container`が実行中（STATUSが`Up xxxx`であること）を確認する
```
docker ps
```

## 3. 動作確認
ブラウザにて下記のURLを実行し、CURLコマンドを使用してAPIの動作確認を行う

下記のコマンドにより「API接続テストが成功しました」のメッセージが表示されることを確認する
```
curl http://localhost:8080
```

続けて下記コマンドにて、「データベース接続テストが成功しました」と表示された上、test_tableの件数が表示されることを確認

```
curl http://localhost:8080/test
```

## 4. 停止および削除
以下の`docker stop`コマンドで、`api-container`を停止する
```
docker stop api-container
```
