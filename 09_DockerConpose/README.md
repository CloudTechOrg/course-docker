# ハンズオン手順

## 1. Docker Composeの起動
```
docker compose up -d
```

## 2. Docker Composeの起動確認
```
docker compose ps
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

# トラブルシューティング

## Docker Compose実行時にfrontコンテナーがunhealthyになり、403エラーが発生する

### 事象
Docker Composeで`docker compose up -d`を実行してコンテナーを起動したところ。front-containerのステータスが `unhealthy` になった。また、`http://localhost:8081` にアクセスすると `403 Forbidden` エラーが発生した。

### 原因
- frontディレクトリ配下のファイルに、他ユーザ（例： コンテナー内の実行ユーザ）に対する読み取り権限がないため、コンテナー内で必要なファイルにアクセスできず、403エラーが発生したと考えられる。
- 必ず発生する事象ではなく、利用者の端末に影響される

### 対応策
1. ターミナルで `front` ディレクトリに移動
2. `front` 配下のすべてのファイルに他ユーザへの読み取り権限を付与

   ```bash
   chmod 644 *
   ```
3. Dockerイメージをキャッシュを使わず再ビルドし、再起動
    ```bash
    docker compose down
    docker compose build --no-cache front
    docker compose up -d
    ```

### 備考
- api-containerにおいても同様の事象が発生した場合、同じ手順で対応可能