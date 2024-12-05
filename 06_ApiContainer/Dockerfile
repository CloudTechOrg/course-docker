# ベースイメージとして公式のGo 1.23を使用
FROM golang:1.23

# 作業ディレクトリを設定
WORKDIR /usr/src/app

# 依存関係のダウンロードに必要なファイルをコピー
COPY go.mod go.sum ./

# go.modに基づいて依存関係をダウンロードし、整合性を検証
RUN go mod download && go mod verify

# アプリケーションの全ファイルをコンテナ内にコピー
COPY . .

# Goアプリケーションをビルドし、実行可能ファイルを指定の場所に出力
RUN go build -v -o /usr/local/bin/app ./...

# コンテナ起動時に実行するコマンドを指定
CMD ["app"]
