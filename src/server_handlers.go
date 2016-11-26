package main

import (
        "encoding/json"
        "fmt"
        "net/http"

        "github.com/MarcosSegovia/MyWins/src/wins/domain"
        "github.com/MarcosSegovia/MyWins/src/wins/infrastructure/mongo"
        "github.com/RangelReale/osin"
)

// Welcome is an info endpoint to provide guidance to user requests.
func Welcome(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Welcome!!!")
}

// Authorization is the endpoint to start the authorization of the current user login attempt.
func Authorization(w http.ResponseWriter, r *http.Request) {
        authorizationServer := osin.NewServer(osin.NewServerConfig(), mongo.NewMongoAPIClient())
        resp := authorizationServer.NewResponse()
        defer resp.Close()

        authorizeRequest := authorizationServer.HandleAuthorizeRequest(resp, r)
        if authorizeRequest != nil {
                //LOGIN
                authorizeRequest.Authorized = true
                authorizationServer.FinishAuthorizeRequest(resp, r, authorizeRequest)
        }
        if resp.IsError && resp.InternalError != nil {
                fmt.Printf("ERROR: %s\n", resp.InternalError)
        }
        osin.OutputJSON(resp, w, r)
}

// AccessToken is the endpoint to generate the Token, after the generation it is redirect to the redirectURL.
func AccessToken(w http.ResponseWriter, r *http.Request) {
        authorizationServer := osin.NewServer(osin.NewServerConfig(), mongo.NewMongoAPIClient())
        resp := authorizationServer.NewResponse()
        defer resp.Close()

        accessRequest := authorizationServer.HandleAccessRequest(resp, r)

        if accessRequest != nil {
                accessRequest.Authorized = true
                authorizationServer.FinishAccessRequest(resp, r, accessRequest)
        }
        if resp.IsError && resp.InternalError != nil {
                fmt.Printf("ERROR: %s\n", resp.InternalError)
        }
        osin.OutputJSON(resp, w, r)
}

// Information retrieves information about the current token provided, such as the number of second before expiry.
func Information(w http.ResponseWriter, r *http.Request) {
        authorizationServer := osin.NewServer(osin.NewServerConfig(), mongo.NewMongoAPIClient())
        resp := authorizationServer.NewResponse()
        defer resp.Close()

        infoRequest := authorizationServer.HandleInfoRequest(resp, r)

        if infoRequest != nil {
                authorizationServer.FinishInfoRequest(resp, r, infoRequest)
        }
        osin.OutputJSON(resp, w, r)
}

// GetAllWins is the endpoint that provides the win list from the user.
func GetAllWins(w http.ResponseWriter, r *http.Request) {
        authorizationServer := osin.NewServer(osin.NewServerConfig(), mongo.NewMongoAPIClient())
        resp := authorizationServer.NewResponse()
        defer resp.Close()

        infoRequest := authorizationServer.HandleInfoRequest(resp, r)

        if infoRequest != nil {
                api := domain.NewAPI(mongo.NewMongoAPIClient())
                wins, err := api.FindAllWins()
                if err != nil {
                        http.Error(w, domain.GeneralError, http.StatusInternalServerError)
                        return
                }
                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(http.StatusOK)
                w.Write(buildResponse(wins))
        } else {
                osin.OutputJSON(resp, w, r)
        }

}

// GetAllFails is the endpoint that provides the fail list from the user.
func GetAllFails(w http.ResponseWriter, r *http.Request) {
        authorizationServer := osin.NewServer(osin.NewServerConfig(), mongo.NewMongoAPIClient())
        resp := authorizationServer.NewResponse()
        defer resp.Close()

        infoRequest := authorizationServer.HandleInfoRequest(resp, r)

        if infoRequest != nil {
                api := domain.NewAPI(mongo.NewMongoAPIClient())
                fails, err := api.FindAllFails()
                if err != nil {
                        http.Error(w, domain.GeneralError, http.StatusInternalServerError)
                        return
                }
                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(http.StatusOK)
                w.Write(buildResponse(fails))
        } else {
                osin.OutputJSON(resp, w, r)
        }
}

// AddWin is the endpoint to submit the win of the day to the user win list.
func AddWin(w http.ResponseWriter, r *http.Request) {
        authorizationServer := osin.NewServer(osin.NewServerConfig(), mongo.NewMongoAPIClient())
        resp := authorizationServer.NewResponse()
        defer resp.Close()

        infoRequest := authorizationServer.HandleInfoRequest(resp, r)

        if infoRequest != nil {
                api := domain.NewAPI(mongo.NewMongoAPIClient())
                err := api.AddWin()
                if err != nil {
                        http.Error(w, domain.GeneralError, http.StatusInternalServerError)
                        return
                }
                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(http.StatusOK)
                responseMessage := []string{"Win added =), Keep it up."}
                w.Write(buildResponse(responseMessage))
        } else {
                osin.OutputJSON(resp, w, r)
        }
}

// AddFail is the endpoint to submit the fail of the day to the user fail list.
func AddFail(w http.ResponseWriter, r *http.Request) {
        authorizationServer := osin.NewServer(osin.NewServerConfig(), mongo.NewMongoAPIClient())
        resp := authorizationServer.NewResponse()
        defer resp.Close()

        infoRequest := authorizationServer.HandleInfoRequest(resp, r)

        if infoRequest != nil {
                api := domain.NewAPI(mongo.NewMongoAPIClient())
                err := api.AddFail()
                if err != nil {
                        http.Error(w, domain.GeneralError, http.StatusInternalServerError)
                        return
                }
                w.Header().Set("Content-Type", "application/json; charset=UTF-8")
                w.WriteHeader(http.StatusOK)
                responseMessage := []string{"Fail added =(, Sorry to hear that."}
                w.Write(buildResponse(responseMessage))
        } else {
                osin.OutputJSON(resp, w, r)
        }
}

func buildResponse(response interface{}) []byte {

        resp, err := json.Marshal(response)
        if err != nil {

        }
        return resp
}
