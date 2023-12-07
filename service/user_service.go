package service

import (
	"fmt"
	"ginchat/global/consts"
	"ginchat/http/api_param"
	"ginchat/http/model"
	"ginchat/pkg/response"
	"ginchat/utils"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	dbConn := model.UseDbConn().Begin()
	defer dbConn.Rollback()
	datas, err := model.CreateUserModelFactory(dbConn).List()
	if err != nil {
		response.Fail(c, consts.UserNotList.Code, consts.UserNotList.Msg, gin.H{})
		return
	}
	dbConn.Commit()
	response.Success(c, datas)
}

func CreateUser(c *gin.Context) {
	dbConn := model.UseDbConn().Begin()
	defer dbConn.Rollback()
	reqData := api_param.CreateUserReq{
		Name:     "Mr.Wang",
		Password: "123456",
		Phone:    "12345678901",
		Email:    "123",
	}

	err := model.CreateUserModelFactory(dbConn).Create(reqData)
	if err != nil {
		response.Fail(c, consts.UserNotList.Code, consts.UserNotList.Msg, gin.H{})
		return
	}
	dbConn.Commit()
	response.Success(c, "")
}

func DeleteUser(c *gin.Context) {
	dbConn := model.UseDbConn().Begin()
	defer dbConn.Rollback()
	err := model.CreateUserModelFactory(dbConn).Delete(2)
	if err != nil {
		response.Fail(c, consts.UserDeleteFail.Code, consts.UserDeleteFail.Msg, gin.H{})
		return
	}
	dbConn.Commit()
	response.Success(c, "")
}

func UpdateUser(c *gin.Context) {
	dbConn := model.UseDbConn().Begin()
	defer dbConn.Rollback()
	reqData := api_param.UpdateUserReq{
		Id:       1,
		Phone:    "123-4567-8901",
		Email:    "123@163.com",
		Password: utils.MakePassword("123456", consts.Salt),
	}

	_, err := govalidator.ValidateStruct(reqData)
	if err != nil {
		fmt.Println(err)
		response.Fail(c, consts.UserDeleteFail.Code, consts.UserDeleteFail.Msg, gin.H{})
		return
	}
	err = model.CreateUserModelFactory(dbConn).Update(reqData)
	if err != nil {
		response.Fail(c, consts.UserDeleteFail.Code, consts.UserDeleteFail.Msg, gin.H{})
		return
	}
	dbConn.Commit()
	response.Success(c, "")
}
