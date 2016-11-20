package main

import (
	"fmt"
	"net/http"

	"github.com/MarcosSegovia/MyWins/src/wins/infrastructure/mongo"
	"github.com/RangelReale/osin"
	"github.com/RangelReale/osincli"
)

var (
	oauthClient      *osincli.Client
	authorizeRequest *osincli.AuthorizeRequest
)

func BootstrapClient() {
	var err error
	persistence := mongo.NewMongoApiClient()
	myClient := &osin.DefaultClient{
		Id:          "1234",
		Secret:      "abcd",
		RedirectUri: "http://localhost:8081/accesstoken",
	}
	err = persistence.SetClient("1234", myClient)

	if err != nil {
		fmt.Println(err.Error())
	}

	clientConfig := &osincli.ClientConfig{
		ClientId:     "1234",
		ClientSecret: "abcd",
		AuthorizeUrl: "http://localhost:8080/authorize",
		TokenUrl:     "http://localhost:8080/token",
		RedirectUrl:  "http://localhost:8081/accesstoken",
	}
	oauthClient, err = osincli.NewClient(clientConfig)
	if err != nil {
		panic(err)
	}
	authorizeRequest = oauthClient.NewAuthorizeRequest(osincli.CODE)
}

/**
 * Entry point to get the Access Token
 */
func Login(w http.ResponseWriter, r *http.Request) {

	url := authorizeRequest.GetAuthorizeUrl()

	w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Login</a>", url.String())))
}

/**
 * Exchanges the Auth Token from the AuthorizeRequest to an AccessToken
 */
func AuthForAccessToken(w http.ResponseWriter, r *http.Request) {
	authorizeRequestData, err := authorizeRequest.HandleRequest(r)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("ERROR: %s\n", err)))
		return
	}
	accessTokenRequest := oauthClient.NewAccessRequest(osincli.AUTHORIZATION_CODE, authorizeRequestData)

	accessData, err := accessTokenRequest.GetToken()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("ERROR: %s\n", err)))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(buildResponse(accessData.ResponseData))
}
