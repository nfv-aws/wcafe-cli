package cmd

import (
	"github.com/nfv-aws/wcafe-cli/config"
	"github.com/spf13/cobra"
)

var (
	// ルートコマンドの設定
	RootCmd = &cobra.Command{
		Use: "wcafe",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
)

// Private DNS
var dns string

// コンフィグの呼び出し
func init() {
	config.Configure()
	dns = config.C.VM.Private_dns
}
