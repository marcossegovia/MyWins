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

	clients := session.DB(TOKEN_DB_NAME).C(CLIENT_COL)
	osinClient := new(osin.DefaultClient)
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
	_, err := authorizations.UpsertId(data.Code, data)
	return err
}

func (client MongoApiClient) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	authorizations := session.DB(TOKEN_DB_NAME).C(AUTHORIZE_COL)
	authData := new(osin.AuthorizeData)
	err := authorizations.FindId(code).One(authData)
	return authData, err
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
	_, err := accesses.UpsertId(data.AccessToken, data)
	return err
}

func (client MongoApiClient) LoadAccess(token string) (*osin.AccessData, error) {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	accesses := session.DB(TOKEN_DB_NAME).C(ACCESS_COL)
	accData := new(osin.AccessData)
	err := accesses.FindId(token).One(accData)
	return accData, err
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

	accesses := session.DB(TOKEN_DB_NAME).C(ACCESS_COL)
	accData := new(osin.AccessData)
	err := accesses.Find(bson.M{REFRESH_TOKEN: token}).One(accData)
	return accData, err
}

func (client MongoApiClient) RemoveRefresh(token string) error {
	session := client.dbClient.Session.Copy()
	defer session.Close()

	accesses := session.DB(TOKEN_DB_NAME).C(ACCESS_COL)
	return accesses.Update(bson.M{REFRESH_TOKEN: token}, bson.M{
		"$unset": bson.M{
			REFRESH_TOKEN: 1,
		}})
}
