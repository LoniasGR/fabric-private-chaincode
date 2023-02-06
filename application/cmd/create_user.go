package cmd

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/hyperledger/fabric-private-chaincode/application/pkg"
	"github.com/spf13/cobra"
)

func init() {
	createCmd.AddCommand(createUserCmd)
}

var createUserCmd = &cobra.Command{
	Use:   "user username",
	Short: "create user on the chaincode",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			panic(fmt.Sprintf("Generation error : %s", err))
		}

		b, err := x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			panic(err)
		}

		block := &pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: b,
		}

		err = ioutil.WriteFile(name, pem.EncodeToMemory(block), 0600)
		if err != nil {
			panic(err)
		}

		b, err = x509.MarshalPKIXPublicKey(publicKey)
		if err != nil {
			panic(err)
		}

		block = &pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: b,
		}

		pubKeyPEM := strings.Trim(strings.ReplaceAll(string(pem.EncodeToMemory(block)), "\n", ""), " ")
		fmt.Println(pubKeyPEM)

		client := pkg.NewClient(config)
		res := client.Invoke("CreateUser", name, pubKeyPEM, "500")
		fmt.Println("> " + res)
	},
}
