package main

import (
	"log"
	"net/http"
	"path"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/pedroh-reis/valu-safe/backend/src/config"
)

func main() {
	log.Println("Loading environment variables")
	loadEnvironment()

	log.Println("Getting db connection")
	db := GetPostgresDbConnection()

	log.Println("Initializing repositories")
	repository := NewRepository(db)

	log.Println("Initializing the API routers")
	NewRouter(repository).init()

	log.Println("Starting the server")
	serverConfig := config.NewServerConfig()
	log.Printf("Listening on http://%s", serverConfig.GetAddress())

	err := http.ListenAndServe(serverConfig.GetAddress(), nil)
	if err != nil {
		log.Println(err)
	}
}

func loadEnvironment() {
	_, filename, _, _ := runtime.Caller(1)
	envPath := path.Join(filename, "..", "..", ".env")

	godotenv.Load(envPath)
}
