package dao

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
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
		panic("dao未初始化")
	}
}

func Init(conf *config.Config) *Dao {
	db := initDB(conf.Db)
	redis := initRedis(conf.Redis)

	dao := &Dao{
		Db: db,
		Redis: redis,
	}
	return dao
}

func initDB(dbConfig *config.DbConfig) (db *gorm.DB){
	db, err := gorm.Open(dbConfig.DbType, getDsn(dbConfig))
	if err != nil {
		fmt.Println("数据库连接失败")
		return nil
	}
	db.DB().SetMaxIdleConns(20)
	db.DB().SetConnMaxLifetime(600)
	db.DB().SetMaxOpenConns(1000)
	return
}

func initRedis(c * config.RedisConfig) (redisClient *redis.Client) {
	fmt.Println("初始化redis")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.Db,
	})
	return
}

// 根据配置获取链接字符串
func getDsn(c *config.DbConfig) string {
	if c.DbType == "mysql" {
		return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.DbName, c.Charset)
	} else if c.DbType == "postgres" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", c.Host, c.Port, c.Username, c.DbName, c.Password)
	}
	return ""
}
