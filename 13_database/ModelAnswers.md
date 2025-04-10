# 模範解答

## 1. db-subnet-01を作成
1. CloudShellを開く
2. 以下のCLIコマンドを実行し、`db-subnet-01`を作成する
```
aws ec2 create-subnet --vpc-id <my-vpcのID> \
  --cidr-block 10.0.4.0/24 \
  --availability-zone ap-northeast-1a \
  --tag-specifications 'ResourceType=subnet,Tags=[{Key=Name,Value=db-subnet-01}]'
```

## 2. db-subnet-02を作成
1. CloudShellを開く
2. 以下のCLIコマンドを実行し、`db-subnet-02`を作成する
```
aws ec2 create-subnet --vpc-id <my-vpcのID> \
  --cidr-block 10.0.5.0/24 \
  --availability-zone ap-northeast-1c \
  --tag-specifications 'ResourceType=subnet,Tags=[{Key=Name,Value=db-subnet-02}]'
```

## 3. サブネットグループを作成
1. RDSのダッシュボードを開く
2. 左側のメニューから`サブネットグループ`を選択
3. 右上の`DBサブネットグループを作成`をクリック
4. サブネットグループの詳細は以下の通り入力
    1. 名前：`my-subnet-group`
    2. 説明：`my subnet group`
    3. VPC：`my-vpc`を選択
    4. アベイラビリティゾーン：`ap-northeast-1a`、`ap-northeast-1c`を選択
    5. サブネット：さきほど作った`db-subnet-01`と`db-subnet-02`を選択
6. `作成`をクリック
7. サブネットグループが正しく作成されることを確認
    
## 4. RDSインスタンスを作成
1. 左側のメニューから、`データベース`を選択する
2. 右上の`データベースの作成`をクリックする
3. データベースの作成方法は、`標準作成`を選択
4. エンジンのタイプは、`MySQL`を選択する
5. テンプレートは、`無料利用枠`を選択する
6. 可用性と耐久性はとくに変更不要（無料利用枠のため変更できない）
7. 設定は以下の通りに選択する
    1. DBインスタンス識別子：`testdb`
    2. マスターユーザー名：`admin`
    3. 認証情報管理：`セルフマネージド`
    4. マスターパスワード：任意
8. インスタンスの設定は、とくに変更なし
9. ストレージはとくに変更不要
10. 接続は以下のように設定
    1. コンピューティングリソース：`EC2コンピューティングリソースに接続しない`
    2. VPC：`my-vpc`
    3. DBサブネットグループ：`my-subnet-group`
    4. パブリックアクセス：`なし`
    5. VPCセキュリティグループ：`新規作成`
    6. 新しいVPCセキュリティグループ名：`db-sg`
    7. アベイラビリティーゾーン：`ap-northeast-1a`
    8. RDS Proxy：`チェックしない`
    9. 認証機関：`デフォルト`
11. タグはとくに設定不要
12. データベース認証は`パスワード認証`を選択
13. モニタリング以下は設定不要
14. `データベースの作成`をクリック
15. データベースが正常に作成されることを確認

## 5. jump-subnetを作成
1. CloudShellを開く
2. 以下のコマンドで、jump-subnetを作成する
```
aws ec2 create-subnet --vpc-id <my-vpcのID> \
  --cidr-block 10.0.6.0/24 \
  --availability-zone ap-northeast-1a \
  --tag-specifications 'ResourceType=subnet,Tags=[{Key=Name,Value=jump-subnet}]'
```
3. 以下のコマンドで、ルートテーブルを`alb-routetble`にするを設定
```
aws ec2 associate-route-table  \
  --route-table-id <alb-routetableのID>  \
  --subnet-id <jump-subnetのID>
```

## 6. 踏み台サーバを作成
1. EC2のダッシュボードを開く
2. 左側のメニューから、`インスタンス`をクリック
3. 右上の`インスタンスを起動`をクリック
4. 名前：`JumpServer`
5. Amazonマシンイメージ：`Amazon Linux 2023 AMI`
6. インスタンスタイプ：`t2.micro`
7. VPC：`my-vpc`
8. サブネット：`jump-subnet`
9. パブリックIPの自動割り当て：`有効化`
10. `セキュリティグループを作成`をクリック
11. セキュリティグループ名：`jump-sg`
12. `インスタンスを起動`をクリック

## 7. db-sgの設定変更（jumpサーバからのMySQL通信を許可）
1. セキュリティグループの一覧から、`db-sg`を選択
2. `インバウンドルールを編集`をクリック
3. 既存のルールを`削除`
4. `ルールを追加`をクリック
5. タイプ：`MYSQL/Aurora`
6. ソース：`jump-sg`
7. `ルールを保存`をクリック

## 8. db-sgの設定変更（APIサーバからのMySQL通信を許可）
1. `ルールを追加`をクリック
2. タイプ：`MYSQL/Aurora`
3. ソース：`api-sg`
4. `ルールを保存`をクリック

## 9. データベースの初期設定
1. [CreateJumpServer.md](CreateJumpServer.md)の手順にしたがってデータベースの初期設定を行う

## 10. ECSタスクに環境変数を設定
1. ECSのダッシュボードを開く
2. 左側のメニューから、`タスク定義`を開く
3. `api-task`をクリック
4. 右上の`新しいリビジョンの作成`をクリック
5. 環境変数（オプション）のところにある環境変数を設定
  - DB_SERVERNAME：RDSのエンドポイント
  - DB_USERNAME：RDS作成時に指定したユーザ名（デフォルトは`admin`）
  - DB_PASSWORD：RDS作成時に指定したパスワード

## 11. api-serviceの更新
1. 左側のメニューから、`クラスター`を選択
2. `MyCluster`を選択
3. サービスから`api-service`を選択
4. 右上の`サービスを更新`をクリック
5. タスク定義のリビジョンを、（最新）とつくものに変更
6. `更新`をクリック
7. タスクが再起動されることを確認、されない場合、タスクを停止し、再起動する

## 12. 動作確認
以下のURLで、データベース接続が成功したメッセージが表示されることを確認
http://<ALBのDNSドメイン名>/dbtest