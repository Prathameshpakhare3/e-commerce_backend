package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prathamesh/ecommerce-project/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/login", controllers.LoginUser())
	incomingRoutes.POST("/users/signup", controllers.SignupUser())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search",controllers.SearchProductByQuery())
}
