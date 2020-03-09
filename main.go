package main

import (
	"compose/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	_loadConfig()
	db := _openDB()
	defer db.Close()
	_initPackages(db)
	_startServer()
}

func _loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func _openDB() *gorm.DB {
	log.Print("Starting database connection")
	dbName := viper.GetString("db.name")
	dbParams := viper.GetString("db.params")
	dbConnectionConfig := viper.GetString("db.username") + ":" + viper.GetString("db.password") +
		"@tcp(" + viper.GetString("db.host") + ":" + viper.GetString("db.port") + ")"
	dbImplementation := viper.GetString("db.implementation")
	dbArgs := dbConnectionConfig + "/" + dbName + "?" + dbParams
	db, err := gorm.Open(dbImplementation, dbArgs)
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
	serverPort := ":" + viper.GetString("server.port")
	err := http.ListenAndServe(serverPort, _getMainRouter())
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
