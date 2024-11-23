package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			ctx.Abort()
			return
		}

		tknStr := cookie

		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("userID", claims.UserID)
		ctx.Next()
	})
}
