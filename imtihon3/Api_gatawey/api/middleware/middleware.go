package middleware

import (
	"api_getway/api/token"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		fmt.Println(tokenParts)
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}
		tokenStr := tokenParts[1]

		valid, err := token.ValidateToken(tokenStr)
		if err != nil || !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		claims, err := token.ExtractClaims(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}
		ctx.Set("claims", claims)

		ctx.Next()
	}
}
