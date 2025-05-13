package handlers

import (
	"bookstore/internal/config"
	"bookstore/internal/repository"
	"bookstore/internal/service"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router    *gin.Engine
	server    *http.Server
	user      *UserHandler
	auth      *AuthHandler
	whishlist *WhishlistHandler
	book      *BookHandler
	review    *ReviewHandler
	order     *OrderHandler
	payment   *PaymentHandler
}

func NewServer(cfg config.Config, user *UserHandler, auth *AuthHandler, whishlist *WhishlistHandler, handler *BookHandler, review *ReviewHandler, order *OrderHandler, payment *PaymentHandler) *Server {
	router := gin.Default()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}
	srv := &Server{
		user:      user,
		auth:      auth,
		whishlist: whishlist,
		book:      handler,
		review:    review,
		router:    router,
		server:    server,
		order:     order,
		payment:   payment,
	}
	srv.Register(cfg)
	return srv
}

func (srv *Server) Start() error {
	return srv.server.ListenAndServe()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.server.Shutdown(ctx)
}

func (srv *Server) Register(cfg config.Config) {
	//public service
	srv.router.POST("/api/v1/user/register", srv.user.Register)
	srv.router.POST("/api/v1/auth/login", srv.auth.Login)
	srv.router.GET("/api/v1/book", srv.book.ListAllBooks)
	srv.router.GET("/api/v1/book/:bookname", srv.book.SearchBookByName)
	srv.router.GET("/api/v1/review/:book id", srv.review.ListReviewByBookId)
	srv.router.GET("/api/v1/reply/:review id", srv.review.ListReplyByReviewId)

	protected := srv.router.Group("/api/v1/service")
	protected.Use(service.AuthMiddleware(repository.NewJWTRepo(cfg.Key.JwtSecret)))
	{
		//book service
		protected.POST("/book", srv.book.CreateBook, service.AdminHandler)
		protected.GET("/book/:id", srv.book.SearchBooks, service.AdminHandler)
		protected.PUT("/book/:id", srv.book.UpdateBook, service.AdminHandler)
		protected.DELETE("/book/:id", srv.book.DeleteBook, service.AdminHandler)

		//review service
		protected.POST("/review", srv.review.CreateReview, service.UserHandler)
		protected.GET("/review", srv.review.ListAllReview, service.AdminHandler)
		protected.PUT("/review/:review id", srv.review.UpdateReviewByReviewId, service.UserHandler)
		protected.DELETE("/review/:review_id", srv.review.DeleteReviewByReviewId, service.UserHandler)

		//reply review service
		protected.POST("/reply", srv.review.CreateReply, service.UserHandler)
		protected.GET("/reply", srv.review.ListAllReply, service.AdminHandler)
		protected.PUT("/reply/:reply id", srv.review.UpdateReplyByReplyId, service.UserHandler)
		protected.DELETE("/reply/:reply id", srv.review.DeleteReplyByReplyId, service.UserHandler)

		//user service
		protected.GET("/user", srv.user.ListUsers, service.AdminHandler)
		protected.GET("/user/:id", srv.user.SearchUser, service.AdminHandler)
		protected.PUT("/user/:id", srv.user.UpdateUser, service.UserHandler)
		protected.DELETE("/user/:id", srv.user.DeleteUser, service.UserHandler)

		//whishlist service
		protected.POST("/whishlist", srv.whishlist.AddWhishItem, service.UserHandler)
		protected.GET("/whishlist/:user id", srv.whishlist.GetWhishlistByUserId, service.UserHandler)
		protected.DELETE("/whishlist/:item id", srv.whishlist.DeleteWhishItemByWishItemId, service.UserHandler)

		//order service
		protected.POST("/order", srv.order.CreateOrder, service.UserHandler)
		protected.GET("/order", srv.order.ListAllOrder, service.AdminHandler)
		protected.GET("/order/:id", srv.order.SearchOrder, service.UserHandler)
		protected.GET("/order/user/:user id", srv.order.ListAllOrderByUserId, service.UserHandler)
		protected.PUT("/order/:id", srv.order.UpdateOrder, service.UserHandler)
		protected.DELETE("/order/:id", srv.order.DeleteOrder, service.UserHandler)

		//payment service
		protected.POST("/payment", srv.payment.CreatePayment, service.UserHandler)
		protected.GET("/payment", srv.payment.ListAllPayments, service.AdminHandler)
		protected.GET("/payment/:order id", srv.payment.SearchPayment, service.UserHandler)
		protected.PUT("/payment/:order id", srv.payment.ConfirmPayment, service.UserHandler)
	}
}
