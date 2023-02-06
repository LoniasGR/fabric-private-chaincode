package cmd

import (
	"fmt"

	"github.com/hyperledger/fabric-private-chaincode/application/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getSlaCmd)
}

var getSlaCmd = &cobra.Command{
	Use:   "sla id ...",
	Short: "query FPC Chaincode for sla",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		client := pkg.NewClient(config)
		res := client.Query("ReadContract", id)
		fmt.Println("> " + res)
	},
}
