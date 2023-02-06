package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/hyperledger/fabric-private-chaincode/application/pkg"
	"github.com/hyperledger/fabric-private-chaincode/lib"
	"github.com/spf13/cobra"
)

func init() {
	createCmd.AddCommand(createSlaCmd)
}

var createSlaCmd = &cobra.Command{
	Use:     "sla sla_json",
	Short:   "create SLA on the chaincode",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jsonFile, err := os.Open(args[0])
		if err != nil {
			panic(err)
		}
		defer jsonFile.Close()

		// read our opened jsonFile as a byte array.
		byteValue, err := io.ReadAll(jsonFile)
		if err != nil {
			panic(err)
		}

		// Verify that the data indeed fits the struct
		// before we send it to the chaincode
		var sla lib.SLA
		err = json.Unmarshal(byteValue, &sla)
		if err != nil {
			panic(err)
		}

		client := pkg.NewClient(config)
		res := client.Invoke("CreateOrUpdateContract", string(byteValue))
		fmt.Println("> " + res)
	},
}
