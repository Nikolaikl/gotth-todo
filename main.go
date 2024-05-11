package main

import (
	"database/sql"
	"fmt"
	"log"
	"todo-gotth/database"
	"todo-gotth/handlers"

	"github.com/gin-gonic/gin"
)

const fileName = "sqlite.db"

func main() {
	conn, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}

	db := database.NewSQLiteDatabase(conn)

	if err := db.Migrate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Migrated")

	router := gin.Default()
	router.HTMLRender = &RenderedTemplate{}

	homeHtmlHandler := handlers.HomeHtmlHandler{db: db}
	todoHtmlHandler := handlers.TodoHtmlHandler{db: db}
	
	router.GET("/", homeHtmlHandler)
	router.GET("/todo", todoHtmlHandler)
	router.POST("/todo", todoHtmlHandler)
	router.PATCH("/todo", todoHtmlHandler)
	router.DELETE("/todo", todoHtmlHandler)
	
	router.Run(":8000")
}
