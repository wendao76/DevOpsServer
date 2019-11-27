package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	confPath = "."
	confFile = "config"
	conf     *Config
)

func init() {
	err := initDefault()
	if err != nil {
	    log.Fatal(err.Error())
	}
}


func GetInstance()  *Config {
	return conf
}

func initDefault() (err error) {
	conf, err = Default()
	if err != nil {
		panic(err)
	}
	return
}

func Init(path string, name string)  (err error) {
	confPath = path
	confFile = name
	conf, err = Default()
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
	err = v.Sub("database").Unmarshal(&dbConfig)
	if err != nil {
	    return nil, err
	}
	err = v.Sub("redis").Unmarshal(&redisConfig)
	if err != nil {
	    return nil, err
	}
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
		return nil, err
	}
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
