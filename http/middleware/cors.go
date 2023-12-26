package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Set("Access-Control-Allow-Origin", "*") //允许任何来源的页面访问
		context.Header("Access-Control-Allow-Methods", "*")             //允许所有http方法进行跨域请求
		context.Header("Access-Control-Allow-Headers", "*")             //允许所有http 头部进行跨域请求
		//响应中可以被暴露给JavaScript的头部信息
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		context.Header("Access-Control-Max-Age", "172800")          //预检请求的缓存时间为172800秒（2天）
		context.Header("Access-Control-Allow-Credentials", "false") //不允许发送凭据（cookies、HTTP认证等）到其他域
		context.Header("content-type", "application/json")          //设置了响应的内容类型为JSON
		//Release all option pre-requests
		if context.Request.Method == http.MethodOptions {
			context.JSON(http.StatusOK, "Options Request!")
		}
		context.Next() //将控制权传递给下一个中间件或路由处理程序
	}
}
