package handlers

import (
	"net/http"
	"todo-gotth/database"
	"todo-gotth/views"

	"github.com/gin-gonic/gin"
)


type HomeHTMLHandler struct {
	Dbconn *database.SQLiteDatabase
}

func (h *HomeHTMLHandler) GetHome(ctx *gin.Context)  {
	
	todoList, err := h.Dbconn.All()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	ctx.HTML(http.StatusOK, "", views.Home(todoList))
}
