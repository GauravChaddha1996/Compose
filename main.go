package main

import (
	"compose/commons"
	"compose/commons/logger"
	"compose/endpoints/article"
	"compose/endpoints/comments"
	"compose/endpoints/like"
	"compose/endpoints/user"
	"compose/middlewares"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	logger.InitLogger()
	logger.InfoPreNewLine("Starting project compose")
	loadConfig()
	db := openDB()
	commons.Init(db)
	startServer()
}

func loadConfig() {
	logger.Info("Loading config")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	commons.PanicIfError(err)
	logger.Info("Config loaded")
}

func openDB() *gorm.DB {
	logger.Info("Starting database connection")
	dbName := viper.GetString("db.name")
	dbParams := viper.GetString("db.params")
	dbConnectionConfig := viper.GetString("db.username") + ":" + viper.GetString("db.password") +
		"@tcp(" + viper.GetString("db.host") + ":" + viper.GetString("db.port") + ")"
	dbArgs := dbConnectionConfig + "/" + dbName + "?" + dbParams
	db, err := gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	commons.PanicIfError(err)
	logger.Info("Database connection established")
	return db
}

func startServer() {
	logger.InfoPostNewLine("Starting server at http://localhost:8000")
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
	router.Use(middlewares.TimeoutHandlingMiddleware)
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
