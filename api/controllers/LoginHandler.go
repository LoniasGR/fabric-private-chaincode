package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-private-chaincode/api/ledger"
	"github.com/hyperledger/fabric-private-chaincode/api/models"
	"github.com/hyperledger/fabric-private-chaincode/api/utils"
)

// login ensures the user is logged in
func Login(c *gin.Context) {

	var userAPI models.User

	if err := c.BindJSON(&userAPI); err != nil {
		c.JSON(401, gin.H{"error": err})
		return
	}

	if userAPI.Mnemonic == "" {
		c.JSON(401, gin.H{"error": "Mnemonic is missing"})
		return
	}

	userLedger := ledger.GetUser(userAPI.Name)

	if userLedger == (models.ContractUser{}) {
		c.JSON(404, gin.H{"error": "User does not exist"})
		return
	}

	match, err := utils.PubKeyMatchesMnemonic(userAPI.Mnemonic, "password", userLedger.PubKey)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	if !match {
		c.AbortWithError(404, fmt.Errorf("Mnemonic and public key do not match"))
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("fabricAuth", userAPI.Name, 1000*60*60, "/", "localhost", false, true)
	return
}
