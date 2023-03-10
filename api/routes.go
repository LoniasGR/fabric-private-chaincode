package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-private-chaincode/api/models"
	"github.com/hyperledger/fabric-private-chaincode/api/utils"
)

// GetName returns the name of this API
func GetRoot(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Response{Name: "fabric-private-chaincode"})
}

// login ensures the user is logged in
func Login(c *gin.Context) {

	var userAPI models.User

	if err := c.BindJSON(&userAPI); err != nil {
		c.JSON(401, gin.H{"error": err})
	}

	if userAPI.Mnemonic == "" {
		c.JSON(401, gin.H{"error": "Mnemonic is missing"})
	}

	userLedger := GetUser(userAPI.Name)
	if userLedger == (models.ContractUser{}) {
		c.JSON(404, gin.H{"error": "User does not exist"})
	}

	c.JSON(200, gin.H{"msg": userLedger})
	match, err := utils.PubKeyMatchesMnemonic(userAPI.Mnemonic, "", []byte(userLedger.PubKey))
	if err != nil {
		c.AbortWithError(401, err)
	}

	if !match {
		c.AbortWithError(404, fmt.Errorf("Mnemonic and public key do not match"))
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("fabricAuth", userAPI.Name, 1000*60*60, "/", "localhost", false, true)
}
