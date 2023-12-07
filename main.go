package main

import (
	"ginchat/http/validator/register_validator"
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	register_validator.ApiRegisterValidator()
	utils.InitConfig()
	utils.InitMySql()
	utils.InitRedis()

	r := router.Router()
	r.Run(":8081")
}
