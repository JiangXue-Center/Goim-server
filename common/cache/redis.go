package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"time"
)

var RedisClient *redis.Client

func InitRedisConn() {
	// 读取配置文件中的Redis配置项
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../user/config") // 指定配置文件的目录
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	// 配置 Redis 客户端
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.db"),
		PoolSize:     viper.GetInt("redis.pool_size"),
		MinIdleConns: viper.GetInt("redis.min_idle_conns"),
		PoolTimeout:  viper.GetDuration("redis.pool_timeout") * time.Second,
	})
}
