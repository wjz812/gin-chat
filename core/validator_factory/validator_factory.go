package validator_factory

import (
	"ginchat/core/container"

	"github.com/gin-gonic/gin"
)

func Create(key string) func(c *gin.Context) {
	if value := container.CreateFactoryContainers().Get(key); value != nil {
		if val, isOk := value.(validatorInterface); isOk {
			return val.CheckParams
		}
	}
	return nil
}

type validatorInterface interface {
	CheckParams(context *gin.Context)
}
