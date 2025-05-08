package handlers

import (
	"bookstore/internal/model"
	"bookstore/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WhishlistHandler struct {
	wishlistService service.WhishlistService
}

func NewWhishlistHandler(wishlistService service.WhishlistService) *WhishlistHandler {
	return &WhishlistHandler{
		wishlistService: wishlistService,
	}
}

func (h *WhishlistHandler) GetWhishlistByUserId(ctx *gin.Context) {
	whishlist, err := h.wishlistService.GetWhishlistByUserId(ctx.Request.Context(), ctx.Param("user id"))
	if err != nil {
		ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(ctx, whishlist)
}

func (h *WhishlistHandler) AddWhishItem(ctx *gin.Context) {
	var request model.Whishlist
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	whishlist, err := h.wishlistService.CreateWhishItem(ctx.Request.Context(), request)
	if err != nil {
		ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(ctx, whishlist)
}

func (h *WhishlistHandler) DeleteWhishItemByWishItemId(ctx *gin.Context) {
	var request model.Whishlist
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ResponseError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	whishlist, err := h.wishlistService.DeleteWhishItemByWishItemId(ctx.Request.Context(), ctx.Param("item id"), request)
	if err != nil {
		ResponseError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(ctx, whishlist)
}
