# wcafe CLI

## これなに
HTTPリクエストを自身に送るコマンドを作成するリポジトリ

## リポジトリクローン
```
cd $GOPATH/src/github.com
mkdir nfv-aws
cd nfv-aws
git clone git@github.com:nfv-aws/wcafe-api-controller.git
```

### パッケージインストール
```
go get github.com/spf13/cobra
go get github.com/jmcvetta/napping
```

### 環境設定
EC2のプライベートDNSを追加

```
vi ~/.bashrc

export WCAFE_VM_PRIVATE_DNS=private_dns

source ~/.bashrc
```

### 動作確認
```
go run main.go stores list
```

### コマンドのインストール
```
go build
go install
```

### 使い方
#### 一覧データの取得の場合
```
wcafe-cli stores list
wcafe-cli pets　list
wcafe-cli users　list
```
#### 新規作成の場合
```
wcafe-cli stores　create
wcafe-cli pets  create <store_id>
wcafe-cli users create
```