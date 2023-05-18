package db

import (
	"github.com/go-redis/redis"
	"github.com/svopper/kalles_weather_dashboard_v2/pkg/common/util"
)

func Init() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     util.GetEnvVariable("REDIS_URL"),
		Password: util.GetEnvVariable("REDIS_DB_PSWD"),
		DB:       0, // use default DB
	})

	// ensure connection
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}

	return rdb
}
