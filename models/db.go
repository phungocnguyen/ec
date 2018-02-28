package models

import (
	"database/sql"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

const DATASOURCE_URL = "postgres://%v:%v@%v/%v?sslmode=%v"

type Config struct {
	Dbname   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Sslmode  string `json:"sslmode"`
	Location string `json:"location"`
	Schema   string `json:"currentSchema"`
	Port	 string `json:"port"`
}

func getConfig() Config {
	raw, err := ioutil.ReadFile("./configs/config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Config
	errj := json.Unmarshal(raw, &c)
	if errj != nil {
		fmt.Println("error parsing json input", err)
	}
	return c[0]
}

func NewDB() (*sql.DB, error) {
	/*
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	hostname := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	*/
	config := getConfig()

	// build our data source url
	dataSourceName := fmt.Sprintf(DATASOURCE_URL, config.Username, config.Password, config.Location, config.Dbname, config.Sslmode)
	fmt.Println("data source",dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
