package Todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "root:iamgroot@tcp(127.0.0.1:3310)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func CreateTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := Todo{Title: c.PostForm("title"), Completed: completed}
	db := Database()
	db.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

func GetTodo(c *gin.Context) {
	var todos []Todo

	db := Database()
	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todos})
}

func GetTodoById(c *gin.Context) {
	var todo Todo
	todoId := c.Param("id")

	db := Database()
	db.First(&todo, todoId)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": todo})
}

func UpdateTodo(c *gin.Context) {
	var todo Todo
	todoId := c.Param("id")
	db := Database()
	db.First(&todo, todoId)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

func DeleteTodo(c *gin.Context) {
	var todo Todo
	todoId := c.Param("id")
	db := Database()
	db.First(&todo, todoId)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
		return
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}
