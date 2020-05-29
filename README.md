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
go install
mv $GOPATH/bin/wcafe-cli $GOPATH/bin/wcafe
```

### 使い方
#### 一覧データの取得の場合
```
wcafe stores list
wcafe pets　list
wcafe users　list
```
#### 新規作成の場合
```
wcafe stores　create
wcafe pets  create <store_id>
wcafe users create
```