# 事前準備
下記のGitHubディレクトリから関連資材がダウンロード（`git clone`）されていること<br>
https://github.com/CloudTechOrg/course-docker.git

以下のように必要なファイルが構成されていること
```
course-container/
└── 07_FrontContainer/
    ├── index.html
    └── Dockerfile
```

cdコマンドにより、カレントディレクトリを`07_FrontContainer`に移動する
```
cd 07_FrontContainer
```

## ハンズオン手順

## 1. イメージの作成
下記の`docker build`コマンドにより、Dockerイメージを生成する
```
docker build -t front-image .
```

下記コマンドにより、`first-image`のイメージが出力されることを確認する
```
docker images
```

## 2. Containerの起動
下記の`docker run`コマンドにより、Containerを起動する
```
docker run -d -p 8081:80 --name front-container front-image
```

下記のコマンドにより`front-container`が実行中（STATUSが`Up xxxx`であること）を確認する
```
docker ps
```

## 3. 動作確認
ブラウザにて下記のURLを実行し、HTMLサイトが表示されることを確認する
```
http://localhost:8081
```
