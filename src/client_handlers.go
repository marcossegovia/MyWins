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

// BootstrapClient is used to create a default client into the DB.
func BootstrapClient() {
	var err error
	persistence := mongo.NewMongoAPIClient()
	myClient := &osin.DefaultClient{
		Id:          "1234",
		Secret:      "abcd",
		RedirectUri: "http://localhost:8081/accesstoken",
	}
	err = persistence.SetClient("1234", myClient)

	if err != nil {
		fmt.Println(err.Error())
	}

}

//Login is the web entry point to get the Access Token to start using MyWins API.
func Login(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<html><body>"))

	w.Write([]byte(fmt.Sprintf("<form action=\"/login\" method=\"POST\">")))

	w.Write([]byte("Login: <input type=\"text\" name=\"client_id\" /><br/>"))
	w.Write([]byte("Password: <input type=\"password\" name=\"client_secret\" /><br/>"))
	w.Write([]byte("<input type=\"submit\"/>"))

	w.Write([]byte("</form>"))

	w.Write([]byte("</body></html>"))
}

//LoginPost is the common entry point to get the Access Token to start using MyWins API by Posting the client_id and the client_secret.
func LoginPost(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	clientID := r.Form.Get("client_id")
	clientSecret := r.Form.Get("client_secret")
	clientConfig := &osincli.ClientConfig{
		ClientId:     clientID,
		ClientSecret: clientSecret,
		AuthorizeUrl: "http://localhost:8080/authorize",
		TokenUrl:     "http://localhost:8080/token",
		RedirectUrl:  "http://localhost:8081/accesstoken",
	}

	var err error
	oauthClient, err = osincli.NewClient(clientConfig)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("ERROR: %s\n", err)))
		return
	}
	authorizeRequest = oauthClient.NewAuthorizeRequest(osincli.CODE)

	authorizedURLToRedirect := authorizeRequest.GetAuthorizeUrl().String()
	http.Redirect(w, r, authorizedURLToRedirect, 301)
}

// AuthForAccessToken exchanges the Auth Token from the AuthorizeRequest to an AccessToken.
func AuthForAccessToken(w http.ResponseWriter, r *http.Request) {
	if authorizeRequest == nil {
		w.Write([]byte("You have to login first."))
		return
	}
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
