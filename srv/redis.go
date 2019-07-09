package srv

import (
	"api/config"
	"api/log"

	"github.com/go-redis/redis"
)

// RedisClientInterface redis 实例拥有的功能
type RedisClientInterface interface {
	redis.Cmdable
	Subscribe(...string) *redis.PubSub
	Close() error
}

type RedisClient struct {
	isCluster bool
	Client    RedisClientInterface
}

var redisClient = &RedisClient{}

// Redis redis客户端实例，集群和非集群都使用该实例
var Redis RedisClientInterface

// Init 初始化redis
func (r *RedisClient) Init(conf *config.Config) {
	redisConf := conf.Redis
	r.isCluster = redisConf.Cluster

	if len(redisConf.Addrs) == 0 {
		log.Infof("redis config is empty, skip init redis")
		return
	}

	if redisConf.Cluster {
		r.Client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisConf.Addrs,
			Password: redisConf.Password,
		})
	} else {
		r.Client = redis.NewClient(&redis.Options{
			Addr:     redisConf.Addrs[0],
			Password: redisConf.Password,
			DB:       redisConf.DB,
		})
	}
	_, err := r.Client.Ping().Result()
	if err != nil {
		log.Errorf("redis 连接失败, err=%v", err)
		panic("connect to redis error")
	}
	Redis = r.Client
}

// Close 关闭redis连接
func (r *RedisClient) Close() {
	if r.Client != nil {
		r.Client.Close()
	}
}
