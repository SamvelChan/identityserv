package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	c "identityserv/controllers"
	u "identityserv/utils"
	_ "io"
	"net/http"
	_ "strconv"
)

var router *chi.Mux
var db *sql.DB



func routers() *chi.Mux {
	router.Get("/v1/account/{id}", c.GetAccount)
	router.Post("/v1/account", c.CreateAccount)

	return router
}


func main() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)

	http.ListenAndServe(":8000", u.Logger(routers()))
}