/*
Copyright IBM Corp. All Rights Reserved.
Copyright 2020 Intel Corporation

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-private-chaincode/api/pkg"
)

var (
	config *pkg.Config
)

func main() {
	initConfig()
	InitLedger()

	router := gin.Default()
	router.GET("/", GetRoot)
	router.GET("/login", Login)

	router.Run(":8000")
}
