package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-private-chaincode/api/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/", controllers.GetIndex)
	g.POST("/login", controllers.Login)
}
