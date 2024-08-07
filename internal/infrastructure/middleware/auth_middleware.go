package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// fungsi untuk otentikasi token JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		// untuk memeriksa token yang ada pada header Authorization
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token is missing"})
			ctx.Abort()
			return
		}
		// untuk memeriksa format token (Bearer token)
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}
		// untuk memparsing token JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		// untuk memeriksa kesalahan dalam parsing atau check token yang tidak valid
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}
		// menyimpan data user yang ada di token JWT ke dalam context
		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("username", claims["username"].(string))
		ctx.Next()
	}
}
