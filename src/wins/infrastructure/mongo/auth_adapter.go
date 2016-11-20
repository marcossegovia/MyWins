package mongo

import (
	"time"
	"github.com/RangelReale/osin"
)

type AuthorizeData struct {
	ClientId    string
	Code        string
	ExpiresIn   int32
	Scope       string
	RedirectUri string
	State       string
	CreatedAt   time.Time
	UserData    interface{}
}

func AuthorizeDataFromOSIN(data *osin.AuthorizeData) *AuthorizeData {
	return &AuthorizeData{
		ClientId:    data.Client.GetId(),
		Code:        data.Code,
		ExpiresIn:   data.ExpiresIn,
		Scope:       data.Scope,
		RedirectUri: data.RedirectUri,
		State:       data.State,
		CreatedAt:   data.CreatedAt,
		UserData:    data.UserData,
	}
}

func (data *AuthorizeData) AuthorizeDataToOSIN(client MongoApiClient) (*osin.AuthorizeData, error) {
	osinClient, err := client.GetClient(data.ClientId)
	if err != nil {
		return nil, err
	}
	return &osin.AuthorizeData{
		Client:      osinClient,
		Code:        data.Code,
		ExpiresIn:   data.ExpiresIn,
		Scope:       data.Scope,
		RedirectUri: data.RedirectUri,
		State:       data.State,
		CreatedAt:   data.CreatedAt,
		UserData:    data.UserData,
	}, nil
}

type AccessData struct {
	ClientId     string
	AccessToken  string
	RefreshToken string
	ExpiresIn    int32
	Scope        string
	RedirectUri  string
	CreatedAt    time.Time
	UserData     interface{}
}

func AccessDataFromOSIN(data *osin.AccessData) *AccessData {
	return &AccessData{
		ClientId:     data.Client.GetId(),
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
		ExpiresIn:    data.ExpiresIn,
		Scope:        data.Scope,
		RedirectUri:  data.RedirectUri,
		CreatedAt:    data.CreatedAt,
		UserData:     data.UserData,
	}
}

func (data *AccessData) AccessDataToOSIN(client MongoApiClient) (*osin.AccessData, error) {
	osinClient, err := client.GetClient(data.ClientId)
	if err != nil {
		return nil, err
	}
	return &osin.AccessData{
		Client:       osinClient,
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
		ExpiresIn:    data.ExpiresIn,
		Scope:        data.Scope,
		RedirectUri:  data.RedirectUri,
		CreatedAt:    data.CreatedAt,
		UserData:     data.UserData,
	}, nil
}
