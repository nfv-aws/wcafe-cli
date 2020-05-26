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
LBのエンドポイントを追加

```
vi ~/.bashrc

export WCAFE_VM_PRIVATE_DNS=private_dns

source ~/.bashrc
```

### コマンドのインストール
```
go build
go install
```

### 使い方
#### getの場合
```
wcafe-cli get stores
wcafe-cli get pets
wcafe-cli get users
```
#### postの場合
```
wcafe-cli post stores
wcafe-cli post pets <store_id>
wcafe-cli post users
```