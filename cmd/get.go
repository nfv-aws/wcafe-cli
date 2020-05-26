package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
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
		newGetStoresCmd(),
		newGetPetsCmd(),
		newGetUsersCmd(),
	)

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

func newGetPetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pets",
		Short: "Get pets list",
		RunE:  runGetPetsCmd,
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

func runGetStoresCmd(cmd *cobra.Command, args []string) error {
	url := "http://" + dns + ":8080/api/v1/stores"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	return nil
}

func runGetPetsCmd(cmd *cobra.Command, args []string) error {
	url := "http://" + dns + ":8080/api/v1/pets"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	return nil
}

func runGetUsersCmd(cmd *cobra.Command, args []string) error {
	url := "http://" + dns + ":8080/api/v1/users"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	return nil
}
