package models

import (
	_ "errors"
	_ "github.com/go-sql-driver/mysql"
)


type Account struct {
	ID    int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
}


func CreateAccount(account Account) (error) {

	_ , err := db.Exec("insert into account(first_name,last_name,email,mobile) values(?,?,?,?)",
		account.FirstName,
		account.LastName,
		account.Email,
		account.Mobile)

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
		&account.Mobile)

	if result != nil {
		return Account{}
	}

	return account
}
