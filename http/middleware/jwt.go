package middleware

import (
	"fmt"
	"ginchat/global/config"
	"ginchat/global/consts"
	"ginchat/pkg/response"
	token_verify "ginchat/pkg/token"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	fmt.Println("jwt")
	return func(c *gin.Context) {
		resp := consts.CurdOK
		token := c.GetHeader(config.Conf.Token.BindContextKeyName)
		fmt.Println("jwt token:", token)
		if token == "" {
			resp = consts.TokenMissing
		} else if claims, err := token_verify.ParseToken(token); err != nil {
			resp = consts.TokenInvalid
		} else {
			fmt.Println(claims)
		}

		if resp != consts.CurdOK {
			response.Fail(c, resp.Code, resp.Msg, gin.H{})
			c.Abort()
			return
		}

		c.Next()
	}
}
