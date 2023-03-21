/*
Copyright IBM Corp. All Rights Reserved.
Copyright 2020 Intel Corporation

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-private-chaincode/api/globals"
	"github.com/hyperledger/fabric-private-chaincode/api/ledger"
	"github.com/hyperledger/fabric-private-chaincode/api/routes"
)

func main() {
	initConfig()
	ledger.InitLedger()

	router := gin.Default()

	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/")

	routes.PublicRoutes(public)

	router.Run(":8000")
}
