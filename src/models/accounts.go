package models

import (
	_ "errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Account struct {
	ID    int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	DOB string `json:"dob"`
	Sex string `json:"sex"`
}

func CreateAccount(account Account) (error) {

	    //set the current time
	    var currentTime = time.Now()

	    //insert into DB
	_ , err := db.Exec("insert into user(first_name,last_name,email,mobile,date_of_birth,sex,create_time,update_time) values(?,?,?,?,?,?,?,?)",
		account.FirstName,
		account.LastName,
		account.Email,
		account.Mobile,
		account.DOB,
		account.Sex,
		currentTime,
		currentTime)

	return err

}

func GetAccount(accountId int) (Account) {
	account := Account{}

	fmt.Println("Account id (%s)",accountId)

	queryString := "select user_id,first_name,last_name,email,mobile,date_of_birth,sex from user where user_id = ? "

	db.QueryRow(queryString, accountId).Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Email,
		&account.Mobile,
		&account.DOB,
		&account.Sex)

	fmt.Println("Account info:",account)
    return account
}
