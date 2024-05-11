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

	dbconn := database.NewSQLiteDatabase(conn)

	if err := dbconn.Migrate(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Migrated")

	router := gin.Default()
	router.HTMLRender = &RenderedTemplate{}

	homeHTMLHandler := handlers.HomeHTMLHandler{Dbconn: dbconn}
	todoHTMLHandler := handlers.ToDoHTMLHandler{Dbconn: dbconn}
	
	router.GET("/", homeHTMLHandler.GetHome)
	router.GET("/todo", todoHTMLHandler.GetAll)
	router.POST("/todo", todoHTMLHandler.Create)
	router.PATCH("/todo", todoHTMLHandler.Update)
	router.DELETE("/todo", todoHTMLHandler.Delete)
	
	router.Run(":42069")
}
