package service

import (
	"bookstore/internal/model"
	"bookstore/internal/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Định dạng token không hợp lệ"})
			return
		}

		tokenString := tokenParts[1]

		var jwtRepo repository.JWTRepo

		claims, err := jwtRepo.VerifyJWT(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Token không hợp lệ: %v", err)})
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}

func ProtectedHandler(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"error": "Không thể lấy thông tin người dùng"})
		return
	}

	claims, ok := user.(*model.UserClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Định dạng thông tin người dùng không hợp lệ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chào mừng!", "user_id": claims.ID, "username": claims.Username})
}
