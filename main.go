package main

import (
	"go-swagger-tutorial/controllers"
	_ "go-swagger-tutorial/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is just a sample server
// @terms of service https://gaming.innv8.ke/tos
//
// @contact.name Rapando
// @contact.url https://innv8.ke
//
// @license.name Apache 2.0
// @license.url https://gaming.innv8.ke/license
//
// @host localhost:5000
// @BasePath /
func main() {
	var base controllers.Base
	base.Init()
}
