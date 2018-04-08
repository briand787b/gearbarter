package models

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
)

var db *sqlx.DB

func init() {
	credentials := struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Database string `json:"database"`
	}{}

	data, err := ioutil.ReadFile("models/database.json")
	if err != nil {
		log.Fatal("could not open database config file: ", err)
	}

	if err := json.Unmarshal(data, &credentials); err != nil {
		log.Fatal("could not marshal database.json to credentials struct: ", err)
	}

	db, err = sqlx.Connect("postgres",
		// CHANGE SSLMODE WHEN I GET REAL CERTS
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
			credentials.Username,
			credentials.Password,
			credentials.Database,
		),
	)

	if err != nil {
		log.Fatal("error connecting to db:  ", err)
	}

}
