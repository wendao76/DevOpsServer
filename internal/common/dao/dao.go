package dao

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_web/internal/common/config"
	"log"
)

var Dao *dao
func init() {
	err := initDefault()
	if err != nil {
	    log.Fatal(err.Error())
	}
}

type dao struct {
	Db *gorm.DB
	Redis *redis.Client
}

func initDefault() error{
	err := Init(config.Config.Db, config.Config.Redis)
	if err != nil {
		return err
	}
	return nil
}

//初始化操作
func Init(dbConf *config.DbConfig, redisConf *config.RedisConfig)(err error) {
	db, err := initDB(dbConf)

	if err != nil {
		return err
	}
	redisClient, err := initRedis(redisConf)
	if err != nil {
		return err
	}
	Dao = &dao{
		Db: db,
		Redis: redisClient,
	}
	fmt.Println("dao初始化成功")
	return nil
}

//初始化数据库相关内容
func initDB(dbConfig *config.DbConfig) (db *gorm.DB, err error){
	db, err = gorm.Open(dbConfig.Type, getDsn(dbConfig))
	if err != nil {
		fmt.Println("数据库连接失败")
		return nil, err
	}
	db.DB().SetMaxIdleConns(20)
	db.DB().SetConnMaxLifetime(600)
	db.DB().SetMaxOpenConns(1000)

	AutoMigrateTables(db)
	return
}

func initRedis(c * config.RedisConfig) (redisClient *redis.Client, err error) {
	fmt.Println("初始化redis")
	fmt.Println(c)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.Db,
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return
}

// 根据配置获取链接字符串
func getDsn(c *config.DbConfig) string {
	if c.Type == "mysql" {
		return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.DbName, c.Charset)
	} else if c.Type == "postgres" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", c.Host, c.Port, c.Username, c.DbName, c.Password)
	}
	return ""
}

//自动创建表
func AutoMigrateTables(db *gorm.DB) {
	fmt.Println("数据表融合")
	//数据结构初始化
	//db.AutoMigrate(po.RunLog{}, po.RunErrorLog{})
}
