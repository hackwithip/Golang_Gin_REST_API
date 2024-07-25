package inits

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPgDB() {
	
	//we read our .env file
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	dbname := os.Getenv("PG_DB_NAME")
	pass := os.Getenv("PG_PASSWORD")

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, dbname)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
    
    if err!= nil {
		panic("Failed to connect to Database!")
	}
    
	fmt.Println("Connected to the database!")
    
    DB = db

}