package main

import (
	"ginchat/global/config"
	"ginchat/http/validator/register_validator"
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	register_validator.ApiRegisterValidator()
	utils.InitConfig()
	utils.InitConfigYmal()
	utils.InitMySql()
	utils.InitRedis()

	r := router.Router()
	r.Run(config.Conf.HTTPServer.API.Port)
}
