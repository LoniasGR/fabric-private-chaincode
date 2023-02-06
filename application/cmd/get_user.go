package cmd

import (
	"fmt"

	"github.com/hyperledger/fabric-private-chaincode/application/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getUserCmd)
}

var getUserCmd = &cobra.Command{
	Use:   "user name ...",
	Short: "query FPC Chaincode for user",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		client := pkg.NewClient(config)
		res := client.Query("ReadUser", name)
		fmt.Println("> " + res)
	},
}
