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

RUN echo "テスト"

# ワーキングディレクトリを設定
WORKDIR /root/

# ビルドした実行ファイルをコピー
COPY --from=builder /app/server .

RUN if [ -f ./server ]; then echo "ある"; else echo "ない"; fi

# ポート8080を開放
EXPOSE 8080

# 実行可能ファイルを実行
CMD ["./server"]
