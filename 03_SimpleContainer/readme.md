# 事前準備
下記のGitHubディレクトリから関連資材がダウンロード（`git clone`）されていること<br>
https://github.com/CloudTechOrg/course-docker.git

以下のように必要なファイルが構成されていること
```
course-container/
└── 03_simple_container/
    ├── index.html
    └── Dockerfile
```

cdコマンドにより、カレントディレクトリを`03_SimpleContainer`に移動する
```
cd 03_SimpleContainer
```

## ハンズオン手順

## 1. イメージの作成
下記の`docker build`コマンドにより、Dockerイメージを生成する
```
docker build -t first-image .
```

下記コマンドにより、`first-image`のイメージが出力されることを確認する
```
docker images
```

## 2. Containerの起動
下記の`docker run`コマンドにより、Containerを起動する
```
docker run -d -p 8080:80 --name first-container first-image
```

下記のコマンドにより`first-container`が実行中（STATUSが`Up xxxx`であること）を確認する
```
docker ps
```

## 3. 動作確認
ブラウザにて下記のURLを実行し、HTMLサイトが表示されることを確認する
```
http://localhost:8080
```

## 4. 停止および削除
以下の`docker stop`コマンドで、`first-container`を停止する
```
docker stop first-container
```

ダウンロードしたイメージを削除する
