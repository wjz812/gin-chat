package config

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GormDB *gorm.DB
	ZapLog *zap.Logger
	Conf   Config
)

type Config struct {
	Mysql struct {
		Write struct {
			Host               string        `yaml:"Host"`
			DataBase           string        `yaml:"DataBase"`
			Port               int           `yaml:"Port"`
			Prefix             string        `yaml:"Prefix"`
			User               string        `yaml:"User"`
			Pass               string        `yaml:"Pass"`
			Charset            string        `yaml:"Charset"`
			SetMaxIdleConns    int           `yaml:"SetMaxIdleConns"`
			SetMaxOpenConns    int           `yaml:"SetMaxOpenConns"`
			SetConnMaxLifetime time.Duration `yaml:"SetConnMaxLifetime"`
		} `yaml:"Write"`
		Read struct {
			Host               string        `yaml:"Host"`
			DataBase           string        `yaml:"DataBase"`
			Port               int           `yaml:"Port"`
			Prefix             string        `yaml:"Prefix"`
			User               string        `yaml:"User"`
			Pass               string        `yaml:"Pass"`
			Charset            string        `yaml:"Charset"`
			SetMaxIdleConns    int           `yaml:"SetMaxIdleConns"`
			SetMaxOpenConns    int           `yaml:"SetMaxOpenConns"`
			SetConnMaxLifetime time.Duration `yaml:"SetConnMaxLifetime"`
		} `yaml:"Read"`
	} `yaml:"Mysql"`
	Redis struct {
		Host               string        `yaml:"Host"`
		Port               int           `yaml:"Port"`
		Auth               string        `yaml:"Auth"`
		MaxIdle            int           `yaml:"MaxIdle"`
		MaxActive          int           `yaml:"MaxActive"`
		IdleTimeout        time.Duration `yaml:"IdleTimeout"`
		IndexDb            int           `yaml:"IndexDb"`
		ConnFailRetryTimes int           `yaml:"ConnFailRetryTimes"`
		ReConnectInterval  time.Duration `yaml:"ReConnectInterval"`
	} `yaml:"Redis"`

	HTTPServer struct {
		API struct {
			Port string `yaml:"Port"`
		} `yaml:"Api"`
		Web struct {
			Port string `yaml:"Port"`
		} `yaml:"Web"`
		AllowCrossDomain bool `yaml:"AllowCrossDomain"`
	} `yaml:"HttpServer"`

	Token struct {
		JwtTokenSignKey         string        `yaml:"JwtTokenSignKey"`
		JwtTokenOnlineUsers     int           `yaml:"JwtTokenOnlineUsers"`
		JwtTokenCreatedExpireAt time.Duration `yaml:"JwtTokenCreatedExpireAt"`
		JwtTokenRefreshAllowSec int64         `yaml:"JwtTokenRefreshAllowSec"`
		JwtTokenRefreshExpireAt int64         `yaml:"JwtTokenRefreshExpireAt"`
		JwtIssuer               string        `yaml:"JwtIssuer"`
		BindContextKeyName      string        `yaml:"BindContextKeyName"`
		IsCacheToRedis          int           `yaml:"IsCacheToRedis"`
		EnqueteTokenExpireAt    int64         `yaml:"EnqueteTokenExpireAt"`
	} `yaml:"Token"`
}
