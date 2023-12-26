package base

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetIp struct{}

func (n GetIp) CheckParams(c *gin.Context) {

	clientIP := c.ClientIP()
	fmt.Println("clientIP :", clientIP)

}
