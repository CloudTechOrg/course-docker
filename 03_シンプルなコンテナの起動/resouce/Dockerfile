# ベースイメージの取得
FROM nginx:alpine

# HTMLファイルのコピー
COPY index.html /usr/share/nginx/html/index.html

# Nginxの起動する
CMD ["nginx", "-g", "daemon off;"]