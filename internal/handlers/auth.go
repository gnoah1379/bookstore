package handlers

import (
	"bookstore/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authSvc service.AuthService
}

func NewAuthHandler(authSvc service.AuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request service.UserLoginRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	userInfo, err := h.authSvc.Login(c.Request.Context(), request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, userInfo)
}
