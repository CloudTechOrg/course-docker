# ベースイメージの取得
FROM nginx:alpine

# curl をインストール（ヘルスチェック用などに必要）
RUN apk update && apk add --no-cache curl

# HTMLファイルのコピー
COPY . /usr/share/nginx/html/

# Nginxの起動する
CMD ["nginx", "-g", "daemon off;"]