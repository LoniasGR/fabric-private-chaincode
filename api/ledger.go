package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-private-chaincode/api/models"
	"github.com/hyperledger/fabric-private-chaincode/api/pkg"
	"github.com/hyperledger/fabric-private-chaincode/api/utils"
	"github.com/tyler-smith/go-bip32"
)

func InitLedger() {
	// client := pkg.NewClient(config)
	// _, err := client.Invoke("InitLedger")
	// if err != nil {
	// 	if err.Error() == "init has already ran" {
	// 		log.Println("Init Ledger has already run. Continuing.")
	// 		return
	// 	}
	// 	log.Fatalln(err)
	// 	return
	// }
	users := [10]string{"Tomoko", "Brad", "Jin Soo", "Max", "Adriana", "Michel", "Mario", "Anton", "Marek", "George"}
	for _, u := range users {
		mnemonic, err := utils.CreateMnemonic()
		if err != nil {
			log.Fatalln(err)
		}
		keysSerialized, err := utils.CreateMasterKey(mnemonic, "password")
		if err != nil {
			log.Fatalln(err)
		}
		keys, _ := bip32.Deserialize(keysSerialized)
		CreateUser(u, keys.PublicKey().B58Serialize())
	}

}

func GetUser(name string) models.ContractUser {
	client := pkg.NewClient(config)
	res, err := client.Query("ReadUser", name)
	if err != nil {
		log.Fatalln(err)
	}

	var user models.ContractUser
	json.Unmarshal([]byte(res), &user)
	return user
}

func CreateUser(name, publicKey string) {
	fmt.Println(name, publicKey)
	client := pkg.NewClient(config)
	res, err := client.Invoke("CreateUser", name, publicKey, "500")
	if err != nil {
		log.Fatalln(err)
	}

}
