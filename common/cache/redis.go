package cache

import (
	"context"
	"fmt"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	} else {
		fmt.Println("Connected to Redis successfully")
	}
}

// SetValue 在 Redis 中设置一个键值对
func SetValue(key string, value string, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := RedisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetValue 从 Redis 中获取一个值
func GetValue(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	value, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}
