package mongo

import (
	"github.com/RangelReale/osin"
	"time"
)

type AuthorizeData struct {
	ClientID    string
	Code        string
	ExpiresIn   int32
	Scope       string
	RedirectURI string
	State       string
	CreatedAt   time.Time
	UserData    interface{}
}

func AuthorizeDataFromOSIN(data *osin.AuthorizeData) *AuthorizeData {
	return &AuthorizeData{
		ClientID:    data.Client.GetId(),
		Code:        data.Code,
		ExpiresIn:   data.ExpiresIn,
		Scope:       data.Scope,
		RedirectURI: data.RedirectUri,
		State:       data.State,
		CreatedAt:   data.CreatedAt,
		UserData:    data.UserData,
	}
}

func (data *AuthorizeData) AuthorizeDataToOSIN(client MongoAPIClient) (*osin.AuthorizeData, error) {
	osinClient, err := client.GetClient(data.ClientID)
	if err != nil {
		return nil, err
	}
	return &osin.AuthorizeData{
		Client:      osinClient,
		Code:        data.Code,
		ExpiresIn:   data.ExpiresIn,
		Scope:       data.Scope,
		RedirectUri: data.RedirectURI,
		State:       data.State,
		CreatedAt:   data.CreatedAt,
		UserData:    data.UserData,
	}, nil
}

type AccessData struct {
	ClientID     string
	AccessToken  string
	RefreshToken string
	ExpiresIn    int32
	Scope        string
	RedirectURI  string
	CreatedAt    time.Time
	UserData     interface{}
}

func AccessDataFromOSIN(data *osin.AccessData) *AccessData {
	return &AccessData{
		ClientID:     data.Client.GetId(),
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
		ExpiresIn:    data.ExpiresIn,
		Scope:        data.Scope,
		RedirectURI:  data.RedirectUri,
		CreatedAt:    data.CreatedAt,
		UserData:     data.UserData,
	}
}

func (data *AccessData) AccessDataToOSIN(client MongoAPIClient) (*osin.AccessData, error) {
	osinClient, err := client.GetClient(data.ClientID)
	if err != nil {
		return nil, err
	}
	return &osin.AccessData{
		Client:       osinClient,
		AccessToken:  data.AccessToken,
		RefreshToken: data.RefreshToken,
		ExpiresIn:    data.ExpiresIn,
		Scope:        data.Scope,
		RedirectUri:  data.RedirectURI,
		CreatedAt:    data.CreatedAt,
		UserData:     data.UserData,
	}, nil
}
