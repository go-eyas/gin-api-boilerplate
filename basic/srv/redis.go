package srv

import (
	"basic/log"

	"github.com/go-eyas/toolkit/redis"
)

type RedisClient struct {
	Client *redis.RedisClient
}

var RedisSrv = &RedisClient{}
var Redis *redis.RedisClient

// Init 初始化redis
func (c *RedisClient) Init(redisConf *RedisConfig) {

	if len(redisConf.Addrs) == 0 {
		// log.Infof("redis config is empty, skip init redis")
		return
	}
	cli, err := redis.New(&redis.Config{
		Cluster:  redisConf.Cluster,
		Addrs:    redisConf.Addrs,
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})

	if err != nil {
		log.Errorf("redis init error: %s", err.Error())
	}

	c.Client = cli
}

