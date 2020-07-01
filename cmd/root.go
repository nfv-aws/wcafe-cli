package cmd

import (
	"net/http"

	"github.com/spf13/cobra"
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

// コマンドの追加
func init() {
	RootCmd.AddCommand(newStoresCmd())
	RootCmd.AddCommand(newPetsCmd())
	RootCmd.AddCommand(newUsersCmd())
	RootCmd.AddCommand(newClerksCmd())
}
