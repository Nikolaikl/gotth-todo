
package handlers

import (
	"net/http"
	"strconv"
	"todo-gotth/database"
	"todo-gotth/models"
	"todo-gotth/views"

	"github.com/gin-gonic/gin"
)

type TodoHtmlHandler struct {
	db *database.SQLiteDatabase
}

func (h *TodoHtmlHandler) GetAll(ctx *gin.Context) {
	todoList, err := h.db.All()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.HTML(http.StatusOK, "", views.TodoTable(todoList))
}

func (h *TodoHtmlHandler) Create(ctx *gin.Context) {
	description := ctx.PostForm("description")
	todo := models.ToDo{
		Description: description,
		Completed:   false,
	}
	_, err := h.db.Create(todo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	todoList, err := h.db.All()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.HTML(http.StatusOK, "", views.TodoTable(todoList))
}

func (h *TodoHtmlHandler) Update(ctx *gin.Context) {
	// this only updates the completed status
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	todoToUpdate, err := h.db.GetByID(int64(intId))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	todoToUpdate.Completed = !todoToUpdate.Completed
	updatedTodo, err := h.db.Update(int64(intId), *todoToUpdate)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "", views.TodoTableItem(*updatedTodo))
}

func (h *TodoHtmlHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = h.db.Delete(int64(intId))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}
