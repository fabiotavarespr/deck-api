package handlers

import (
	"net/http"

	internalerrors "deck-api/errors"

	"github.com/gin-gonic/gin"
)

const ErrDefaultMessage = "Something went wrong, contact the system administrator to obtain more details"

type Handler interface {
	Routes(router *gin.Engine)
}

func HandleErr(context *gin.Context, err error) {
	switch err.(type) {
	case internalerrors.ErrInvalidEntry, internalerrors.ErrInsufficientResources:
		context.JSON(http.StatusBadRequest, err.Error())
	case internalerrors.ErrNotFound:
		context.JSON(http.StatusNotFound, err.Error())
	default:
		context.JSON(http.StatusInternalServerError, ErrDefaultMessage)
	}
}
