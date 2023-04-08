package middlewares

import (
	"challenge-12/database"
	"challenge-12/helpers"
	"challenge-12/models"
	"net/http"
	"strconv"

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

func BookAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		bookId, err := strconv.Atoi(c.Param("bookId"))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Bad Request",
				"message": "Invalid parameter!",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userId := uint(userData["id"].(float64))
		Book := models.ProductModel{}

		err = db.Select("user_id").First(&Book, uint(bookId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Data Not Found",
				"message": "Data doesn't exist!",
			})
			return
		}

		if Book.UserId != userId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
				"message": "Access is denied!",
			})
			return
		}

		c.Next() 
	}
}