package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ResponseError(c *gin.Context, code int, message string) {
	err := &ErrorResponse{Code: code, Message: message}
	c.JSON(err.Code, err)
}

type SuccessResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
	})
}
