package models

import (
	_ "errors"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Account struct {
	ID    int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	DOB string `json:"mobile"`
	Sex string `json:"mobile"`
}

func CreateAccount(account Account) (error) {

	    //set the current time
	    var currentTime = time.Now()

	    //insert into DB
	_ , err := db.Exec("insert into account(first_name,last_name,email,mobile,date_of_birth,sex,create_time,update_time) values(?,?,?,?,?,?,?,?)",
		account.FirstName,
		account.LastName,
		account.Email,
		account.Mobile,
		account.Sex,
		account.DOB,
		currentTime,
		currentTime)

	return err

}

func GetAccount(accountId int) (Account) {
	account := Account{}

	queryString := "select * from account where id = ? "

	result := db.QueryRow(queryString, accountId).Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Email,
		&account.Mobile,
		&account.DOB,
		&account.Sex)

	if result != nil {
		return Account{}
	}

	return account
}
