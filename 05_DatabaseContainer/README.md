## ハンズオン手順

## 1. イメージの取得
`mysql:latest`のベースイメージを取得する
```
docker pull mysql:latest
```

## 2. Containerの起動
`db-container`を起動する
```
docker container run --name db-container -e YSQL_ROOT_PASSWORD=cloudtech mysql:latest
```

## 3. MySQLにログイン
MySQLにログインする
```
docker exec -it db-container mysql -u root -p
```

パスワードの入力が促されるので、Container起動時に指定した`MYSQL_ROOT_PASSWORD`の値を入力する
```
Enter password: 
```

## 5. テストデータを作成
`testdb`データベースを作成する
```sql
CREATE DATABASE testdb;
```

`test_table`テーブルを作成する
```sql
CREATE TABLE testdb.test_table (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    age INT
);
```

テスト用のデータを登録する
```sql
INSERT INTO
    testdb.test_table (name, age) 
    VALUES ('Test Taro', 30);
```

```sql
INSERT INTO
    testdb.test_table (name, age) 
    VALUES ('Test Jiro', 22);
```

```sql
INSERT INTO
    testdb.test_table (name, age) 
    VALUES ('Test Hanako', 25);
```

登録されたデータの確認をする
```sql
SELECT * FROM testdb.test_table;
```

## 9. MySQLからログアウト
MySQLからログアウトする
```
exit
```

起動したdb-containerは、このあとも利用するため削除せずに残しておく