package router

import (
	"fmt"
	validatorFactory "ginchat/core/validator_factory"
	"ginchat/global/consts"
	"ginchat/http/middleware"
	"ginchat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.PUT("/api/login", validatorFactory.Create(consts.ValidatorPrefix+"UserLogin"))
	fmt.Println("=================")
	r.Use(middleware.JWT())

	vApi := r.Group("/api/")
	{
		user := vApi.Group("user")
		{
			user.GET("list", validatorFactory.Create(consts.ValidatorPrefix+"UserList"))
			user.PUT("new", service.CreateUser)
			user.DELETE("del", service.DeleteUser)
			user.POST("edit", service.UpdateUser)
		}
	}

	// r.GET("/api/sendmsg", validatorFactory.Create(consts.ValidatorPrefix+"SendMsg"))

	return r
}
