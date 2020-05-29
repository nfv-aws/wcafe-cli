package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	math_rand "math/rand"
	"strconv"
	"time"

	//	"github.com/jmcvetta/napping"
	"github.com/pkg/errors"
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
		RunE:  RunUsersListCmd,
	}
	return cmd
}

func newUsersCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a user",
		RunE:  RunUsersCreateCmd,
	}
	return cmd
}

// users listの出力
func RunUsersListCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.UserList(ctx)
	if err != nil {
		return errors.Wrapf(err, "StoreList was failed:res = %+v", res)
	}
	fmt.Println(res)

	return nil
}

// users listの処理
func (client *Client) UserList(ctx context.Context) (string, error) {
	subPath := fmt.Sprintf("/users")
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

// users createの出力
func RunUsersCreateCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := client.UserCreate(ctx)
	if err != nil {
		return errors.Wrapf(err, "UserCreate was failed:res = %+v", res)
	}
	fmt.Println(res)

	return nil
}

// users createの処理
func (client *Client) UserCreate(ctx context.Context) (string, error) {
	subPath := fmt.Sprintf("/users")
	math_rand.Seed(time.Now().UnixNano())
	random_num := math_rand.Intn(10000)
	body := `{"number":` + strconv.Itoa(random_num) + `}`
	httpRequest, err := client.newRequest(ctx, "POST", subPath, bytes.NewBuffer([]byte(body)))
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
