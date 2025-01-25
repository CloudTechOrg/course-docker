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

## 1. イメージの作成
`front-image`を作成する
```
docker imagebuild -t front-image .
```

`front-image`が作成されていることを確認する
```
docker image ls
```

## 2. コンテナーの起動
`front-container`を起動する
```
docker container run -p 8081:80 --name front-container front-image
```

`front-container`が実行中（STATUSが`Up xxxx`で）あることを確認する
```
docker container ls
```

## 3. 動作確認
ブラウザにて下記のURLを実行し、HTMLサイトが表示されることを確認する
```
http://localhost:8081
```

## 4. コンテナーとイメージの削除
コンテナーの停止
```
docker container stop db-container
docker container stop api-container
docker container stop front-container
```

コンテナーの削除
```
docker container rm db-container
docker container rm api-container
docker container rm front-container
```

イメージの削除
```
docker image rm db-image
docker image rm api-image
docker image rm front-image
```

## 5. Docker Networkの削除
`api-network`の削除
```
docker network rm api-image
```

`api-network`の削除確認
```
docker network ls
```