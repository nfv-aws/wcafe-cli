package cmd

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// コマンドの追加
func init() {
	RootCmd.AddCommand(newStoresCmd())
	RootCmd.AddCommand(newPetsCmd())
	RootCmd.AddCommand(newusersCmd())
}

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
	)
	return cmd
}

func newStoresListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get stores list",
		RunE:  RunStoresListCmd,
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

// 以下、stores listの処理
// stores list の出力
func RunStoresListCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.StoreList(ctx)
	if err != nil {
		return errors.Wrapf(err, "StoreList was failed:res = %+v", res)
	}
	fmt.Println(res)

	return nil
}

// GET storesの呼び出し
func (client *Client) StoreList(ctx context.Context) (string, error) {
	subPath := fmt.Sprintf("/stores")
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

// 以下、createの処理
// ランダムな文字列の生成
func random() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

// stores createの処理
func runStoresCreateCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.StoreCreate(ctx)
	if err != nil {
		return errors.Wrapf(err, "StoreCreate was failed:res = %+v", res)
	}
	fmt.Println(res)

	return nil
}

// POST storesの呼び出し
func (client *Client) StoreCreate(ctx context.Context) (string, error) {
	subPath := fmt.Sprintf("/stores")

	// POSTするデータ
	jsonStr := `{
	"name": "` + random() + `",
    "tag":"CLI",
    "address":"Okinawa"
	}`
	httpRequest, err := client.newRequest(ctx, "POST", subPath, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		log.Println("create NewRequest error ")
		return "error", err
	}

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		log.Println("create HTTPClient Do Error")
		log.Println(httpResponse)
		log.Println(httpRequest)
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
