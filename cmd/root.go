package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

var cfgFile string

var (
	// RootCmd defines root command
	RootCmd = &cobra.Command{
		Use: "wcafe",
		// SilenceErrors: true,
		// SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}
)

func Run() {
	RootCmd.Execute()
}

// Exit finishes a runnning action.
func Exit(err error, codes ...int) {
	var code int
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 2
	}
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func newDefaultClient() (*Client, error) {
	endpointURL := viper.GetString("url")
	httpClient := &http.Client{}
	userAgent := fmt.Sprintf("wcafe/%s (%s)")
	return newClient(endpointURL, httpClient, userAgent)
}
