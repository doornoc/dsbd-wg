package cmd

import (
	"fmt"
	"github.com/doornoc/dsbd-wg/pkg/api"
	"github.com/doornoc/dsbd-wg/pkg/core/config"
	"github.com/doornoc/dsbd-wg/pkg/core/peer"
	"github.com/spf13/cobra"
	"log"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start application",
	Long:  `start application`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		config.DbPath, err = cmd.Flags().GetString("database")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		if config.DbPath == "" {
			log.Fatalf("[error] invalid database path")
		}

		fmt.Println("------Application Start(User)------")

		err = peer.WgInit()
		if err != nil {
			log.Println(err)
			log.Fatalf("[error] Wireguard init process")
		}

		api.RestAPI()
		fmt.Println("------End------")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringP("database", "p", "", "database path")
}
