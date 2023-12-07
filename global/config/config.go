package config

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GormDB *gorm.DB
	ZapLog *zap.Logger
)
