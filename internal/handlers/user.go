package handlers

import (
	"bookstore/internal/model"
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

func (h *UserHandler) ListUsers(c *gin.Context) {
	user, err := h.userSvc.ListAllUser(c.Request.Context())
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
	}
	ResponseSuccess(c, user)
}

func (h *UserHandler) SearchUser(c *gin.Context) {
	var id = c.Param("id")
	user, err := h.userSvc.SearchUserById(c.Request.Context(), id)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var request model.User
	var id = c.Param("id")
	err := c.ShouldBind(&request)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.userSvc.UpdateUserById(c.Request.Context(), id, request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var id = c.Param("id")
	user, err := h.userSvc.DeleteUserById(c.Request.Context(), id)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, user)
}
