# 事前準備
- Dockerがインストールされていること

# ハンズオン手順

## 1. イメージの入手
`hello-world`のイメージを取得する
```
docker pull hello-world
```

## 2. イメージの一覧取得
イメージの一覧を表示する

```
docker image ls
```

## 3.コンテナーの起動
`hello-world`のコンテナーを起動する

```
docker container run hello-world
```

## 4. コンテナーの一覧表示
下記コマンドで`hello-world`のコンテンナーが表示`されないこと`を確認（起動後すぐに停止するため）

```
docker container ls
```

下記コマンドで`hello-world`のContainerが表示されることを確認
```
docker container ls --all
```

## 5. Containerの削除
コンテナーを削除する。（`<container-id>`の部分は、`docker container ls --all`コマンドで表示される`CONTAINER ID`を設定する。

```
docker container rm <container-id>
```

`hello-world`のコンテナーが表示されないことを確認

```
docker container ls -a
```

## 6. イメージの削除
`hello-world`のイメージで削除する

```
docker image rm hello-world
```

以下のコマンドで、`hello-world`のイメージが表示されないことを確認する

```
docker image ls
```