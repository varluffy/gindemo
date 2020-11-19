package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"github.com/varluffy/gindemo/pkg/logger"
	"time"
)

// InitRedis 初始化redis
func InitRedis() (*redis.ClusterClient, func(), error) {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        viper.GetStringSlice("Redis.Addrs"),
		Username:     viper.GetString("Redis.Username"),
		Password:     viper.GetString("Redis.Password"),
		PoolSize:     viper.GetInt("Redis.PoolSize"),
		MinIdleConns: viper.GetInt("Redis.MinIdleConns"),
		MaxConnAge:   time.Second * time.Duration(viper.GetInt("Redis.MaxConnAge")),
	},
	)
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Errorf(context.Background(), "Redis client ping error: %s", err.Error())
		return nil, nil, err
	}
	cleanFunc := func() {
		err := client.Close()
		if err != nil {
			logger.Errorf(context.Background(), "Redis client close error: %s", err.Error())
		}
	}
	return client, cleanFunc, err
}
