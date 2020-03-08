package main

import (
	"compose/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {
	db := _openDB()
	defer db.Close()
	_initPackages(db)
	_startServer()
}

func _openDB() *gorm.DB {
	log.Print("Starting database connection")
	dbName := "compose"
	dbParams := "charset=utf8&parseTime=True&loc=Local"
	dbConnectionConfig := "root:@tcp(127.0.0.1:3306)"
	db, err := gorm.Open("mysql", dbConnectionConfig+"/"+dbName+"?"+dbParams)
	if err != nil {
		log.Print("Cannot open db")
		panic(err)
	}
	log.Print("Database connection established")
	return db
}

func _initPackages(db *gorm.DB) {
	user.Init(db)
}

func _startServer() {
	log.Print("Starting server at http://localhost:8000")
	err := http.ListenAndServe(":8000", _getMainRouter())
	if err != nil {
		log.Print("Error starting server. Returning")
		panic(err)
	}
}

func _getMainRouter() *mux.Router {
	router := mux.NewRouter()
	_addApiRoutes(router)
	return router
}

func _addApiRoutes(router *mux.Router) {
	router.HandleFunc("/", home)
	user.AddApiRoutes(router)
}

func home(writer http.ResponseWriter, _ *http.Request) {
	_, _ = writer.Write([]byte("Hello world"))
}
