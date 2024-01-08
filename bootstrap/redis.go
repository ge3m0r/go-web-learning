package bootstrap

import(
	"fmt"
	"golearning/pkg/config"
	"golearning/pkg/redis"
)

func SetupRedis(){
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"),config.GetString("redis.port")),
		config.GetString("redis.usernanme"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}