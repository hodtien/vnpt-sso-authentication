package oauth2

import (
	"encoding/json"
	"fmt"
	"log"

	"bitbucket.org/cloud-platform/vnpt-sso-authentication/initial"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	modelOAuth "github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"

	mongodb "github.com/mixi-gaminh/core-framework/repository/mongodb"
)

var mgodb mongodb.Mgo
var Manager *manage.Manager
var Srv *server.Server
var ClientStore *store.ClientStore

func init() {
	Manager = manage.NewDefaultManager()
	Manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	Manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr: viper.GetString("redis.url"),
		DB:   15,
	}))

	ClientStore = store.NewClientStore()
	Manager.MapClientStorage(ClientStore)

	Srv = server.NewDefaultServer(Manager)
	Srv.SetAllowGetAccessRequest(true)
	Srv.SetClientInfoHandler(server.ClientFormHandler)
	Manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	Srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	Srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	LoadCredetialsFromDB()
}

func LoadCredetialsFromDB() error {
	collection := "VSSO_Credentials"
	listRecord, err := mgodb.GetAllRecordInCollection(initial.MgoDBName, collection)
	if err != nil {
		log.Printf("Get All Credentials from MongoDB Failed: %s", err)
		return err
	}

	for _, record := range listRecord {
		clientID := fmt.Sprintf("%v", record["_id"])
		bodyBytes, _ := json.Marshal(record)
		credential := new(modelOAuth.Client)
		err := json.Unmarshal(bodyBytes, &credential)
		if err != nil {
			log.Printf("Load Credentials %v Failed: %s", clientID, err)
			continue
		}
		err = ClientStore.Set(clientID, credential)
		if err != nil {
			log.Printf("Set Credentials %v Failed: %s", clientID, err)
			continue
		}
	}
	return nil
}
