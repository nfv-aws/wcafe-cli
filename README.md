# wcafe CLI

## これなに
HTTPリクエストを自身(localhost)に送るコマンドを作成するリポジトリ

## リポジトリクローン
```
cd $GOPATH/src/github.com
mkdir nfv-aws
cd nfv-aws
git clone git@github.com:nfv-aws/wcafe-api-controller.git
```

## パッケージインストール
```
go get github.com/spf13/cobra
go get github.com/jmcvetta/napping
```

## 動作確認
```
go run main.go stores list
```

## UnitTest
```
go test -v ./cmd/...

PASS
ok      github.com/nfv-aws/wcafe-cli/cmd        0.005s
```

## 使い方

### コマンドのインストール
```
go install
mv $GOPATH/bin/wcafe-cli $GOPATH/bin/wcafe
```
### 一覧データの取得の場合
```
wcafe stores list
wcafe pets　list
wcafe users　list
wcafe clerks list
```
### 新規作成の場合
```
wcafe stores　create
wcafe pets  create <store_id>
wcafe users create
wcafe clerks create -n(--name) hogehoge
(オプションは指定しなくても利用可)
```