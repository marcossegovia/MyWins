package mongo

import (
	"github.com/RangelReale/osin"
	"gopkg.in/mgo.v2/bson"
)

const (
	CLIENT_COL = "clients"
	AUTHORIZE_COL = "authorizations"
)

func (client MongoApiClient) Clone() osin.Storage {
	return client
}

func (client MongoApiClient) Close() {

}

func (client MongoApiClient) GetClient(id string) (osin.Client, error) {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	osinClient := new(osin.DefaultClient)
	clients := session.DB(TOKEN_DB_NAME).C(CLIENT_COL)
	err := clients.FindId(id).One(osinClient)

	return osinClient, err
}

func (client MongoApiClient) SetClient(id string, osinClient osin.Client) error {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	clients := session.DB(TOKEN_DB_NAME).C(CLIENT_COL)
	_, err := clients.UpsertId(id, osinClient)
	return err
}

func (client MongoApiClient) SaveAuthorize(data *osin.AuthorizeData) error {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	authorizations := session.DB(TOKEN_DB_NAME).C(AUTHORIZE_COL)
	_, err := authorizations.UpsertId(data.Code, AuthorizeDataFromOSIN(data))
	return err
}

func (client MongoApiClient) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	var authData AuthorizeData
	authorizations := session.DB(TOKEN_DB_NAME).C(AUTHORIZE_COL)
	err := authorizations.FindId(code).One(&authData)
	if err != nil {
		return nil, err
	}
	pam, err := authData.AuthorizeDataToOSIN(client)
	return pam, err
}

func (client MongoApiClient) RemoveAuthorize(code string) error {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	authorizations := session.DB(TOKEN_DB_NAME).C(AUTHORIZE_COL)
	return authorizations.RemoveId(code)
}

func (client MongoApiClient) SaveAccess(data *osin.AccessData) error {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	accesses := session.DB(TOKEN_DB_NAME).C(ACCESS_COL)
	_, err := accesses.UpsertId(data.AccessToken, AccessDataFromOSIN(data))
	return err
}

func (client MongoApiClient) LoadAccess(token string) (*osin.AccessData, error) {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	var accData AccessData
	accesses := session.DB(TOKEN_DB_NAME).C(ACCESS_COL)
	err := accesses.FindId(token).One(&accData)
	if err != nil {
		return nil, err
	}
	return accData.AccessDataToOSIN(client)
}

func (client MongoApiClient) RemoveAccess(token string) error {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	accesses := session.DB(TOKEN_DB_NAME).C(ACCESS_COL)
	return accesses.RemoveId(token)
}

func (client MongoApiClient) LoadRefresh(token string) (*osin.AccessData, error) {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	var accData AccessData

	accesses := session.DB(TOKEN_DB_NAME).C(ACCESS_COL)
	err := accesses.Find(bson.M{REFRESH_TOKEN: token}).One(&accData)
	if err != nil {
		return nil, err
	}
	return accData.AccessDataToOSIN(client)
}

func (client MongoApiClient) RemoveRefresh(token string) error {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	accesses := session.DB(TOKEN_DB_NAME).C(ACCESS_COL)
	return accesses.Remove(bson.M{REFRESH_TOKEN: token})
}
