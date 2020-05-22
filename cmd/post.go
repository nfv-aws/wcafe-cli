package cmd

import (
	"encoding/json"
	"github.com/jmcvetta/napping"

	"github.com/spf13/cobra"
	"log"
	math_rand "math/rand"
	"net/http"
	"strconv"
	"time"
	// 以下、**TODO WCAF-123　**　で利用予定
	// "crypto/rand"
	// "encoding/binary"
	// "github.com/nfv-aws/wcafe-api-controller/db"
	// "github.com/nfv-aws/wcafe-api-controller/entity"
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
		// **TODO WCAF-123**
		// newPostStoresCmd(),
		// newPostPetsCmd(),
		newPostUsersCmd(),
	)

	return cmd
}

// **TODO WCAF-123**
// func newPostStoresCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "stores",
// 		Short: "Post a store",
// 		RunE:  runPostStoresCmd,
// 	}
// 	return cmd
// }

// func newPostPetsCmd() *cobra.Command {
// 	cmd := &cobra.Command{
// 		Use:   "pets",
// 		Short: "Post a pet",
// 		RunE:  runPostPetsCmd,
// 	}
// 	return cmd
// }

func newPostUsersCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "users",
		Short: "Post a user",
		RunE:  runPostUsersCmd,
	}
	return cmd
}

func runPostUsersCmd(cmd *cobra.Command, args []string) error {
	url := endpoint + "/api/v1/users"
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

// **TODO WCAF-123**
// // ランダムな文字列の生成
// func random() string {
// 	var n uint64
// 	binary.Read(rand.Reader, binary.LittleEndian, &n)
// 	return strconv.FormatUint(n, 36)
// }

// func runPostStoresCmd(cmd *cobra.Command, args []string) error {
// 	url := endpoint + "/api/v1/stores"
// 	log.Println("URL:>", url)

// 	s := napping.Session{}
// 	h := &http.Header{}
// 	h.Set("X-Custom-Header", "myvalue")
// 	s.Header = h

// 	var jsonStr = []byte(`
// {
//     "name": "` + random() + `",
//     "tag":"CLI",
//     "address":"Okinawa"
// }`)

// 	var data map[string]json.RawMessage
// 	err := json.Unmarshal(jsonStr, &data)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	resp, err := s.Post(url, &data, nil, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("response Status:", resp.Status())
// 	log.Println("response Headers:", resp.HttpResponse().Header)
// 	log.Println("response Body:", resp.RawText())

// 	return nil
// }

// func runPostPetsCmd(cmd *cobra.Command, args []string) error {
// 	url := endpoint + "/api/v1/pets"
// 	log.Println("URL:>", url)

// 	s := napping.Session{}
// 	h := &http.Header{}
// 	h.Set("X-Custom-Header", "myvalue")
// 	s.Header = h

// 	var store []entity.Store
// 	db := db.GetDB()
// 	db.Find(&store)

// 	var jsonStr = []byte(`
// {
//     "species": "Inu",
//     "name":"Pug",
//     "age": 3,
//     "store_id":"` + store[0].Id + `"
// }`)

// 	var data map[string]json.RawMessage
// 	err := json.Unmarshal(jsonStr, &data)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	resp, err := s.Post(url, &data, nil, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("response Status:", resp.Status())
// 	log.Println("response Headers:", resp.HttpResponse().Header)
// 	log.Println("response Body:", resp.RawText())

// 	return nil
// }
