package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"log"
	"strconv"

	"github.com/jmcvetta/napping"
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

// stores listの処理
func runStoresListCmd(cmd *cobra.Command, args []string) error {
	url := "http://" + dns + ":8080/api/v1/stores"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	return nil
}

// // ランダムな文字列の生成
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
