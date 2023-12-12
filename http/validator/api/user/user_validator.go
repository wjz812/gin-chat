package user

import (
	"fmt"
	"ginchat/http/api_param"
	"ginchat/http/controller/user"
	"ginchat/pkg/response"

	"github.com/gin-gonic/gin"
)

// 用户登录
type UserLogin struct{}

// 用户注册
type UserRegister struct{}

// 用户列表
type UserList struct{}

type SendMsg struct{}

func (n UserLogin) CheckParams(c *gin.Context) {
	params := api_param.UserLoginReq{}

	if err := c.ShouldBindJSON(&params); err != nil {
		response.ValidatorError(c, err)
		return
	}

	user.Login(c, params)
}

func (n UserRegister) CheckParams(c *gin.Context) {
	params := api_param.CreateUserReq{}

	if err := c.ShouldBindJSON(&params); err != nil {
		response.ValidatorError(c, err)
		return
	}

	//user.
}

func (n SendMsg) CheckParams(c *gin.Context) {
	// params := api_param.UserLoginReq{}

	// if err := c.ShouldBindJSON(&params); err != nil {
	// 	response.ValidatorError(c, err)
	// 	return
	// }

	user.SendMsg(c)
}

func (n UserList) CheckParams(c *gin.Context) {
	params := api_param.ListUserReq{}
	fmt.Println("====> user List check params")
	// if err := c.ShouldBindJSON(&params); err != nil {
	// 	response.ValidatorError(c, err)
	// 	fmt.Println("====> err:", err)
	// 	return
	// }

	user.List(c, params)
}
