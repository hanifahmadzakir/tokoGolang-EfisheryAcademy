package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

//struct DB Config
type DBConfig struct{
	DBHost		string
	DBUser		string
	DBPassword	string
	DBName		string
	DBPort		string
}

// fungsi untuk inisialisasi
func (server *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	var err error 
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("Failed connecting to database")
	}

	server.Router = mux.NewRouter()
	server.initializeRoutes()

}

// fungsi untuk print port listener
func (server *Server) Run(addr string) {
	fmt.Printf("Listening to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

// fungsi untuk cek .env jika tidak ada akan di fallback ke default
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// fungsi untuk running server
func Run() {
	var server = Server{}
	var appConfig = AppConfig{}
	var dbConfig = DBConfig{}

	//cek file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "Toko Golang")
	appConfig.AppEnv = getEnv("APP_ENV", "PENGEMBANGAN")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBName = getEnv("DB_NAME", "tokogolang_db")
	dbConfig.DBUser = getEnv("DB_USER", "postgres")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "RootAccess1108")
	dbConfig.DBPort	= getEnv("DB_PORT", "5432")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)

}
