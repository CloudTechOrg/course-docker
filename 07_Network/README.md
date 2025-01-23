# 事前準備
以下のコンテナーが起動中であること
- api-container
- db-container

もし停止中である場合、以下のコマンドで起動させる
```
docker container start api-container
docker container start db-container
```

# ハンズオン手順

## 1. Dockerネットワークを作成
`api-network`を作成する
```
docker network create api-network
```

`api-network`が作成されていることを確認する
```
docker network ls
```

## 2. ネットワークにコンテナーを割り当て
下記のコマンドで、すでに起動されているコンテナーを`my-network`に接続する

`db-container`を`api-network`に割り当て
```
docker network connect api-network db-container
```

`api-container`を`api-network`に割り当て
```
docker network connect api-network api-container
```

ネットワークの状態を確認
```
docker network inspect api-network
```

## 3. 動作確認
下記の`curl`コマンドにより「データベース接続テストが成功しました（test_tableの件数：3）」のメッセージが表示されることを確認する（件数は実際に登録したデータ件数に応じる）

```
curl http://localhost:8080/dbtest
```

# 補足
`docker container run`実行時に`--network`オプションを指定すると、コンテナー起動時にネットワークを関連付けることが可能

```
docker container run --name <container-name> --network <network-name> <image-name>
```