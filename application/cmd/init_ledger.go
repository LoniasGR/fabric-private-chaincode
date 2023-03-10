/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package cmd

import (
	"fmt"

	"github.com/hyperledger/fabric-private-chaincode/application/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initLedgerCmd)
}

var initLedgerCmd = &cobra.Command{
	Use:   "initl",
	Short: "initialize ledger",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		client := pkg.NewClient(config)
		res := client.Invoke("InitLedger")
		fmt.Println(">" + res)
	},
}
