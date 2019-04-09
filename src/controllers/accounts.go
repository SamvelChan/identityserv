package controllers

import (
	model "../models"
	util "../utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	account := model.Account {}
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

	err = model.CreateAccount(account)


	if err == nil {
		util.RespondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
	} else {
		util.RespondWithError(w, http.StatusInternalServerError,  "oops! internal server error")
	}
}

func GetAccount(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	account := model.Account {}
	accountId, err := strconv.Atoi(path.Base(r.RequestURI))

	if err == nil {
		account = model.GetAccount(accountId)
	}

	if account.ID == 0 {
		util.RespondwithJSON(w, http.StatusNoContent, map[string]string {"message":"sorry! account not found"})
	} else {
		enc := json.NewEncoder(w)
		enc.Encode(&account)
	}

}

func ReturnTestString(w http.ResponseWriter, r *http.Request) {
	util.RespondwithJSON(w, http.StatusOK, map[string]string {"message":"test String"})
}