package handler

import (
	"context"
	"fmt"

	"bitbucket.org/cloud-platform/vnpt-sso-authentication/initial"
	"bitbucket.org/cloud-platform/vnpt-sso-authentication/oauth2"
	modelOAuth "github.com/go-oauth2/oauth2/v4/models"
	"github.com/google/uuid"
)

func SaveCredentialsToDB(ctx context.Context, userID string) (interface{}, error) {
	clientID := uuid.New().String()
	clientSecret := uuid.New().String()

	ret := &modelOAuth.Client{
		ID:     clientID,
		Secret: clientSecret,
		UserID: userID,
	}

	err := oauth2.ClientStore.Set(clientID, ret)
	if err != nil {
		fmt.Println("Set Credential Failed: " + err.Error())
		return nil, err
	}

	collection := "VSSO_Credentials"
	err = mgodb.SaveMongo(initial.MgoDBName, collection, clientID, ret)
	if err != nil {
		fmt.Println("Save Mongodb Credential Failed: " + err.Error())
		return nil, err
	}

	return ret, nil
}
