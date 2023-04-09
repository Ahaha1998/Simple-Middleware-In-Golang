package controllers

import (
	"challenge-12/database"
	"challenge-12/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	Product := []models.ProductModel{}

	err := db.Preload("User").Where("user_id = ?", userId).Find(&Product).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
			"message": "Currently there is no product for this user!",
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func GetProductById(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	productId, _ := strconv.Atoi(c.Param("productId"))
	Product := models.ProductModel{}

	err := db.Find(&Product, productId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
			"message": "There is no product with that id!",
		})
		return
	}

	if Product.UserId != userId {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "Can't access others product!",
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := c.ContentType()
	userId := uint(userData["id"].(float64))

	Product := models.ProductModel{}
	User := models.UserModel{}

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	db.First(&User, userId)

	Product.UserId = userId
	Product.User = &User

	err := db.Debug().Create(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Product)
}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := c.ContentType()

	Product := models.ProductModel{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	
	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserId = uint(userData["id"].(float64))
	Product.ID = uint(productId)

	err := db.Model(&Product).Updates(models.ProductModel{Title: Product.Title, Description: Product.Description}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Product)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	Product := models.ProductModel{}

	productId, _ := strconv.Atoi(c.Param("productId"))

	err := db.Where("id=?", productId).Delete(&Product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Product with id %d has been successfully deleted", productId),
	})
}