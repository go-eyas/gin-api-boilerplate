package srv

import (
	"api/config"
	"toolkit/log"
	"toolkit/redis"
)

type RedisClient struct{}

var redisClient = &RedisClient{}

// Init 初始化redis
func (RedisClient) Init(conf *config.Config) {
	redisConf := conf.Redis

	if len(redisConf.Addrs) == 0 {
		log.Infof("redis config is empty, skip init redis")
		return
	}
	err := redis.Init(&redis.Config{
		Cluster:  conf.Redis.Cluster,
		Addrs:    conf.Redis.Addrs,
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})

	if err != nil {
		log.Errorf("redis init error: %s", err.Error())
	}
}

// Close 关闭redis连接
func (r *RedisClient) Close() {
	redis.Client.Close()
}
