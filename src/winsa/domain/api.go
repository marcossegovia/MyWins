package domain

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

type MyWinsAPI struct{}

var api *MyWinsAPI

func NewApi() *MyWinsAPI {
	if api != nil {
		return api
	}
	api := new(MyWinsAPI)
	return api
}

func (api *MyWinsAPI) FindAllWins() (*win, error) {

	wins, err := readFileToDomain()
	return wins, err
}

func (api *MyWinsAPI) AddWin() error {
	wins, err := readFileToDomain()
	if err != nil {
		return err
	}
	wins.Success = append(wins.Success, time.Now().Format(time.RFC3339))
	err = writeDomainToFile(wins)
	if err != nil {
		return err
	}

	return nil
}

func (api *MyWinsAPI) AddFail() error {
	wins, err := readFileToDomain()
	if err != nil {
		return err
	}
	wins.Fails = append(wins.Fails, time.Now().Format(time.RFC3339))
	err = writeDomainToFile(wins)
	if err != nil {
		return err
	}

	return nil
}

func readFileToDomain() (*win, error) {
	absPath, _ := filepath.Abs("files/wins.json")
	jsonFile, err := os.Open(absPath)

	if err != nil {
		log.Print("Error when trying to open json file", err.Error())
		return nil, err
	}

	wins := new(win)
	jsonParser := json.NewDecoder(jsonFile)
	if err = jsonParser.Decode(&wins); err != nil {
		log.Print("Error on parsing json file", err.Error())
		return nil, err
	}
	return wins, nil
}

func writeDomainToFile(w *win) error {
	serialized_wins, err := json.Marshal(w)
	if err != nil {
		return err
	}
	absPath, _ := filepath.Abs("files/wins.json")
	err = ioutil.WriteFile(absPath, serialized_wins, 0644)
	if err != nil {
		log.Print("Could not write on the file", err.Error())
		return err
	}

	return nil
}
