# 事前準備
下記のGitHubディレクトリから関連資材がダウンロード（`git clone`）されていること<br>
https://github.com/CloudTechOrg/course-docker.git

cdコマンドにより、カレントディレクトリを`04_DatabaseContainer`に移動する
```
cd 05_DatabaseContainer
```

## ハンズオン手順

## 1. イメージの取得
下記の`docker pull`コマンドにより、MySQLをベースとしたイメージを取得する
```
docker pull mysql:latest
```

## 2. Containerの起動
下記の`docker run`コマンドにより
```
docker run --name database-container -e MYSQL_ROOT_PASSWORD=cloudtech -d mysql:latest
```

## 3. MySQLにログイン
下記の`docker exec`コマンドにより、MySQLにログインする
```
docker exec -it database-container mysql -u root -p
```

パスワードの入力が促されるので、Container起動時に指定した`MYSQL_ROOT_PASSWORD`の値を入力
```
Enter password: 
```

## 5. データベースの作成
下記のSQLを実行し、`testdb`データベースを作成する
```sql
CREATE DATABASE testdb;
```

## 6. テーブルの作成
下記のSQLを実行し、`test_table`テーブルを作成する
```sql
CREATE TABLE testdb.test_table (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    age INT
);
```

## 7. データの作成
下記のSQLを実行し、テスト用のデータを登録する
```sql
INSERT INTO
    testdb.test_table (name, age) 
    VALUES ('テスト太郎', 25);
```

## 8. データの確認
下記のSQLを実行し、登録されたデータの確認をする
```sql
SELECT * FROM testdb.test_table;
```

## 9. MySQLからログアウト
以下のコマンドを実行し、MySQLからログアウトする
```
exit
```