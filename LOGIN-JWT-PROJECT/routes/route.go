package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prathamesh/login_jwt_project/controllers"
)

func SetUpRouters(router *gin.Engine) {
	router.POST("/signupuser", controllers.SignUpUser)
	//router.POST("/loginuser", controllers.LoginUser)
}
