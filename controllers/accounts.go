package controllers

import (
	"encoding/json"
	"fmt"
	m "identityserv/models"
	u "identityserv/utils"
	"io"
	"net/http"
	"path"
	"strconv"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	account := m.Account {}
	out := make([]byte,1024)
	bodyLen, err := r.Body.Read(out)

	if err != io.EOF {
		fmt.Println(err.Error())
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}


	err = json.Unmarshal(out[:bodyLen],&account)

	if err != nil {
		w.Write([]byte("{error:" + err.Error() + "}"))
		return
	}

	err = m.CreateAccount(account)


	if err == nil {
		u.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
	} else {
		u.RespondWithError(w, http.StatusInternalServerError,  "oops! internal server error")
	}
}

func GetAccount(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	account := m.Account {}
	accountId, err := strconv.Atoi(path.Base(r.RequestURI))

	if err == nil {
		account = m.GetAccount(accountId)
	}

	if account.ID == 0 {
		u.RespondwithJSON(w, http.StatusNoContent, map[string]string {"message":"sorry! account not found"})
	} else {
		enc := json.NewEncoder(w)
		enc.Encode(&account)
	}

}
