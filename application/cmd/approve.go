package cmd

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"os"

	"github.com/hyperledger/fabric-private-chaincode/application/pkg"
	"github.com/hyperledger/fabric-private-chaincode/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(approveCmd)
}

var approveCmd = &cobra.Command{
	Use:   "approve sla_json username priv_key",
	Short: "approve SLA on the chaincode",
	Args:  cobra.MinimumNArgs(3),
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

		privKeyFile, err := os.Open(args[2])
		if err != nil {
			panic(err)
		}
		defer privKeyFile.Close()

		privKeyBytes, err := io.ReadAll(privKeyFile)
		if err != nil {
			panic(err)
		}

		block, _ := pem.Decode(privKeyBytes)
		if block == nil {
			panic("Key file could not be read")
		}

		if block.Type == "CERTIFICATE" {
			panic("Public key provided instead of private key")
		}

		privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			panic(err)
		}

		key, ok := privateKey.(ed25519.PrivateKey)
		if !ok {
			panic("Wrong key type")
		}

		// We need to marshal the bytes, to make sure it's consistent
		slaBytes, err := json.Marshal(sla)
		if err != nil {
			panic(err)
		}
		signature := ed25519.Sign(key, slaBytes)
		client := pkg.NewClient(config)
		res := client.Invoke("Approve", sla.ID, args[1], hex.EncodeToString(signature))
		fmt.Println("> " + res)
	},
}
