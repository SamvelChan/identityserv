package models

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *sql.DB

func init() {
	// load the variables from .env file
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbUser := os.Getenv("db_user")
	dbPass := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	fmt.Println("Db User(%s",dbUser);
	fmt.Println("Db Pass(%s",dbPass);
	fmt.Println("Db Name(%s",dbName);
	fmt.Println("Db Host(%s",dbHost);
	fmt.Println("Db Port(%s",dbPort);

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)

	var err error
	db, err = sql.Open("mysql", dbSource)

	if err != nil {
		log.Println(err.Error())
	}
}

