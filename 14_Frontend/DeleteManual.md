# 1. ECS
## 1.1 クラスター
### MyCluster
- MyClusterをクリックし、`クラスターの削除`をクリック
- エラーが発生する場合、`api-service`と`front-service`を削除してから実施
## 1.2 タスク定義
### api-task
- `api-task`をクリックし、作成されているタスク定義をすべてチェックしたうえで、`アクション`→`登録解除`をクリック
### front-task
- `front-task`をクリックし、作成されているタスク定義をすべてチェックしたうえで、`アクション`→`登録解除`をクリック
# 2. ECR
## 2.1 Repositories
### my-repository
- `my-repository`をチェックし、`削除`をクリック
### front-repository
- `front-repository`をチェックし、`削除`をクリック
# 3. RDS
## 3.1 データベース
### testdb
- `testdb`をチェックし、`アクション`→`削除`
## 3.2 サブネットグループ
### my-subnet
- `my-subnet`をチェックし、`削除`
# 4. EC2
## 4.1 インスタンス
### JumpServer
- `JumpServer`をチェックし、`インスタンスの状態`→`インスタンスの終了（削除）`
## 4.2 ロードバランサー
### alb-lb
- `alb-sg`をチェックし、`アクション`→`ロードバランサーの削除`
## 4.3 ターゲットグループ
### api-tg
- `api-tg`をチェックし、`アクション`→`削除`
# 5. VPCダッシュボード
## 5.1 NATゲートウェイ
### my-nat-gateway
- `my-nat-gateway`をチェックし、`アクション`→`NATゲートウェイを削除`
## 5.2 VPC
### my-vpc
- `my-vpc`を選択し、`アクション`→`削除`をクリック
- サブネット、ルートテーブル、インターネットゲートウェイ、セキュリティグループも一緒に削除される
# 6. IAM
## 6.1 IAMユーザ
### EcrPushUser
- 該当のユーザを選択し、`削除`をクリック
### FrontEcrPushUser
- 該当のユーザを選択し、`削除`をクリック
## 6.2 IAMポリシー
### EcrPushPolicy
- 該当のポリシーを選択し、`削除`をクリック
### FrontEcrPushPolicy
- 該当のポリシーを選択し、`削除`をクリック
