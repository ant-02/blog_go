package util

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var (
	once     sync.Once    // 确保初始化只执行一次
	instance *RedisClient // 唯一的 Redis 客户端实例
)

// RedisClient 定义 Redis 客户端结构体
type RedisClient struct {
	client *redis.Client
}

func GetRedis(addr, password string, db int) *RedisClient {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     addr,     // Redis 地址
			Password: password, // Redis 密码
			DB:       db,       // 使用的数据库编号
			PoolSize: 10,       // 连接池大小
		})

		// 测试连接是否成功
		_, err := client.Ping().Result()
		if err != nil {
			panic(fmt.Sprintf("无法连接到 Redis: %v", err))
		}

		instance = &RedisClient{client: client}
	})
	return instance
}

func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(key, value, expiration).Err()
}

func (r *RedisClient) Get(key string) (string, error) {
    val, err := r.client.Get(key).Result()
    if err == redis.Nil {
        return "", fmt.Errorf("键 %s 不存在", key)
    } else if err != nil {
        return "", fmt.Errorf("获取键值对失败: %v", err)
    }
    return val, nil
}
