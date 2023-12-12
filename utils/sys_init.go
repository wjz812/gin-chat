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
	"gopkg.in/yaml.v3"
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

func InitConfigYmal() {
	yamlBytes, err := os.ReadFile("global/config/config.yml")
	if err != nil {
		log.Fatalf("Failed to read Yaml file: %v", err)
	}

	err = yaml.Unmarshal(yamlBytes, &config.Conf)
	if err != nil {
		log.Fatalf("Failed to decode Yaml: %v", err)
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

	Host := config.Conf.Mysql.Write.Host
	DataBase := config.Conf.Mysql.Write.DataBase
	Port := config.Conf.Mysql.Write.Port
	User := config.Conf.Mysql.Write.User
	Pass := config.Conf.Mysql.Write.Pass
	Charset := config.Conf.Mysql.Write.Charset

	// fmt.Println("=====:", Port, Host, DataBase, User, Pass, Charset)

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=false&loc=Local",
		User, Pass, Host, Port, DataBase, Charset)

	//"root:root@tcp(localhost:3306)/ginchat?charset=utf8&parseTime=True&loc=Local"
	fmt.Println("=====config mysql:", dns)

	DB, _ = gorm.Open(mysql.Open(dns),
		&gorm.Config{Logger: newLogger})

	// DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
	// 	&gorm.Config{Logger: newLogger})
	// fmt.Println("config mysql:", viper.Get("mysql"))
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
