package response

import (
	"ginchat/global/consts"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Trans ut.Translator

func Success(c *gin.Context, data interface{}) {
	returnJson(c, http.StatusOK, consts.CurdOK.Code, consts.CurdOK.Msg, "", data)
}

func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	returnJson(c, http.StatusOK, dataCode, msg, "", data)
	c.Abort()
}

func ValidatorError(c *gin.Context, err error) {
	if errs, ok := err.(validator.ValidationErrors); ok {
		wrongParam := removeTopStruct(errs.Translate(Trans))
		returnJson(c, http.StatusBadRequest, consts.ParamsCheckFail.Code, consts.ParamsCheckFail.Msg, "", wrongParam)
	} else {
		errStr := err.Error()
		if strings.ReplaceAll(strings.ToLower(errStr), " ", "") == "multipart:nextpart:eof" {
			returnJson(c, http.StatusBadRequest, consts.ParamsCheckFail.Code, consts.ParamsCheckFail.Msg, "", gin.H{"tips": consts.ErrorNotAllParamsIsBlank})
		} else {
			returnJson(c, http.StatusBadRequest, consts.ParamsCheckFail.Code, consts.ParamsCheckFail.Msg, "", gin.H{"tips": errStr})
		}
	}
	c.Abort()
}

func ErrorSystem(c *gin.Context, msg string, data interface{}) {
	returnJson(c, http.StatusInternalServerError, consts.ServerError.Code, consts.ServerError.Msg, msg, data)
	c.Abort()
}

func returnJson(c *gin.Context, httpCode, dataCode int, msg, dep string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"access_status": gin.H{
			"code":        dataCode,
			"message":     msg,
			"description": dep,
		},
		"result": data,
	})
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.LastIndex(field, ".")+1:]] = err
	}
	return res
}
