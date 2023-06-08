package main

import (
	"time"

	"github.com/gin-gonic/gin"

	"onlyfounds/api/route"
	"onlyfounds/config"
)

func main() {
	config.LoadEnvVariables()
	db := config.InitDB()
	gin := gin.Default()
	// gin.Static("/", "./public")
	// gin.LoadHTMLGlob("templates/*")
	timeout := time.Duration(2) * time.Second
	route.Setup(timeout, db, gin)

	gin.Run()
}
