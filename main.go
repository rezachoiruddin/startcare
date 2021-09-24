package main

import (
	"fmt"
	"log"
	"startcare/auth"
	"startcare/handler"
	"startcare/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func main() {
// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
// 	// 	dsn := "root:1n73GR!ya@tcp(127.0.0.1:3306)/gotrial?charset=utf8mb4&parseTime=True&loc=Local"
// 	// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	// 	if err != nil {
// 	// 		log.Fatal(err.Error())
// 	// 	}

// 	// 	fmt.Println("Database connected")

// 	// 	var users []user.User
// 	// 	db.Find(&users)

// 	// 	for _, user := range users {
// 	// 		fmt.Println(user.Name)
// 	// 		fmt.Println(user.Email)
// 	// 		fmt.Println("================")
// 	// 	}
// 	router := gin.Default()
// 	router.GET("/handler", handler)
// 	router.Run()
// }

func main() {
	dsn := "root:1n73GR!ya@tcp(127.0.0.1:3306)/gotrial?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjo0fQ.5uLQPYy4C1VKJ9cr05xswjgH0gwk8uTSiGltkAdHhVM")
	if err != nil {
		fmt.Println("ERROR")
	}

	if token.Valid {
		fmt.Println("VALID")
	} else {
		fmt.Println("INVALID")
	}

	// fmt.Println(authService.GenerateToken(1001))

	// userService.SaveAvatar(1, "images/1-profile.png")

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

	// input
	// handler mapping input to struct
	// service mapping struct User
	// repository save struct User to db
	// db
}
