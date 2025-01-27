# ハンズオン手順

## 1. Docker Composeの起動
```
docker compose up -d
```

## 2. Docker Composeの起動確認
```
docker compose -p
```

## 3. 動作確認
ブラウザにて下記のURLを実行し、HTMLサイトが表示されることを確認する
```
http://localhost:8081
```

## 4. Docker Composeの削除
```
docker compose down --rmi all
```