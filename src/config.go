package x90

import (
	"encoding/json"
	"log"
	"os"
)

type Params struct {
	Token    string `json : "Token"`
	Database string `json : "Database"`
}

func GetConfig() Params {
	var params Params
	file, err := os.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &params)

	if err != nil {
		log.Fatal(err)
	}

	return params

}
