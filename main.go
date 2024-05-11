package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"todo.de/home/database"
)

const fileName = "sqlite.db"

func main () {
	conn, err := sql.Open("sqlite3", fileName)
	
	if err != nil {
		log.Fatal(err)
	}
	
	db := database.NewSQLiteDatabase(conn)
	
}


