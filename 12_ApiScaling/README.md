# ハンズオン手順

## api-subnet-01をプライベートサブネットにする
1. VPCのダッシュボードを開く
2. 左側のメニューから、`ルートテーブル`をクリック
3. `api-routetable`をクリック
4. `ルートを編集`をクリック
5. インターネットゲートウェイに向いているルートを、`削除`ボタンにより削除
6. `変更を保存`をクリック
7. 正常終了のメッセージの表示を確認

## api-subnet-02を作成する
1. CloudShellを開く
2. 下記のCLIコマンドを実行し、`api-subnet-02`を作成する
```
aws ec2 create-subnet --vpc-id <my-vpcのID> \
  --cidr-block 10.0.1.0/24 \
  --availability-zone ap-northeast-1c \
  --tag-specifications 'ResourceType=subnet,Tags=[{Key=Name,Value=api-subnet-02}]'
```

3. 以下のコマンドを実行し、ルートテーブルの関連付けを行う

```
aws ec2 associate-route-table  \
  --route-table-id <api-routetableのID>  \
  --subnet-id <api-subnet-02のID>
```

## alb-routetableの作成
1. 以下のコマンドを実行し、alb-routetableを作成する
```
aws ec2 create-route-table  \
  --vpc-id <my-vpcのID> \
  --tag-specifications 'ResourceType=route-table,Tags=[{Key=Name,Value=alb-routetable}]'
```

2. 以下のコマンドを実行し、alb-routetableをmy-internet-gatewayに関連付ける
```
aws ec2 create-route  \
  --route-table-id <alb-routetableのID>  \
  --destination-cidr-block 0.0.0.0/0  \
  --gateway-id <my-internet-gatewayのID>
```

## alb-subnet-01を作成
1. 以下のコマンドを実行し、alb-subnet-01を作成する
```
aws ec2 create-subnet --vpc-id <my-vpcのID> \
  --cidr-block 10.0.3.0/24 \
  --availability-zone ap-northeast-1c \
  --tag-specifications 'ResourceType=subnet,Tags=[{Key=Name,Value=alb-subnet-02}]'
```

## alb-subnet-02を作成
1. 以下のコマンドを実行し、alb-subnet-01を作成する
```
aws ec2 create-subnet --vpc-id <my-vpcのID> \
  --cidr-block 10.0.2.0/24 \
  --availability-zone ap-northeast-1a \
  --tag-specifications 'ResourceType=subnet,Tags=[{Key=Name,Value=alb-subnet-01}]'
```

## alb-subnet-01にalb-routetableを関連付ける
1. 以下のコマンドを実行し、alb-subnet-01をalb-routetableに関連付けする
```
aws ec2 associate-route-table  \
  --route-table-id <alb-routetableのID>  \
  --subnet-id <alb-subnet-01のID>
```

## alb-subnet-02にalb-routetableを関連付ける
1. 以下のコマンドを実行し、alb-subnet-02をalb-routetableに関連付けする
```
aws ec2 associate-route-table  \
  --route-table-id <alb-routetableのID>  \
  --subnet-id <alb-subnet-02のID>
```

## albのセキュリティグループを作成
1. 以下のコマンドを実行し、alb-sgを作成
```
aws ec2 create-security-group  \
  --group-name alb-sg  \
  --description "alb-sg"  \
  --vpc-id <my-vpcのID>
```

2. 以下のコマンドを実行し、alb-sgのインバウンドルールを作成
```
aws ec2 authorize-security-group-ingress  \
  --group-id <alb-sgのID> \
  --protocol tcp --port 80 --cidr 0.0.0.0/0
```

## api-sgのアクセス制限
1. セキュリティグループの一覧から、`api-sg`を選択
2. `インバウンドルールの編集`をクリック
3. 現在のインバウンドルールに対して、`削除`をクリック
4. `ルールを追加`をクリック
5. インバウンドルールを設定
    1. タイプ：`カスタムTCP`を入力
    2. ポート範囲：`8080`を入力
    3. ソース：`alb-sg`を選択
6. `ルールを保存`をクリック
7. 正常終了を表すメッセージを確認

## NATゲートウェイを作成
1. VPCのダッシュボードを開く
2. 左側のメニューから、`NATゲートウェイ`を開く
3. `NATゲートウェイを作成`をクリック
4. NATゲートウェイの設定を行う
    1. 名前：`my-nat-gateway`
    2. サブネット：`alb-subnet-01`（パブリクサブネットなら何でも良い）
    3. 接続タイプ：`パブリック`
    4. Elastic IP割り当てID：`Elastic IPを割り当て`
5. `NATゲートウェイを作成`をクリック
6. NATゲートウェイが正常に作成されることを確認
7. ルートテーブルの一覧から、`api-routetable`を開く    
1. `api-routetable`に、以下を追加し、`変更を保存`をクリック
    1. 送信先：`0.0.00/0`
    2. ターゲット：`api-nat-gateway`

# ターゲットグループを作成
1. EC2のダッシュボードを開く
2. 左側のメニューから`ターゲットグループ`を選択
3. `ターゲットグループの作成`をクリック
4. ターゲットグループの選択に、`IPアドレス`を選択
5. ターゲットグループに`api-tg`、ポートに`8080`、vpcに`my-vpc`を選択
6. `次へ`をクリック
7. `ターゲットグループの作成`をクリック
8. 無事に`alb-sg`が作成されていることを確認

# ALBを作成
1. 左側のメニューから`ロードバランサー`を選択
2. `ロードバランサーの作成`をクリック
3. `Application Load Balancer`を選択
4. 基本的な設定を入力
    1. ロードバランサー名は、`api-lb`
    2. スキームは、外部からアクセスしたいので`インターネット向け`を選択
5. ネットワークマッピングを入力
    1. VPCは、`my-vpc`を選択
    2. アベイラビリティゾーンは、`ap-northeast-1a`と`ap-northeast-1b`を選択
    3. そのうえで、`alb-subnet01`と`alb-subnet-02`を選択
6. セキュリティグループは、`alb-sg`を選択します。
7. リスナーとルーティングの部分を選択
    1. プロトコルは`HTTP`として、ポートは`80`を指定します
    2. デフォルトアクションとして、ターゲットグループの`api-tg`を指定
8. `ロードバランサーの作成`をクリックします。
9. 無事に`api-lb`が作成されていれば、ここまでの操作は成功です。

# ECSサービスを再作成
1. ECSのダッシュボードを開く
2. 左側のメニューから、`クラスター`を選択します。
3. `MyCluster`を選択
4. `api-service`を選択し、`サービスを削除`をクリック

## ECSサービスの再作成
1. デプロイ設定を入力
    1. ファミリーは`api-task`、リビジョンは`最新`とつくものを選択
    2. 必要なタスクについて今回は`2`と指定
2. ネットワーキングの設定を行います。
    1. VPCは、`my-vpc`を選択
    2. サブネットは、`api-subnet-01`と`api-subnet-02`を選択
3. ロードバランシングの設定を行います。
    1. `ロードバランサーを使用`をチェック
    2. ロードバランサーの種類として、今回は`Application Load Balancer`を選択
    3. `既存のロードバランサーを使用`を選択
    4. ロードバランサーとして、先ほど作成した`api-lb`を指定
    5. リスナーは`既存のリスナーを使用`を選択して、`80:HTTP`を指定
    6. ターゲットグループは`既存のターゲットグループを使用`を選択し、`api-tg`を指定    
4. オートスケーリングとして、`サービスの自動スケーリング`の設定
    1. `サービスの自動スケーリング`を使用をチェックします。
    2. タスクの最小数として`2`、タスクの最大数として`4`を指定
    3. スケーリングポリシータイプとして、`ターゲットの追跡`を指定
    4. ポリシー名は、ここでは`api-auto-scaling`を指定します。
    5. ECSサービスメトリクスは、`ECSServiceAverageCPUUtilization`を選択
    6. ターゲット値は`70`とsuru

    
# 動作確認
1. ALBのドメイン名を取得
2. ブラウザの新しいタブを開く
3. `http://<albのドメイン名>により、APIにアクセス