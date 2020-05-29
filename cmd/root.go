package cmd

import (
	"github.com/spf13/cobra"
	"net/http"
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

func newDefaultClient() (*Client, error) {
	endpointURL := "http://localhost:8080/api/v1"
	httpClient := &http.Client{}
	return newClient(endpointURL, httpClient)
}
