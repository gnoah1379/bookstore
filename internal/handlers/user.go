package handlers

import (
	"bookstore/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userSvc service.UserService
}

func NewUserHandler(userSvc service.UserService) *UserHandler {
	return &UserHandler{userSvc: userSvc}
}

func (h *UserHandler) Register(c *gin.Context) {
	var request service.UserRegistration
	err := c.ShouldBind(&request)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	userInfo, err := h.userSvc.Register(c.Request.Context(), request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, userInfo)
}
