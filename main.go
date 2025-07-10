package main

import (
	"golang/golang-rest-api/pkg/database"
	"golang/golang-rest-api/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	router := gin.Default()
	routes.ListRoutes(router)
}
