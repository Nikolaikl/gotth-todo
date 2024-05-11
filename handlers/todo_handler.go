package handlers

import (
	"net/http"
	"strconv"
	"todo-gotth/database"
	"todo-gotth/models"
	"todo-gotth/views"

	"github.com/gin-gonic/gin"
)

type ToDoHTMLHandler struct {
	Dbconn *database.SQLiteDatabase
}

func (h *ToDoHTMLHandler) GetAll(ctx *gin.Context) {
	todoList, err := h.Dbconn.All()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.HTML(http.StatusOK, "", views.ToDoTable(todoList))
}

func (h *ToDoHTMLHandler) Create(ctx *gin.Context) {
	description := ctx.PostForm("description")
	todo := models.ToDo{
		Description: description,
		Completed:   false,
	}
	_, err := h.Dbconn.Create(todo)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	todoList, err := h.Dbconn.All()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.HTML(http.StatusOK, "", views.ToDoTable(todoList))
}

func (h *ToDoHTMLHandler) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	todoToUpdate, err := h.Dbconn.GetByID(int64(intID))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	todoToUpdate.Completed = !todoToUpdate.Completed
	updatedTodo, err := h.Dbconn.Update(int64(intID), *todoToUpdate)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.HTML(http.StatusOK, "", views.ToDoTableItem(*updatedTodo))
}

func (h *ToDoHTMLHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = h.Dbconn.Delete(int64(intID))
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}
