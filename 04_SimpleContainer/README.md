# 事前準備
下記のGitHubディレクトリから関連資材がダウンロード（`git clone`）されていること<br>
https://github.com/CloudTechOrg/course-docker.git

以下のように必要なファイルが構成されていること
```
course-container/
└── 04_Dockerfile/
    ├── index.html
    └── Dockerfile
```

cdコマンドにより、カレントディレクトリを`04_Dockerfile`に移動する
```
cd 04_Dockerfile
```

## ハンズオン手順

## 1. イメージの作成
下記の`docker build`コマンドにより、Dockerイメージを生成する
```
docker image build -t first-image .
```

下記コマンドにより、`first-image`のイメージが出力されることを確認する
```
docker image ls
```

## 2. Containerの起動
下記の`docker run`コマンドにより、Containerを起動する
```
docker run -d -p 8080:80 --name first-container first-image
```

下記のコマンドにより`first-container`が実行中（STATUSが`Up xxxx`であること）を確認する
```
docker container ls
```

## 3. 動作確認
ブラウザにて下記のURLを実行し、HTMLサイトが表示されることを確認する
```
http://localhost:8080
```

## 4. Containerの停止

下記コマンドで、起動中のContainerを表示する
```
docker container ls
```

続いて下記コマンドで、Containerを停止する
```
docker container stop first-container
```

下記コマンドで起動中のContainerに表示されないことを確認する
```
docker container ls
```

## 4. 停止および削除
以下のコマンドでContainerを削除する
```
docker container rm first-container
```

続いて下記のコマンドでイメージを削除する
```
docker image rm  first-image
```
