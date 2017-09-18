package main

import (
	"golang-test-api/model/Todo"
	"golang-test-api/model/User"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database() *gorm.DB {
	//open a db connection
	db, err := gorm.Open("mysql", "root:iamgroot@tcp(127.0.0.1:3310)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func main() {

	//Migrate the schema
	// db := Database()
	// db.AutoMigrate(&User{})

	router := gin.Default()

	v1 := router.Group("/api/v1/")
	{

		//TODO
		v1.POST("todos/", Todo.CreateTodo)
		v1.GET("todos/", Todo.GetTodo)
		v1.GET("todos/:id", Todo.GetTodoById)
		v1.POST("todos/:id", Todo.UpdateTodo)
		v1.DELETE("todos/:id", Todo.DeleteTodo)

		//User
		v1.POST("users/", User.CreateUser)
		v1.GET("users/", User.GetUser)
		v1.GET("users/:id", User.GetUserById)
		v1.POST("users/:id", User.UpdateUser)
		v1.DELETE("users/:id", User.DeleteUser)

	}
	router.Run(":1200")

}
