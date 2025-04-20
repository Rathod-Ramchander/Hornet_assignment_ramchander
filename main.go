package main

import (
	"beneficiary-tracer/config"
	"beneficiary-tracer/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	router := gin.Default()
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.GET("/beneficiary", handlers.GetBeneficiaries)
	router.Run(":8082")
}
