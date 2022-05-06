package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "dsbd wg",
	Short: "dashboard wg",
	Long:  "Source Code: https://github.com/doornoc/dsbd-wg",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
