# 模範解答

# ハンズオン手順

## 1. APIの向き先を変更
1. [config.js](config.js)を開く
2. baseURLの部分を、`http://<APIサーバのALBのDNS名>`に変更する

## 2.ECRリポジトリの作成
1. 「Elastic Container Registry」のダッシュボードを開く
2. `リポジトリを作成`をクリックします。 
3. レジストリ名に`front-repository`と入力し、それ以外の項目は変更不要として、`作成`をクリック    
4. 無事に`front-repository`が作成されていることを確認
5. `URI`の値をコピーしておく

## 2. IAMポリシーを作成
1. IAMのダッシュボードを開く
2. 左側のメニューから、「ポリシー」を選択
3. 右上の「ポリシーの作成」をクリック。
4. アクセス許可の指定部分で、「JSON」タブを選択
5. ポリシーエディターの部分に、下記のjsonを貼り付け（account-idは自分のIDに置き換え）
    
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "ecr:GetAuthorizationToken"
            ],
            "Resource": "*"
        },
        {
            "Effect": "Allow",
            "Action": [
                "ecr:BatchCheckLayerAvailability",
                "ecr:InitiateLayerUpload",
                "ecr:UploadLayerPart",
                "ecr:CompleteLayerUpload",
                "ecr:PutImage"
            ],
            "Resource": "arn:aws:ecr:ap-northeast-1:<account-id>:repository/front-repository"
        }
    ]
}
```

6. 入力が完了したら、`次へ`をクリック
7. ポリシー名に`FrontEcrPushPolicy`を入力
8. それ以外はデフォルトのまま、`ポリシーの作成`をクリックし
7. 無事に`FrontEcrPushPolicy`が作成されていることを確認

## 3. IAMユーザを作成
1. 左側のメニューから`ユーザー`を選択
2. `ユーザーの作成`をクリックします。
3. ユーザー名として、`FrontEcrPushUser`と入力して、`次へ`をクリック
4. 許可のオプションでは、`ポリシーを直接アタッチする`を選択
5. 許可ポリシーからさきほど作成した`FrontEcrPushPolicy`を検索し、それをチェック
6. `次へ`をクリックします。 
7. 確認画面が表示されるので、問題がなければ`ユーザの作成`をクリック    
6. 無事に`FrontEcrPushUser`が作成されていることを確認

## 4. アクセスキーを作成
1. まずは、IAMユーザの一覧から`FrontEcrPushUser`を選択
2. 続いて、`セキュリティ認証情報`のタブを選択
3. `アクセスキーを作成`をクリック
4. ユースケースとして、`コマンドラインインターフェース`を選択
5. `上記のレコメンデーションを確認し、アクセスキーを作成します。`をチェックしたうえで、`次へ`をクリック
6. 説明タグは特は必要に応じて入力し、`アクセスキーを作成`をクリック
7. アクセスキーとシークレットアクセスキーが作成されるので、忘れずにメモしておくか、またはcsvファイルのダウンロードを行う
8. メモが完了したら、`終了`をクリック
   
# 5. アクセスキーの設定
1. `aws configure`コマンドを実行
2. アクセスキー、シークレットアクセスキーを入力

## 6. イメージの作成
下記コマンドを実行

```docker
docker image build --platform linux/x86_64 -t front-image . 
```

## 7. タグ付け
下記コマンドを実行（`your-ecr-uri`の部分はご自身のものに置き換え）

```docker
docker image tag front-image:latest <your-ecr-uri>:latest
```

## 8. ECRにログイン
下記コマンドを実行（`your-ecr-uri`の部分はご自身のものに置き換え）

```docker
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin 125605020607.dkr.ecr.ap-northeast-1.amazonaws.com/front-repository
```

## 9. プッシュ
下記コマンドを実行（`your-ecr-uri`の部分はご自身のものに置き換え）

```docker
docker image push <your-ecr-uri>>y:latest
```

# 10. ECRリポジトリの確認
1. ECRリポジトリの一覧から`front-repository`を選択
2. latestタグが付いているイメージが存在することを確認

# 11. front-routetableを作成
1. 以下のコマンドを実行し、front-routetableを作成する
```
aws ec2 create-route-table  \
  --vpc-id <my-vpcのID> \
  --tag-specifications 'ResourceType=route-table,Tags=[{Key=Name,Value=front-routetable}]'
```
2. 以下のコマンドを実行し、front-routetableをmy-internet-gatewayに関連付ける
```
aws ec2 create-route  \
  --route-table-id <front-routetableのID>  \
  --destination-cidr-block 0.0.0.0/0  \
  --gateway-id <my-internet-gatewayのID>
```

# 12. front-subntt-01の作成
1. 以下のコマンドを実行し、alb-subnet-01を作成する
```
aws ec2 create-subnet --vpc-id <my-vpcのID> \
  --cidr-block 10.0.7.0/24 \
  --availability-zone ap-northeast-1a \
  --tag-specifications 'ResourceType=subnet,Tags=[{Key=Name,Value=front-subnet-01}]'
```
2. 以下のコマンドを実行し、front-subnet-01をfront-routetableに関連付けする
```
aws ec2 associate-route-table  \
  --route-table-id <front-routetableのID>  \
  --subnet-id <front-subnet-01のID>
```

# 13. front-sgの作成
1. 以下のコマンドを実行し、alb-sgを作成
```
aws ec2 create-security-group  \
  --group-name front-sg  \
  --description "front-sg"  \
  --vpc-id <my-vpcのID>
```
2. 以下のコマンドを実行し、alb-sgのインバウンドルールを作成
```
aws ec2 authorize-security-group-ingress  \
  --group-id <front-sgのID> \
  --protocol tcp --port 80 --cidr 0.0.0.0/0
```

# 14. タスク定義の設定
1. ECSのダッシュボードを開く
2. 左側のメニューから、`タスク定義`を選択
3. タスク定義ファミリーに`front-task`を入力
5. コンテナーの定義以下のように入力
    - 名前：`front-container`
    - イメージURI：front-image、latestのECRリポジトリのURI
    - コンテナポートは：`80`とします
6. それ以外はデフォルトのまま`作成`ボタンをクリックします。

# 15. ECSサービスの作成
1. 左側のメニューから`クラスター`を選択し、`MyCluster`をクリック
2. `作成`をクリック
3. `タスク定義のファミリー`に、`front-task`のタスク定義を選択。リビジョンは、`(最新)`とつくものを選択
4. サービス名に`front-service`と入力
4. ネットワーキングの部分を下記のように選択
    - ネットワーク：`my-vpc`
    - サブネット：`front-subnet-01`
    - セキュリティグループ：`api-sg`
5. `作成`をクリックします。
7. `api-task`が起動することを確認

# 16. 動作確認
1. `front-service`をクリック
2. `タスク`のタブを開く
3. 表示されるタスクのIDをクリック
4. パブリックIPアドレスを控える
5. ブラウザで新しいタブを開き、`http://<パブリックIPアドレス>`を実行
6. `API Test`ボタンをクリックすると、`API接続テストが成功しました`というメッセージが表示されることを確認