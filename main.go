package main

import (
	"compose/article"
	"compose/commons"
	"compose/serviceContracts"
	"compose/user"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	println("")
	log.Println("Starting project compose")
	loadConfig()
	db := openDB()
	defer db.Close()
	initPackages(db)
	startServer()
}

func loadConfig() {
	log.Println("Loading config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	commons.PanicIfError(err)
	log.Println("Config loaded")
}

func openDB() *gorm.DB {
	log.Print("Starting database connection")
	dbName := viper.GetString("db.name")
	dbParams := viper.GetString("db.params")
	dbConnectionConfig := viper.GetString("db.username") + ":" + viper.GetString("db.password") +
		"@tcp(" + viper.GetString("db.host") + ":" + viper.GetString("db.port") + ")"
	dbImplementation := viper.GetString("db.implementation")
	dbArgs := dbConnectionConfig + "/" + dbName + "?" + dbParams
	db, err := gorm.Open(dbImplementation, dbArgs)
	commons.PanicIfError(err)
	log.Print("Database connection established")
	return db
}

func initPackages(db *gorm.DB) {
	// Init with common things
	commons.Init(db)
	user.Init(db)
	article.Init(db)

	// Save all service impls
	serviceContracts.Init(user.GetServiceContractImpl())

	// Attach all service impls to other services
	article.SetServiceContractImpl(serviceContracts.GetUserServiceContract())
}

func startServer() {
	log.Println("Starting server at http://localhost:8000")
	println("")
	serverPort := ":" + viper.GetString("server.port")
	err := http.ListenAndServe(serverPort, getMainRouter())
	commons.PanicIfError(err)
}

func getMainRouter() *mux.Router {
	router := mux.NewRouter()
	addMiddlewares(router)
	addApiRoutes(router)
	return router
}

func addMiddlewares(router *mux.Router) {
	router.Use(commons.RequestLoggingMiddleware)
	router.Use(commons.ExtractCommonModelMiddleware)
	router.Use(commons.GeneralSecurityMiddleware)
	router.Use(commons.CommonResponseHeadersMiddleware)
	router.Use(commons.ResponseLoggingMiddleware)
}

func addApiRoutes(router *mux.Router) {
	router.HandleFunc("/", home)
	user.AddSubRoutes(router.PathPrefix("/user").Subrouter())
	article.AddSubRoutes(router.PathPrefix("/article").Subrouter())
}

func home(writer http.ResponseWriter, _ *http.Request) {
	_, _ = writer.Write([]byte("Hello world"))
}
