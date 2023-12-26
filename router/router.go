package router

import (
	validatorFactory "ginchat/core/validator_factory"
	"ginchat/global/config"
	"ginchat/global/consts"
	"ginchat/http/middleware"
	"ginchat/service"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	pprof.Register(r)

	//获取真实ip地址
	_ = r.SetTrustedProxies(nil)
	r.TrustedPlatform = "X-Real-IP"

	//域名访问限制
	if config.Conf.HTTPServer.AllowCrossDomain {
		r.Use(middleware.Cors())
	}

	r.PUT("/api/login", validatorFactory.Create(consts.ValidatorPrefix+"UserLogin"))
	r.GET("/api/ip", validatorFactory.Create(consts.ValidatorPrefix+"GetIp"))
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
