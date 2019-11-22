package dao

import (
	"fmt"
	"errors"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_web/internal/common/config"
)

var dao *Dao

type Dao struct {
	Db *gorm.DB
	Redis *redis.Client
}

func Get() (*Dao, error) {
	if dao!= nil {
		return dao, nil
	} else {
		return nil, errors.New("dao初始化失败")
	}
}

func Init(conf *config.Config)(err error) {
	db, err := initDB(conf.Db)

	if err != nil {
		return err
	}
	redis, err := initRedis(conf.Redis)
	if err != nil {
		return err
	}
	dao = &Dao{
		Db: db,
		Redis: redis,
	}
	fmt.Println("dao初始化成功")
	return nil
}

func initDB(dbConfig *config.DbConfig) (db *gorm.DB, err error){
	db, err = gorm.Open(dbConfig.Type, getDsn(dbConfig))
	if err != nil {
		fmt.Println("数据库连接失败")
		return nil, err
	}
	db.DB().SetMaxIdleConns(20)
	db.DB().SetConnMaxLifetime(600)
	db.DB().SetMaxOpenConns(1000)
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
