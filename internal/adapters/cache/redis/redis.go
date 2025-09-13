package cache

import (
	"fmt"

	"github.com/pawannn/cashtrack/internal/pkg/config"
	"github.com/pawannn/cashtrack/internal/ports"
	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	rClient *redis.Client
}

func InitRedisService(config config.CashTrackCfg) ports.CacheRepo {
	addr := fmt.Sprintf("%s:%d", config.CacheHost, config.CachePort)
	redisClient := redis.NewClient(&redis.Options{
		DB:       config.CacheDB,
		Password: config.CachePass,
		Addr:     addr,
	})
	return &RedisService{
		rClient: redisClient,
	}
}
