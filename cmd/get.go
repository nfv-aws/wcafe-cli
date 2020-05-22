package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

// コマンドの追加
func init() {
	RootCmd.AddCommand(newGetCmd())
	RootCmd.AddCommand(newPostCmd())
}

// getコマンドの定義
func newGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get resources",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	// サブコマンドの追加
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
		Short: "Get pets list",
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
