package srv

import (
	"api/config"
	"api/srv/db"
)

// SrvClient 服务接口
type SrvClient interface {
	Init(*config.Config) // 用于初始化具体服务
	Close()              // 关闭服务
}

var clients []SrvClient

// Init 初始化服务
func Init(conf *config.Config) {
	// init redis
	clients = []SrvClient{
		db.Gorm,
		db.Xorm,
		redisClient,
	}
	for _, cl := range clients {
		if cl != nil {
			cl.Init(conf)
		}
	}
}

// Close 关闭所有服务连接
func Close() {
	if len(clients) > 0 {
		for _, cl := range clients {
			if cl != nil {
				cl.Close()
			}
		}
	}
}
