package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	math_rand "math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/jmcvetta/napping"
	"github.com/spf13/cobra"
)

// usersコマンドの定義
func newusersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "users command",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	// サブコマンドの追加
	cmd.AddCommand(
		newUsersListCmd(),
		newUsersCreateCmd(),
	)
	return cmd
}

func newUsersListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get pets list",
		RunE:  runUsersListCmd,
	}
	return cmd
}

func newUsersCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a user",
		RunE:  runUsersCreateCmd,
	}
	return cmd
}

// users listの処理
func runUsersListCmd(cmd *cobra.Command, args []string) error {
	url := "http://" + dns + ":8080/api/v1/users"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	resp, _ := client.Do(req)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))

	return nil
}

// users createの処理
func runUsersCreateCmd(cmd *cobra.Command, args []string) error {
	url := "http://" + dns + ":8080/api/v1/users"
	log.Println("URL:>", url)

	s := napping.Session{}
	h := &http.Header{}
	h.Set("X-Custom-Header", "myvalue")
	s.Header = h

	math_rand.Seed(time.Now().UnixNano())
	random_num := math_rand.Intn(10000)

	var jsonStr = []byte(`
{
    "number":` + strconv.Itoa(random_num) + `,
    "name":"gogogo"
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
