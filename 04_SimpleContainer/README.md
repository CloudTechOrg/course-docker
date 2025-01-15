# 事前準備
任意のフォルダーに以下のファイルがダウンロードされていること
```
任意のフォルダ
├── index.html
└── Dockerfile
```

# ハンズオン手順

## 1. イメージの作成
`first-image`のイメージを作成する
```
docker image build --tag first-image .
```

`first-image`のイメージが出力されることを確認する
```
docker image ls
```

## 2. コンテナーの起動
`first-container`のコンテナーを起動する
```
docker run -d -p 8080:80 --name first-container first-image
```

`first-container`が実行中（STATUSが`Up xxxx`）であることを確認する
```
docker container ls
```

## 3. 動作確認
ブラウザにて下記のURLを実行し、HTMLサイトが表示されることを確認する
```
http://localhost:8080
```

## 4. Containerの停止
`first-container`を停止する
```
docker container stop first-container
```

`first-container`を表示されないことを確認する
```
docker container ls
```

## 4. 停止および削除
`first-container`を削除する
```
docker container rm first-container
```

`first-image`を削除する
```
docker image rm  first-image
```
