package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
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
		newPetsDeleteCmd(),
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

func newPetsDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <pets_id>",
		Short: "Delete a pet",
		RunE:  runPetsDeleteCmd,
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

	body, err := client.PetList(ctx)
	if err != nil {
		return errors.Wrapf(err, "PetList was failed:body = %+v", body)
	}
	fmt.Println(body)

	return nil
}

// pets listの処理
func (client *Client) PetList(ctx context.Context) (string, error) {
	subPath := fmt.Sprintf("/pets")
	req, err := client.newRequest(ctx, "GET", subPath, nil)
	if err != nil {
		return "error", errors.Wrapf(err, "newRequest was faild:req= %+v", req)
	}

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return "error", errors.Wrapf(err, "HTTPClient Do was faild:res=%+v", res)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "error", errors.Wrapf(err, "ReadAll was faild:body=%+v", body)
	}
	return string(body), nil
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

	body, err := client.PetCreate(ctx, store_id)
	if err != nil {
		return errors.Wrapf(err, "PetCreate was failed:body = %+v", body)
	}
	fmt.Println(body)

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
	req, err := client.newRequest(ctx, "POST", subPath, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return "error", errors.Wrapf(err, "NewRequest was failed:req = %+v", req)
	}

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return "error", errors.Wrapf(err, "HTTPClient Do was faild:res=%+v", res)
	}
	defer res.Body.Close()
	// レスポンスを取得し出力
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "error", errors.Wrapf(err, "ReadAll was faild:body=%+v", body)
	}
	return string(body), nil

}

// pets deleteの出力
func runPetsDeleteCmd(cmd *cobra.Command, args []string) error {
	// pet_idが指定されているか確認
	if len(args) == 0 {
		return errors.New("pet_id is required")
	}
	pet_id := args[0]

	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body, err := client.PetDelete(ctx, pet_id)
	if err != nil {
		return errors.Wrapf(err, "PetDelete was failed:body = %+v", body)
	}
	fmt.Println(body)

	return nil
}

// pets deleteの処理
func (client *Client) PetDelete(ctx context.Context, pet_id string) (string, error) {
	subPath := fmt.Sprintf("/pets/" + pet_id)
	req, err := client.newRequest(ctx, "DELETE", subPath, nil)
	if err != nil {
		return "error", errors.Wrapf(err, "newRequest was faild:req= %+v", req)
	}

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return "error", errors.Wrapf(err, "HTTPClient Do was faild:res=%+v", res)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "error", errors.Wrapf(err, "ReadAll was faild:body=%+v", body)
	}
	return string(body), nil
}
