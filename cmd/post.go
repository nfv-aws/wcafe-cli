package cmd

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"log"
	math_rand "math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/jmcvetta/napping"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// postコマンドの定義
func newPostCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post",
		Short: "post  a resource",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	// サブコマンドの追加
	cmd.AddCommand(
		newPostStoresCmd(),
		newPostPetsCmd(),
		newPostUsersCmd(),
	)

	return cmd
}

func newPostStoresCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stores",
		Short: "Post a store",
		RunE:  runPostStoresCmd,
	}
	return cmd
}

func newPostPetsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pets <store_id>",
		Short: "Post a pet",
		RunE:  runPostPetsCmd,
	}
	return cmd
}

func newPostUsersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "Post a user",
		RunE:  runPostUsersCmd,
	}
	return cmd
}

// // ランダムな文字列の生成
func random() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func runPostStoresCmd(cmd *cobra.Command, args []string) error {
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

func runPostPetsCmd(cmd *cobra.Command, args []string) error {
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

func runPostUsersCmd(cmd *cobra.Command, args []string) error {
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
