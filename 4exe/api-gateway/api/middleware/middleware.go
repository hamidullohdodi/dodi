package middleware

import (
	"api-geteway/api/token"
	"errors"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type casbinPermission struct {
	enforcer *casbin.Enforcer
}

func (c *casbinPermission) GetRole(ctx *gin.Context) (string, int) {
	Token := ctx.GetHeader("Authorization")
	if Token == "" {
		return "unauthorized", http.StatusUnauthorized
	}

	claims, err := token.ExtractClaims(Token)
	if err != nil || claims == nil {
		fmt.Println(err)
		return "unauthorized", http.StatusUnauthorized
	}

	role, ok := claims["role"].(string)
	if !ok {
		fmt.Println(err)
		return "unauthorized", http.StatusUnauthorized
	}
	return role, 0
}

func (c *casbinPermission) CheckPermission(ctx *gin.Context) (bool, error) {
	subject, status := c.GetRole(ctx)
	if status != 0 {
		return false, errors.New("error while getting a role: " + subject)
	}
	action := ctx.Request.Method
	object := ctx.Request.URL.Path

	allow, err := c.enforcer.Enforce(subject, object, action)
	if err != nil {
		return false, err
	}
	return allow, nil
}

func PermissionMiddleware(enf *casbin.Enforcer) gin.HandlerFunc {
	casbHandler := &casbinPermission{
		enforcer: enf,
	}

	return func(ctx *gin.Context) {
		res, err := casbHandler.CheckPermission(ctx)

		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}
		if !res {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "You don't have permission"})
			return
		}
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": "authorization header is required"})
			return
		}

		valid, err := token.ValidateToken(auth)
		if err != nil || !valid {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("invalid token: %s", err)})
			return
		}

		claims, err := token.ExtractClaims(auth)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("invalid token claims: %s", err)})
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
