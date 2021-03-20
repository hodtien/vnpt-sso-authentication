package usershandler

import (
	mongodb "github.com/mixi-gaminh/core-framework/repository/mongodb"
	redisdb "github.com/mixi-gaminh/core-framework/repository/redisdb"
)

var (
	redis redisdb.Cache
	mgodb mongodb.Mgo
)
