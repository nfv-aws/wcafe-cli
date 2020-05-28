package cmd

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jmcvetta/napping"
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

// stores listの処理
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

	stores := res.Store
	fmt.Printf(
		"id: %s, name: %s, tag: %s, address: %s, strong_point:%s, created_time:%s, updated_time:%s\n",
		stores.Id, stores.Name, stores.Tag, stores.Address, stores.StrongPoint, stores.CreatedAt, stores.UpdatedAt,
	)

	return nil
}

// stores listの処理
func (client *Client) StoreList(ctx context.Context) (*StoreShowResponse, error) {
	subPath := fmt.Sprintf("/stores")
	httpRequest, err := client.newRequest(ctx, "GET", subPath, nil)
	if err != nil {
		return nil, err
	}

	httpResponse, err := client.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	var apiResponse StoreShowResponse
	if err := decodeBody(httpResponse, &apiResponse); err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

type StoreShowResponse struct {
	Store Store `json:"stores"`
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
	url := "http://" + dns + ":8080/api/v1/stores"
	log.Println("URL:>", url)

	s := napping.Session{}
	h := &http.Header{}
	h.Set("X-Custom-Header", "myvalue")
	s.Header = h

	var jsonStr = []byte(`
{
    "name": "` + random() + `",
    "tag":"CLI",
    "address":"Okinawa"
}`)

	var data map[string]json.RawMessage
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		log.Println(err)
	}

	resp, err := s.Post(url, &data, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("response Status:", resp.Status())
	log.Println("response Headers:", resp.HttpResponse().Header)
	log.Println("response Body:", resp.RawText())

	return nil
}
