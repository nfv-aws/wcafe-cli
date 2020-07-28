package cmd

import (
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use: "wcafe",
	}
	// サブコマンドの追加
	cmd.AddCommand(newStoresCmd())
	cmd.AddCommand(newPetsCmd())
	cmd.AddCommand(newUsersCmd())
	cmd.AddCommand(newClerksCmd())
	return cmd
}

func Execute() {
	cmd := NewCmdRoot()
	cmd.SetOutput(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOutput(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}

func newDefaultClient() (*Client, error) {
	endpointURL := "http://localhost:8080/api/v1"
	httpClient := &http.Client{}
	return newClient(endpointURL, httpClient)
}

// コマンドの追加
func init() {
}
