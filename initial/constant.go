package initial

import (
	"github.com/spf13/viper"

	mongodb "github.com/mixi-gaminh/core-framework/repository/mongodb"
	redisdb "github.com/mixi-gaminh/core-framework/repository/redisdb"
)

var (
	redis redisdb.Cache
	mgodb mongodb.Mgo
	MgoDBName string
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	redis.RedisConstructor(
		viper.GetString(`redis.url`),
		viper.GetInt(`redis.max_clients`),
		viper.GetInt(`redis.min_idle`),
	)

	// initialize mongo constructor
	mgodb.MongoDBConstructor(viper.GetStringSlice(`mongodb.url`),
		viper.GetString(`mongodb.username`),
		viper.GetString(`mongodb.password`),
	)
	MgoDBName = viper.GetString("mongodb.dbname")
}