package main

import (
	controller "../controllers"
	util "../utils"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	_ "io"
	"net/http"
	_ "strconv"
)

var router *chi.Mux
var db *sql.DB



func routers() *chi.Mux {
	router.Get("/v1/account/{id}", controller.GetAccount)
	router.Post("/v1/account", controller.CreateAccount)
	router.Post("/v1/test", controller.ReturnTestString)

	return router
}


func main() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	fmt.Println("Starting server...")
	http.ListenAndServe(":8000", util.Logger(routers()))
}