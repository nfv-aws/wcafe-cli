package cmd

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"fmt"
	"io/ioutil"

	"github.com/jmcvetta/napping"
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

// pets listの処理
func runPetsListCmd(cmd *cobra.Command, args []string) error {
	url := "http://" + dns + ":8080/api/v1/pets"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	return nil
}

// pets createの処理
func runPetsCreateCmd(cmd *cobra.Command, args []string) error {
	url := "http://" + dns + ":8080/api/v1/pets"
	log.Println("URL:>", url)

	// store_idが指定されているか確認
	if len(args) == 0 {
		return errors.New("store_id is required")
	}
	_, err := strconv.Atoi(args[0])

	s := napping.Session{}
	h := &http.Header{}
	h.Set("X-Custom-Header", "myvalue")
	s.Header = h

	var jsonStr = []byte(`
{
    "species": "Inu",
    "name":"Pug",
    "age": 3,
    "store_id":"` + args[0] + `"
}`)

	var data map[string]json.RawMessage
	err = json.Unmarshal(jsonStr, &data)
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
