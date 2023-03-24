package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-private-chaincode/api/globals"
	"github.com/hyperledger/fabric-private-chaincode/api/ledger"
	"github.com/hyperledger/fabric-private-chaincode/lib"
)

func GetSingleSLA(c *gin.Context) {
	var id string
	if err := c.ShouldBindUri(&id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	username := session.Get(globals.Userkey)

	user := ledger.GetUser(username.(string))

	if !slaInUserContracts(user, id) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User does not have access to this contract"})
		return
	}

	asset, err := ledger.GetSLA(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": asset})
	return
}

func CreateSLA(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get(globals.Userkey)

	var sla lib.SLA

	if err := c.BindJSON(&sla); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if sla.Details.Provider.Name != username {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not the provider of the SLA"})
	}

	err := ledger.CreateSLA(sla)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func slaInUserContracts(user lib.User, slaId string) bool {
	clientList := strings.Split(user.ClientOf, ",")
	// Slice of size 1 means that the delimiter was not found in the string
	if len(clientList) != 1 {
		for _, sla := range clientList {
			if sla == slaId {
				return true
			}
		}
	}

	providerList := strings.Split(user.ProviderOf, ",")
	// Slice of size 1 means that the delimiter was not found in the string
	if len(providerList) != 1 {
		for _, sla := range providerList {
			if sla == slaId {
				return true
			}
		}
	}
	return false
}
