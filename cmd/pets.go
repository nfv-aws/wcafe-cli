package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	// "strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// petsコマンドの定義
func newPetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pets",
		Short: "pets  command",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	// サブコマンドの追加
	cmd.AddCommand(
		newPetsListCmd(),
		newPetsCreateCmd(),
	)

	return cmd
}

func newPetsListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get pets list",
		RunE:  runPetsListCmd,
	}
	return cmd
}

func newPetsCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <store_id>",
		Short: "Create a pet",
		RunE:  runPetsCreateCmd,
	}
	return cmd
}

// pets listの出力
func runPetsListCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.PetList(ctx)
	if err != nil {
		return errors.Wrapf(err, "PetList was failed:res = %+v", res)
	}
	fmt.Println(res)

	return nil
}

// pets listの処理
func (client *Client) PetList(ctx context.Context) (string, error) {
	subPath := fmt.Sprintf("/pets")
	httpRequest, err := client.newRequest(ctx, "GET", subPath, nil)
	if err != nil {
		log.Println("newRequest Error")
		return "error", err
	}

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		log.Println("HTTPClient Do Error")
		return "error", err
	}
	defer httpResponse.Body.Close()
	res, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Println("ReadAll Error")
		return "error", err
	}
	return string(res), nil
}

// pets createの出力
func runPetsCreateCmd(cmd *cobra.Command, args []string) error {
	// store_idが指定されているか確認
	if len(args) == 0 {
		return errors.New("store_id is required")
	}
	store_id := args[0]

	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.PetCreate(ctx, store_id)
	if err != nil {
		return errors.Wrapf(err, "PetCreate was failed:res = %+v", res)
	}
	fmt.Println(res)

	return nil
}

func (client *Client) PetCreate(ctx context.Context, store_id string) (string, error) {
	subPath := fmt.Sprintf("/pets")

	// POSTするデータ
	jsonStr := `{
    "species": "Inu",
    "name":"Pug",
    "age": 3,
    "store_id":"` + store_id + `"
	}`
	httpRequest, err := client.newRequest(ctx, "POST", subPath, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Println("create NewRequest error ")
		return "error", err
	}

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		log.Println("create HTTPClient Do Error")
		return "error", err
	}
	defer httpResponse.Body.Close()
	// レスポンスを取得し出力
	res, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Println("ReadAll Error")
		return "error", err
	}
	return string(res), nil

}
