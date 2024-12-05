## ハンズオン手順

## 1. Docker Networkを作成
以下のコマンドで、`my-network`というDocker Networkを作成する
```
docker network create my-network
```

`my-network`が作成されていることを確認する
```
docker network ls
```

## 2. DatabaseのCotainerにネットワークを割り当て
下記のコマンドで、すでに起動されている`database-container`に`my-network`を接続する

```
docker network connect my-network database-container
```

## 3. api-containerを一度作り直す
環境変数を付与するために`api-container`を再作成する。

下記コマンドでapi-containerを停止する
```
docker stop api-container 
```

続いて下記コマンドでapi-containerを削除する
```
docker rm api-container
```

api-containerを再度起動する
```
docker run -d -p 8080:8080 \
  --name api-container \
  -e DB_USERNAME=root \
  -e DB_PASSWORD=cloudtech \
  -e DB_SERVERNAME=database-container \
  api-image
```

## 4. ネットワークに所属していることを確認
以下のコマンドで、`my-network`に`database-container`と`api-container`が所属しhていることを確認する
```
docker network inspect my-network
```

## 5. 動作確認
`api-container`にて起動しているAPIを呼び出し、データベースにアクセスして結果（データベース接続テストに成功しました）が返却されることを確認する

```
curl http://localhost:8080/test
```

