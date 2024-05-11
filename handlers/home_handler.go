package handlers

import (
	"net/http"
	"todo-gotth/database"
	"todo-gotth/views"

	"github.com/gin-gonic/gin"
)


type HomeHtmlHandler struct {
	db *database.SQLiteDatabase
}

func (h *HomeHtmlHandler) GetHome(ctx *gin.Context)  {
	
	todoList, err := h.db.All()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	
	ctx.HTML(http.StatusOK, "", views.Home(todoList))
}
