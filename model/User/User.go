package User

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	First_Name string
	Last_Name  string
	Address    string
}

func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "root:iamgroot@tcp(127.0.0.1:3310)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func CreateUser(c *gin.Context) {
	user := User{First_Name: c.PostForm("first_name"), Last_Name: c.PostForm("last_name"), Address: c.PostForm("address")}
	db := Database()
	db.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "user created successfully!", "resourceId": user.ID})
}

func GetUser(c *gin.Context) {
	var users []User

	db := Database()
	db.Find(&users)

	if len(users) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": users})
}

func GetUserById(c *gin.Context) {
	var user User
	userId := c.Param("id")

	db := Database()
	db.First(&user, userId)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}

func UpdateUser(c *gin.Context) {
	var user User
	userId := c.Param("id")
	db := Database()
	db.First(&user, userId)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	db.Model(&user).Update("first_name", c.PostForm("first_name"))
	db.Model(&user).Update("last_name", c.PostForm("last_name"))
	db.Model(&user).Update("address", c.PostForm("address"))
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User updated successfully!"})
}

func DeleteUser(c *gin.Context) {
	var user User
	userId := c.Param("id")
	db := Database()
	db.First(&user, userId)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
		return
	}

	db.Unscoped().Where("id = ?", userId).Delete(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "user deleted successfully!"})
}
