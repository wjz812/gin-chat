package user

import (
	"fmt"
	"ginchat/global/consts"
	"ginchat/http/api_param"
	"ginchat/http/model"
	"ginchat/pkg/response"
	"ginchat/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

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

	//token 加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := utils.MD5Encode(str)

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
