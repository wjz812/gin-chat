package model

import (
	"ginchat/global/config"
	"ginchat/global/consts"

	"gorm.io/gorm"
)

type BaseModel struct {
	*gorm.DB  `gorm:"-" json:"-"`
	Id        int64  `gorm:"primaryKey" json:"id"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}

func UseDbConn() *gorm.DB {
	if config.GormDB == nil {
		panic(consts.MysqlConfigErrorMsg)
	}
	return config.GormDB
}
