package middlewares

import (
	"challenge-12/helpers"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthenticated",
				"message": err.Error(),
			})
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}


func CheckUserLevel() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)
		isAdmin := userData["isadmin"].(bool)
		
		method := c.Request.Method

		if !isAdmin {
			if method == "PUT" || method == "DELETE" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthenticated",
					"message": "This method is forbidden or only admin can afford it!",
				})
				return
			}
		}

		c.Next()
	}
}