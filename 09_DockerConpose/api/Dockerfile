# --------------------
# ビルドステージ
# --------------------
# ベースイメージを指定（Go言語 ver1.23）
FROM golang:1.23 AS builder

# ワーキングディレクトリを設定
WORKDIR /app

# ソースコードをコピー
COPY . .

# 依存関係をダウンロード
RUN go mod download

# 実行可能ファイルをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# --------------------
# 実行ステージ
# --------------------
# ベースイメージを指定（Alpine Linuxの最新版）
FROM alpine:latest

# curl をインストール（ヘルスチェック用などに必要）
RUN apk update && apk add --no-cache curl

# ワーキングディレクトリを設定
WORKDIR /root/

# ビルドした実行ファイルをコピー
COPY --from=builder /app/server .

# ポート8080を開放
EXPOSE 8080

# 実行可能ファイルを実行
CMD ["./server"]
