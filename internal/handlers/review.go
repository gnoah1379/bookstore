package handlers

import (
	"bookstore/internal/model"
	"bookstore/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReviewHandler struct {
	reviewSvc service.ReviewService
}

func NewReviewHandler(reviewSvc service.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		reviewSvc: reviewSvc,
	}
}

// review
func (r *ReviewHandler) CreateReview(c *gin.Context) {
	var request model.Review
	err := c.ShouldBindJSON(&request)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	review, err := r.reviewSvc.CreateReview(c.Request.Context(), request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, review)
}

func (r *ReviewHandler) ListAllReview(c *gin.Context) {
	reviews, err := r.reviewSvc.ListAllReview(c.Request.Context())
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
	}
	ResponseSuccess(c, reviews)
}

func (r *ReviewHandler) ListReviewByBookId(c *gin.Context) {
	var request = c.Param("book id")
	reviews, err := r.reviewSvc.ListReviewByBookId(c.Request.Context(), request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, reviews)
}

func (r *ReviewHandler) UpdateReviewByReviewId(c *gin.Context) {
	var request model.Review
	var id = c.Param("review id")
	err := c.ShouldBindJSON(&request)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	review, err := r.reviewSvc.UpdateReviewByReviewId(c.Request.Context(), id, request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, review)
}

func (r *ReviewHandler) DeleteReviewByReviewId(c *gin.Context) {
	var request model.Review
	var id = c.Param("review id")
	err := c.ShouldBindJSON(&request)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	review, err := r.reviewSvc.DeleteReviewByReviewId(c.Request.Context(), id, request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, review)
}

// reply
func (r *ReviewHandler) CreateReply(c *gin.Context) {
	var request model.ReplyReview
	err := c.ShouldBindJSON(&request)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	reply, err := r.reviewSvc.CreateReply(c.Request.Context(), request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, reply)
}

func (r *ReviewHandler) ListAllReply(c *gin.Context) {
	reply, err := r.reviewSvc.ListAllReply(c.Request.Context())
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
	}
	ResponseSuccess(c, reply)
}

func (r *ReviewHandler) ListReplyByReviewId(c *gin.Context) {
	var request = c.Param("review id")
	reply, err := r.reviewSvc.ListReplyByReviewId(c.Request.Context(), request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, reply)
}

func (r *ReviewHandler) UpdateReplyByReplyId(c *gin.Context) {
	var request model.ReplyReview
	var id = c.Param("reply id")
	err := c.ShouldBindJSON(&request)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	reply, err := r.reviewSvc.UpdateReplyByReplyId(c.Request.Context(), id, request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, reply)
}

func (r *ReviewHandler) DeleteReplyByReplyId(c *gin.Context) {
	var request model.ReplyReview
	var id = c.Param("reply id")
	err := c.ShouldBindJSON(&request)
	if err != nil {
		ResponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	reply, err := r.reviewSvc.DeleteReplyByReplyId(c.Request.Context(), id, request)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseSuccess(c, reply)
}
