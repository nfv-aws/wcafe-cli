package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// clerksコマンドの定義
func newClerksCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clerks",
		Short: "clerks command",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	// サブコマンドの追加
	cmd.AddCommand(
		newClerksListCmd(),
		newClerksCreateCmd(),
	)
	return cmd
}

func newClerksListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Get clerks list",
		RunE:  runClerksListCmd,
	}
	return cmd
}

type Clerks struct {
	Name string
}

func newClerksCreateCmd() *cobra.Command {

	var (
		cl = &Clerks{}
	)
	cmd := &cobra.Command{
		Use:   "create [flag]",
		Short: "Create a clerk",
		Run: func(cmd *cobra.Command, args []string) {
			runClerksCreateOpt(cmd, cl)
		},
	}
	cmd.Flags().StringVarP(&cl.Name, "name", "n", "climan", "change name")
	return cmd
}

// clerks list の出力
func runClerksListCmd(cmd *cobra.Command, args []string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body, err := client.ClerkList(ctx)
	if err != nil {
		return errors.Wrapf(err, "ClerkList was failed:body = %+v", body)
	}
	fmt.Println(body)

	return nil
}

// clerks list の処理
func (client *Client) ClerkList(ctx context.Context) (string, error) {
	subPath := fmt.Sprintf("/clerks")
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

// create時のオプション有無の判断
func runClerksCreateOpt(cmd *cobra.Command, opt *Clerks) {
	if len(opt.Name) != 0 {
		runClerksCreateCmd(cmd, opt.Name)
	} else {
		opt.Name = "climan"
		runClerksCreateCmd(cmd, opt.Name)
	}
}

// clerks createの出力
func runClerksCreateCmd(cmd *cobra.Command, optName string) error {
	client, err := newDefaultClient()
	if err != nil {
		return errors.Wrap(err, "newClient failed:")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body, err := client.ClerkCreate(ctx, optName)
	if err != nil {
		return errors.Wrapf(err, "ClerkCreate was failed:body = %+v", body)
	}
	fmt.Println(body)

	return nil
}

// clerks createの処理
func (client *Client) ClerkCreate(ctx context.Context, optName string) (string, error) {
	subPath := fmt.Sprintf("/clerks")
	// POSTするデータ
	jsonStr := `{
	"Name": "` + optName + `"
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
	// レスポンスを取得し出力
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "error", errors.Wrapf(err, "ReadAll was faild:body=%+v", body)
	}
	return string(body), nil
}
