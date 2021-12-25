package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/book/domain"
	"github.com/book/service"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func sanityCheck() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined...")
	}
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	dbClient := getDbClient()

	bookRepositoryDb := domain.NewBookRepositoryDb(dbClient)

	bh := BookHandlers{service: service.NewBookService(bookRepositoryDb)}
	router.HandleFunc("/books", bh.getAllBook).
		Methods(http.MethodGet, "OPTIONS").
		Name("GetAllBooks")

	log.Fatal(http.ListenAndServe(":8001", router))

}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	//client.SetMax(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
