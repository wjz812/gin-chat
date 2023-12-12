package user

import (
	"fmt"
	"ginchat/global/consts"
	"ginchat/http/api_param"
	"ginchat/http/model"
	"ginchat/pkg/response"
	token_verify "ginchat/pkg/token"
	"ginchat/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

/**
 * @description 用户登录
 * @param
 * @return
 * @date 2023/12/12 14:58:35
 * @version: 1.0.0
 * @author
 */
func Login(c *gin.Context, req api_param.UserLoginReq) {
	dbConn := model.UseDbConn().Begin()
	defer dbConn.Rollback()
	userInfo, err := model.CreateUserModelFactory(dbConn).Login(req)
	if err != nil {
		fmt.Println(err)
		response.Fail(c, consts.UserLoginFail.Code, consts.UserLoginFail.Msg, gin.H{})
		return
	}

	if !utils.ValidPassword(req.Password, consts.Salt, userInfo.Password) {
		response.Fail(c, consts.PasswordFail.Code, consts.PasswordFail.Msg, gin.H{})
		return
	}

	token, err := token_verify.GenerateToken(req.Name) // 获取token
	if err != nil {
		response.Fail(c, consts.TokenGenerateErr.Code, consts.TokenGenerateErr.Msg, gin.H{})
		return
	}

	//存储token
	err = model.CreateUserModelFactory(dbConn).Token(userInfo.Id, token)
	if err != nil {
		fmt.Println(err)
		response.Fail(c, consts.UserLoginFail.Code, consts.UserLoginFail.Msg, gin.H{})
		return
	}

	result := api_param.UserLoginResp{
		Id:       userInfo.Id,
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Identity: token,
	}

	dbConn.Commit()
	response.Success(c, result)
}

/**
 * @description 用户一览
 * @param
 * @return
 * @date 2023/12/12 14:58:53
 * @version: 1.0.0
 * @author
 */
func List(c *gin.Context, req api_param.ListUserReq) {
	dbConn := model.UseDbConn().Begin()
	defer dbConn.Rollback()

	result, err := model.CreateUserModelFactory(dbConn).List()
	if err != nil {
		response.Fail(c, consts.UserNotList.Code, consts.UserNotList.Msg, gin.H{})
		return
	}

	dbConn.Commit()
	response.Success(c, result)
}

var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	fmt.Println("SendMsg")
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return
	}

	MsgHandler(ws, c)
	defer func(ws *websocket.Conn) {
		err = ws.Close()

		if err != nil {
			return
		}
	}(ws)

}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	msg, err := utils.Subscribe(c, utils.PublishKey)
	if err != nil {
		fmt.Println(err)
	}

	tm := time.Now().Format("2006-01-02 15:04:05")

	m := fmt.Sprintf("[ws][%s]:%s", tm, msg)

	err = ws.WriteMessage(1, []byte(m))

	if err != nil {
		return
	}
}
