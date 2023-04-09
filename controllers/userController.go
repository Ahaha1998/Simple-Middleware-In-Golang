package controllers

import (
	"challenge-12/database"
	"challenge-12/helpers"
	"challenge-12/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

func HasData(u *models.UserModel) bool {
    if u.ID != 0 || u.Username != "" || u.Email != "" {
        return true
    }
    return false
}

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := c.ContentType()

	TempUser, User := models.UserModel{}, models.UserModel{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	db.Where("email = ?", User.Email).First(&TempUser)
	checkUser := HasData(&TempUser)

	if checkUser {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "Email is already registered!",
		})
		return
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": User.ID,
		"email": User.Email,
		"username": User.Username,
		"isadmin": User.IsAdmin,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := c.ContentType()

	User := models.UserModel{}
	
	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password := User.Password

	err := db.Where("email = ?", User.Email).First(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "Email is not found!",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "Password incorrect!",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, User.IsAdmin)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}