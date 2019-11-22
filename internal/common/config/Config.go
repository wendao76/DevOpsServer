package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	confPath = "."
	confFile = "config"
	Conf     *Config
)

func InitDefault() (err error) {
	Conf, err = Default()
	if err != nil {
		panic(err)
	}
	return
}

func Init(path string, name string)  (err error) {
	confPath = path
	confFile = name
	Conf, err = Default()
	if err != nil {
		panic(err)
	}
	return
}

func Default() (conf *Config, err error) {
	v := viper.New()
	v.AddConfigPath(confPath)
	v.SetConfigName(confFile)
	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var (
		dbConfig    *DbConfig
		redisConfig *RedisConfig
	)
	v.Sub("database").Unmarshal(&dbConfig)
	v.Sub("redis").Unmarshal(&redisConfig)
	fmt.Println("fmt.Println(redisConfig)")
	fmt.Println(redisConfig)
	conf = &Config{
		Db:    dbConfig,
		Redis: redisConfig,
	}
	return
}

func New(configPath, configFile string) (conf *Config, err error) {
	confPath, confFile = configPath, configFile
	conf, err = Default()
	if err != nil {
		panic(err)
	}
	Conf = conf
	return
}

type Config struct {
	Db    *DbConfig
	Redis *RedisConfig
}

//数据库配置
type DbConfig struct {
	Host     string
	DbName   string
	Type   string
	Username string
	Password string
	Port     int
	Charset  string
}

//redis配置
type RedisConfig struct {
	Addr string
	Db   int
	Password string
}
