# 前提
- 踏み台サーバ（JUMPサーバ）にsshなどでログインしていること
- EC2インスタンスが`Amazon Linux 2023`で起動されていること

# 手順

## 1. mysqlのインストール

パッケージマネージャー（yum）の更新を行う
```shell
sudo yum update -y
```

MySQLの公式リポジトリをシステムに追加する
```shell
sudo yum install https://dev.mysql.com/get/mysql84-community-release-el9-1.noarch.rpm
```

MySQLサーバのインストールを行う
```shell
sudo yum install mysql-community-server -y
```

MySQLサービスの起動を行う
```shell
sudo systemctl start mysqld
```

システム起動時にMySQLが自動起動するように設定する
```shell
sudo systemctl enable mysqld
```

### 2. RDSに接続

以下のコマンドで、RDSのMySQLに接続
```
mysql -h testdb.co7vcyxdczfg.ap-northeast-1.rds.amazonaws.com -P 3306 -u admin -p
```

AWS RDSインスタンスのMySQLデータベースに接続します。以下のコマンドを実行前に、適切なエンドポイントとユーザ名に置き換えてください。また、パスワードの入力を求められるため、はRDSインスタンス作成時に指定したものを入力してください。


## 3. テストデータを作成
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