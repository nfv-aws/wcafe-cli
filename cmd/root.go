package cmd

import (
	"github.com/spf13/cobra"
	"net/http"

	"github.com/nfv-aws/wcafe-cli/config"
)

var (
	// ルートコマンドの設定
	RootCmd = &cobra.Command{
		Use: "wcafe",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
)

// Private DNS
var dns string

// コンフィグの呼び出し
func init() {
	config.Configure()
	dns = config.C.VM.Private_dns
}

func newDefaultClient() (*Client, error) {
	endpointURL := "http://localhost:8080/api/v1"
	httpClient := &http.Client{}
	return newClient(endpointURL, httpClient)
}
