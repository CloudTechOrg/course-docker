-- データベースを作成して選択
CREATE DATABASE IF NOT EXISTS testdb;
USE testdb;

-- テーブルが既に存在する場合は削除して再作成
DROP TABLE IF EXISTS test_table;

CREATE TABLE test_table (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50),
    age INT
);

-- 初期データを挿入
INSERT INTO test_table (name, age) VALUES ('Test Taro', 30);
INSERT INTO test_table (name, age) VALUES ('Test Jiro', 22);
INSERT INTO test_table (name, age) VALUES ('Test Hanako', 25);
INSERT INTO test_table (name, age) VALUES ('Test Youko', 25);
