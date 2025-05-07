# 事前準備
今回のハンズオンでは、AWS ClIを使用します。
CLIをインストールしていない場合、下記手順にしたがって事前にインストールをお願いします。

[AWS CLI の最新バージョンのインストールまたは更新](https://docs.aws.amazon.com/ja_jp/cli/latest/userguide/getting-started-install.html)

# ハンズオン手順
## 1.ECRリポジトリの作成
1. 「Elastic Container Registry」のダッシュボードを開く
2. `リポジトリを作成`をクリックします。 
3. レジストリ名に`my-repository`と入力し、それ以外の項目は変更不要として、`作成`をクリック    
4. 無事に`my-repository`が作成されていることを確認
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
            "Resource": "arn:aws:ecr:ap-northeast-1:<account-id>:repository/my-repository"
        }
    ]
}
```

6. 入力が完了したら、`次へ`をクリック
7. ポリシー名に`EcrPushPolicy`を入力
8. それ以外はデフォルトのまま、`ポリシーの作成`をクリックし
7. 無事に`EcrPushPolicy`が作成されていることを確認

## 3. IAMユーザを作成
1. 左側のメニューから`ユーザー`を選択
2. `ユーザーの作成`をクリックします。
3. ユーザー名として、`EcrPushUser`と入力して、`次へ`をクリック
4. 許可のオプションでは、`ポリシーを直接アタッチする`を選択
5. 許可ポリシーからさきほど作成した`EcrPushPolicy`を検索し、それをチェック
6. `次へ`をクリックします。 
7. 確認画面が表示されるので、問題がなければ`ユーザの作成`をクリック    
6. 無事に`EcrPushUser`が作成されていることを確認

## 4. アクセスキーを作成
1. まずは、IAMユーザの一覧から`EcrPushUser`を選択
2. 続いて、`セキュリティ認証情報`のタブを選択
3. `アクセスキーを作成`をクリック
4. ユースケースとして、`コマンドラインインターフェース`を選択
5. `上記のレコメンデーションを確認し、アクセスキーを作成します。`をチェックしたうえで、`次へ`をクリックします。
6. 説明タグは特は必要に応じて入力し、`アクセスキーを作成`をクリックします。   
7. アクセスキーとシークレットアクセスキーが作成されるので、忘れずにメモしておくか、またはcsvファイルのダウンロードを行う
8. メモが完了したら、`終了`をクリックします。
   
# 5. アクセスキーの設定
1. `aws configure`コマンドを実行
2. アクセスキー、シークレットアクセスキーを入力

## 6. イメージの作成
下記コマンドを実行

```docker
docker image build --platform linux/x86_64 -t api-image . 
```

## 7. タグ付け
下記コマンドを実行（`your-ecr-uri`の部分はご自身のものに置き換え）

```docker
docker image tag api-image:latest <your-ecr-uri>:latest
```


## 8. ECRにログイン
下記コマンドを実行（`your-ecr-uri`の部分はご自身のものに置き換え）

```docker
aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin <your-ecr-uri>
```

## 9. プッシュ
下記コマンドを実行（`your-ecr-uri`の部分はご自身のものに置き換え）

```docker
docker image push <your-ecr-uri>>y:latest
```

# 10. ECRリポジトリの確認
1. ECRリポジトリの一覧から`my-repository`を選択
2. latestタグが付いているイメージが存在することを確認
