package redis

import (
	"context"
	"golearning/pkg/logger"
	"sync"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type RedisClient struct{
	Client *redis.Client
	Context context.Context
}

var once sync.Once

var Redis *RedisClient

func ConnectRedis(address string, username string,password string, db int){
	once.Do(func () {
		Redis = NewClient(address,username,password,db)
		
	})
}

func NewClient(address string, username string, password string, db int) *RedisClient{
	rds := &RedisClient{}

	rds.Context = context.Background()

	rds.Client = redis.NewClient(&redis.Options{
		Addr:        address,
		Username:    username,
		Password:    password,
		DB:          db,
	})

	err := rds.Ping()
	logger.LogIf(err)

	return rds
}

func (rds RedisClient) Ping() error{
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool{
	if err := rds.Client.Set(rds.Context, key , value,expiration).Err(); err != nil{
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

func (rds RedisClient) Get(key string) string{
	 result, err := rds.Client.Get(rds.Context, key).Result()
	 if err != nil{
		if err != redis.Nil{
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

func (rds RedisClient) Has(key string) bool{
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil{
	   if err != redis.Nil{
		   logger.ErrorString("Redis", "Has", err.Error())
	   }
	   return false
   }
   return true
}

func (rds RedisClient) Del(keys ...string) bool{
    if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil{
	   logger.ErrorString("Redis", "Del", err.Error())
	   return false
    }
   return true
}

func (rds RedisClient) FlushDB() bool{
    if err := rds.Client.FlushDB(rds.Context).Err(); err != nil{
	   logger.ErrorString("Redis", "FlushDB", err.Error())
	   return false
    }
   return true
}

func (rds RedisClient) Increment(parameters ...interface{}) bool{
	switch len(parameters){
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Incr(rds.Context,key).Err(); err != nil{
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.IncrBy(rds.Context,key,value).Err(); err != nil{
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Increment", "参数过多")
		return false
	}
	return true
}

func (rds RedisClient) Decrement(parameters ...interface{}) bool{
	switch len(parameters){
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Decr(rds.Context,key).Err(); err != nil{
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.DecrBy(rds.Context,key,value).Err(); err != nil{
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Decrement", "参数过多")
		return false
	}
	return true
}