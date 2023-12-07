package utils

import (
	"context"
	"fmt"
	"ginchat/global/config"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
	ctx = context.Background()
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("global/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("inin err :", err)
	}
}

func InitMySql() {
	//自定义日志模板，打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)

	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{Logger: newLogger})
	fmt.Println("config mysql:", viper.Get("mysql"))
	config.GormDB = DB
}

func InitRedis() {
	fmt.Println("==== InitRedis ====")
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConns"),
		Password:     viper.GetString("redis.password"),
	})

	// err := Red.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// 	fmt.Println("====key", err)
	// }

	// val, err := Red.Get(ctx, "key").Result()
	// if err != nil {
	// 	fmt.Println("====key", err)
	// }
	// fmt.Println("====key", val)

	// fmt.Println("==== InitRedis ==== End")
}

const (
	PublishKey = "websocket"
)

func Publish(ctx context.Context, channel string, msg string) error {
	var err error

	err = Red.Publish(ctx, channel, msg).Err()

	return err
}

func Subscribe(ctx context.Context, channel string) (string, error) {

	sub := Red.Subscribe(ctx, channel)

	msg, err := sub.ReceiveMessage(ctx)

	fmt.Println("Subscribe .....", msg.Payload)
	return msg.Payload, err
}
