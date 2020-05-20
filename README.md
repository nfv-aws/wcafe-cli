# wcafe CLI

## これなに
curl文を実行するコマンドを作成するリポジトリ

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
go get github.com/spf13/cobra/cobra
```

### 使い方（想定）
wcafe get pets
⇒curl localhost:8080/api/v1/pets | jq .　　が実行される。