package cmd

import (
	"fmt"
	"github.com/nfv-aws/wcafe-cli/config"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

var (
	endpoint string
)

// サブコマンドの追加
func init() {
	config.Configure()
	endpoint = config.C.LB.Endpoint
	RootCmd.AddCommand(newGetCmd())
	RootCmd.AddCommand(newPostCmd())
}

// getコマンド
func newGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	// getコマンドのオプションの追加
	cmd.AddCommand(
		newGetPetsCmd(),
		newGetStoresCmd(),
		newGetUsersCmd(),
	)

	return cmd
}

func newGetPetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pets",
		Short: "Get Pets list",
		RunE:  runGetPetsCmd,
	}
	return cmd
}

func newGetStoresCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stores",
		Short: "Get stores list",
		RunE:  runGetStoresCmd,
	}
	return cmd
}

func newGetUsersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "Get users list",
		RunE:  runGetUsersCmd,
	}
	return cmd
}

func runGetPetsCmd(cmd *cobra.Command, args []string) error {
	url := endpoint + "/api/v1/pets"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	return nil
}

func runGetStoresCmd(cmd *cobra.Command, args []string) error {
	url := endpoint + "/api/v1/stores"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	return nil
}

func runGetUsersCmd(cmd *cobra.Command, args []string) error {
	url := endpoint + "/api/v1/users"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	return nil
}
