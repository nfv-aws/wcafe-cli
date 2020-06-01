package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	math_rand "math/rand"
	"strconv"
	"time"

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

	body, err := client.UserList(ctx)
	if err != nil {
		return errors.Wrapf(err, "StoreList was failed:body = %+v", body)
	}
	fmt.Println(body)

	return nil
}

// users listの処理
func (client *Client) UserList(ctx context.Context) (string, error) {
	subPath := fmt.Sprintf("/users")
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

// users createの出力
func RunUsersCreateCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body, err := client.UserCreate(ctx)
	if err != nil {
		return errors.Wrapf(err, "UserCreate was failed:body = %+v", body)
	}
	fmt.Println(body)

	return nil
}

// users createの処理
func (client *Client) UserCreate(ctx context.Context) (string, error) {
	subPath := fmt.Sprintf("/users")
	math_rand.Seed(time.Now().UnixNano())
	random_num := math_rand.Intn(10000)
	jsonStr := `{
		"number":` + strconv.Itoa(random_num) + `
		
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "error", errors.Wrapf(err, "ReadAll was faild:body=%+v", body)
	}

	return string(body), nil
}
