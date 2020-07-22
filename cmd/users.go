package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	math_rand "math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// usersコマンドの定義
func newUsersCmd() *cobra.Command {
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
		newUsersUpdateCmd(),
		newUsersDeleteCmd(),
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

func newUsersUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <user_id>",
		Short: "Update a user",
		RunE:  runUsersUpdateCmd,
	}
	return cmd
}

func newUsersDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <user_id>",
		Short: "Delete a user",
		RunE:  runUsersDeleteCmd,
	}
	return cmd
}

// users listの出力
func runUsersListCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body, err := client.UserList(ctx)
	if err != nil {
		return errors.Wrapf(err, "UserList was failed:body = %+v", body)
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
func runUsersCreateCmd(cmd *cobra.Command, args []string) error {
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

// users updateの出力
func runUsersUpdateCmd(cmd *cobra.Command, args []string) error {
	// user_idが指定されているか確認
	if len(args) == 0 {
		return errors.New("user_id is required")
	}
	user_id := args[0]

	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body, err := client.UserUpdate(ctx, user_id)
	if err != nil {
		return errors.Wrapf(err, "UserUpdate was failed:body = %+v", body)
	}
	fmt.Println(body)

	return nil
}

func (client *Client) UserUpdate(ctx context.Context, user_id string) (string, error) {
	subPath := fmt.Sprintf("/users/" + user_id)

	// idが存在するか確認
	get_req, err := client.newRequest(ctx, "GET", subPath, nil)
	if err != nil {
		return "error", errors.Wrapf(err, "newRequest was faild:get_req= %+v", get_req)
	}

	get_res, err := client.HTTPClient.Do(get_req)
	if err != nil {
		return "error", errors.Wrapf(err, "HTTPClient Do was faild:get_res=%+v", get_res)
	}
	defer get_res.Body.Close()
	data, err := ioutil.ReadAll(get_res.Body)

	// POSTするデータ
	jsonStr := `{
    "name": "Hinata",
    "address": "Yokohama",
    "email": "test@example.com"
	}`
	// idが存在する場合はデータを削除
	if string(data) != "" {

		req, err := client.newRequest(ctx, "PATCH", subPath, bytes.NewBuffer([]byte(jsonStr)))
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
		if string(body) != "" {
			return string(body), nil
		} else {
			fmt.Println(http.StatusBadRequest)
			return "BadRequest", nil
		}
	} else {
		fmt.Println(http.StatusNotFound)
		return "NotFound", nil
	}
}

// users deleteの出力
func runUsersDeleteCmd(cmd *cobra.Command, args []string) error {
	// user_idが指定されているか確認
	if len(args) == 0 {
		return errors.New("user_id is required")
	}
	user_id := args[0]

	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body, err := client.UserDelete(ctx, user_id)
	if err != nil {
		return errors.Wrapf(err, "UserDelete was failed:body = %+v", body)
	}
	fmt.Println(body)

	return nil
}

// users deleteの処理
func (client *Client) UserDelete(ctx context.Context, user_id string) (string, error) {
	subPath := fmt.Sprintf("/users/" + user_id)

	// idが存在するか確認
	get_req, err := client.newRequest(ctx, "GET", subPath, nil)
	if err != nil {
		return "error", errors.Wrapf(err, "newRequest was faild:get_req= %+v", get_req)
	}
	get_res, err := client.HTTPClient.Do(get_req)
	if err != nil {
		return "error", errors.Wrapf(err, "HTTPClient Do was faild:get_res=%+v", get_res)
	}
	defer get_res.Body.Close()
	data, err := ioutil.ReadAll(get_res.Body)

	// idが存在する場合はデータを削除
	if string(data) != "" {
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
		fmt.Println("user delete success")
		return string(body), nil
	} else {
		fmt.Println(http.StatusNotFound)
		return "NotFound", nil
	}
}
