package mongo

import (
        "github.com/RangelReale/osin"
        "gopkg.in/mgo.v2/bson"
)

const (
        CLIENT_COL = "clients"
        AUTHORIZE_COL = "authorizations"
)

// Clone implementation of osin.Storage.Clone using MongoDB
func (client MongoAPIClient) Clone() osin.Storage {
        return client
}

// Close implementation of osin.Storage.Close using MongoDB
func (client MongoAPIClient) Close() {

}

// GetClient implementation of osin.Storage.GetClient using MongoDB
func (client MongoAPIClient) GetClient(id string) (osin.Client, error) {
        session := client.dbClient.Session.Copy()
        defer session.Close()

        osinClient := new(osin.DefaultClient)
        clients := session.DB(TOKEN_DB_NAME).C(CLIENT_COL)
        err := clients.FindId(id).One(osinClient)

        return osinClient, err
}

// SetClient implementation of osin.Storage.SetClient using MongoDB
func (client MongoAPIClient) SetClient(id string, osinClient osin.Client) error {
        session := client.dbClient.Session.Copy()
        defer session.Close()

        clients := session.DB(TOKEN_DB_NAME).C(CLIENT_COL)
        _, err := clients.UpsertId(id, osinClient)
        return err
}

// SaveAuthorize implementation of osin.Storage.SaveAuthorize using MongoDB
func (client MongoAPIClient) SaveAuthorize(data *osin.AuthorizeData) error {
        session := client.dbClient.Session.Copy()
        defer session.Close()

        authorizations := session.DB(TOKEN_DB_NAME).C(AUTHORIZE_COL)
        _, err := authorizations.UpsertId(data.Code, AuthorizeDataFromOSIN(data))
        return err
}

// LoadAuthorize implementation of osin.Storage.LoadAuthorize using MongoDB
func (client MongoAPIClient) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
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

// RemoveAuthorize implementation of osin.Storage.RemoveAuthorize using MongoDB
func (client MongoAPIClient) RemoveAuthorize(code string) error {
        session := client.dbClient.Session.Copy()
        defer session.Close()

        authorizations := session.DB(TOKEN_DB_NAME).C(AUTHORIZE_COL)
        return authorizations.RemoveId(code)
}

// SaveAccess implementation of osin.Storage.SaveAccess using MongoDB
func (client MongoAPIClient) SaveAccess(data *osin.AccessData) error {
        session := client.dbClient.Session.Copy()
        defer session.Close()

        accesses := session.DB(TOKEN_DB_NAME).C(ACCESS_COL)
        _, err := accesses.UpsertId(data.AccessToken, AccessDataFromOSIN(data))
        return err
}

// LoadAccess implementation of osin.Storage.LoadAccess using MongoDB
func (client MongoAPIClient) LoadAccess(token string) (*osin.AccessData, error) {
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

// RemoveAccess implementation of osin.Storage.RemoveAccess using MongoDB
func (client MongoAPIClient) RemoveAccess(token string) error {
        session := client.dbClient.Session.Copy()
        defer session.Close()

        accesses := session.DB(TOKEN_DB_NAME).C(ACCESS_COL)
        return accesses.RemoveId(token)
}

// LoadRefresh implementation of osin.Storage.LoadRefresh using MongoDB
func (client MongoAPIClient) LoadRefresh(token string) (*osin.AccessData, error) {
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

// RemoveRefresh implementation of osin.Storage.RemoveRefresh using MongoDB
func (client MongoAPIClient) RemoveRefresh(token string) error {
        session := client.dbClient.Session.Copy()
        defer session.Close()

        accesses := session.DB(TOKEN_DB_NAME).C(ACCESS_COL)
        return accesses.Remove(bson.M{REFRESH_TOKEN: token})
}
