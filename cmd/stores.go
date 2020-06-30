package cmd

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// storesコマンドの定義
func newStoresCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stores",
		Short: "stores command",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	// サブコマンドの追加
	cmd.AddCommand(
		newStoresListCmd(),
		newStoresCreateCmd(),
		newStoresDeleteCmd(),
	)
	return cmd
}

func newStoresListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get stores list",
		RunE:  runStoresListCmd,
	}
	return cmd
}

func newStoresCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a store",
		RunE:  runStoresCreateCmd,
	}
	return cmd
}

func newStoresDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <store_id>",
		Short: "Delete a store",
		RunE:  runStoresDeleteCmd,
	}
	return cmd
}

// stores list の出力
func runStoresListCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body, err := client.StoreList(ctx)
	if err != nil {
		return errors.Wrapf(err, "StoreList was failed:body = %+v", body)
	}
	fmt.Println(body)

	return nil
}

// stores list の処理
func (client *Client) StoreList(ctx context.Context) (string, error) {
	subPath := fmt.Sprintf("/stores")
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

// ランダムな文字列の生成
func random() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

// stores createの出力
func runStoresCreateCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body, err := client.StoreCreate(ctx)
	if err != nil {
		return errors.Wrapf(err, "StoreCreate was failed:body = %+v", body)
	}
	fmt.Println(body)

	return nil
}

// stores createの処理
func (client *Client) StoreCreate(ctx context.Context) (string, error) {
	subPath := fmt.Sprintf("/stores")

	// POSTするデータ
	jsonStr := `{
	"name": "` + random() + `",
    "tag":"CLI",
    "address":"Okinawa"
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

// stores deleteの出力
func runStoresDeleteCmd(cmd *cobra.Command, args []string) error {
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

	body, err := client.StoreDelete(ctx, store_id)
	if err != nil {
		return errors.Wrapf(err, "StoreDelete was failed:body = %+v", body)
	}
	fmt.Println(body)

	return nil
}

// stores deleteの処理
func (client *Client) StoreDelete(ctx context.Context, store_id string) (string, error) {
	subPath := fmt.Sprintf("/stores/" + store_id)
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
