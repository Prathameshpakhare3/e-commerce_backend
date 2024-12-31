package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prathamesh/login_jwt_project/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	routes.SetUpRouters(router)

	log.Println("Server is starting on port 9090...")

	log.Fatal(router.Run(":9090"))
}
