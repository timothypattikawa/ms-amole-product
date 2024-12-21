package rd

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	addr     string
	password string
	db       int
}

func NewRedisConfig(v *viper.Viper) *RedisConfig {
	return &RedisConfig{
		addr:     v.GetString("database.redis.addr"),
		password: "",
		db:       0,
	}
}

func (rc *RedisConfig) GetClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     rc.addr,
		Password: rc.password,
		DB:       rc.db,
	})

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// Test connection
	pong, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Printf("Connected to Redis: %v", pong)

	return redisClient
}
