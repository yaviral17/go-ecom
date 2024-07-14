package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/yaviral17/go-ecom/cmd/api"
	"github.com/yaviral17/go-ecom/cmd/config"
	"github.com/yaviral17/go-ecom/cmd/db"
)

func main() {
	// Call the function from the package
	// and print the result
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBuser,
		Passwd:               config.Envs.DBpass,
		Addr:                 config.Envs.DBAddr,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Successfully connected")
}
