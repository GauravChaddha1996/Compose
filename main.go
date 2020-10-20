package main

import (
	"compose/article"
	"compose/comments"
	"compose/commons"
	"compose/like"
	"compose/middlewares"
	"compose/serviceContracts"
	"compose/user"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	println("")
	log.Println("Starting project compose")
	loadConfig()
	db := openDB()
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
	dbArgs := dbConnectionConfig + "/" + dbName + "?" + dbParams
	db, err := gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	commons.PanicIfError(err)
	log.Print("Database connection established")
	return db
}

func initPackages(db *gorm.DB) {
	// Init with common things
	commons.Init(db)
	user.Init(db)
	article.Init(db)
	like.Init(db)
	comments.Init(db)

	// Save all service impls
	serviceContracts.Init(user.GetServiceContractImpl(), article.GetServiceContractImpl(), like.GetServiceContractImpl(), comments.GetCommentServiceContractImpl())

	// Attach all service impls to other services
	user.SetServiceContractImpls(serviceContracts.GetArticleServiceContract(), serviceContracts.GetLikeServiceContract())
	article.SetServiceContractImpl(serviceContracts.GetUserServiceContract(), serviceContracts.GetCommentServiceContract(), serviceContracts.GetLikeServiceContract())
	like.SetServiceContractImpl(serviceContracts.GetArticleServiceContract(), serviceContracts.GetUserServiceContract())
	comments.SetServiceContractImpl(serviceContracts.GetArticleServiceContract(), serviceContracts.GetUserServiceContract())
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
	router.Use(middlewares.RequestLoggingMiddleware)
	router.Use(middlewares.ExtractCommonModelMiddleware)
	router.Use(middlewares.GeneralSecurityMiddleware)
	router.Use(middlewares.CommonResponseHeadersMiddleware)
	router.Use(middlewares.ResponseLoggingMiddleware)
}

func addApiRoutes(router *mux.Router) {
	router.HandleFunc("/", home)
	user.AddSubRoutes(router.PathPrefix("/user").Subrouter())
	article.AddSubRoutes(router.PathPrefix("/article").Subrouter())
	like.AddSubRoutes(router.PathPrefix("/article").Subrouter())
	comments.AddSubRoutes(router.PathPrefix("/article").Subrouter())
}

func home(writer http.ResponseWriter, _ *http.Request) {
	_, _ = writer.Write([]byte("Hello world"))
}
