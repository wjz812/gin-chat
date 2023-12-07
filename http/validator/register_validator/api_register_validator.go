package register_validator

import (
	"ginchat/core/container"
	"ginchat/global/consts"
	"ginchat/http/validator/api/user"
)

// 注册api
func ApiRegisterValidator() {
	containers := container.CreateFactoryContainers()

	containers.Set(consts.ValidatorPrefix+"UserLogin", user.UserLogin{})

	// containers.Set(consts.ValidatorPrefix+"SendMsg", user.SendMsg{})
}
