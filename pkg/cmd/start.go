package cmd

import (
	"fmt"
	"github.com/doornoc/dsbd-wg/pkg/api"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start application",
	Long:  `start application`,
	Run: func(cmd *cobra.Command, args []string) {
		//confPath, err := cmd.Flags().GetString("config")
		//if err != nil {
		//	log.Fatalf("could not greet: %v", err)
		//}
		//if config.GetConfig(confPath) != nil {
		//	log.Fatalf("error config process |%v", err)
		//}

		fmt.Println("------Application Start(User)------")

		api.RestAPI()
		fmt.Println("------End------")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringP("config", "c", "", "config path")
}
