package main

import "to-do-api/database"

func main() {
	dsn := "host=localhost user=myuser password=mypassword dbname=go-db port=5433 sslmode=disable"
	// dsn will be got from config.go
	database.ConnectDatabase(dsn)
}
