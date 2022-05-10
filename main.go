package main

import (
	"fmt"
	"log"
	"mediaislam/auth"
	"mediaislam/handler"
	"mediaislam/helper"
	"mediaislam/materi"
	"mediaislam/user"
	"mediaislam/ustadz"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/mediaislam?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	materiRepository := materi.NewRepository(db)
	ustadzRepository := ustadz.NewRepository(db) 

	userService := user.NewService(userRepository)
	materiService := materi.NewService(materiRepository)
	ustadzService := ustadz.NewService(ustadzRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	materiHandler := handler.NewMateriHandler(materiService)
	ustadzHandler := handler.NewUstadzHandler(ustadzService)
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.POST("/ustadz", ustadzHandler.RegisterUstadz)
	api.GET("/ustadz", ustadzHandler.GetUstadzList)
	api.GET("/ustadz/:id",  ustadzHandler.GetUstadz)


	api.GET("/materiall", materiHandler.GetMateriList)
	api.GET("/materiall/:id", materiHandler.GetMateri)
	router.Run(":8080")
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		fmt.Println(claim)
		if !ok || !token.Valid {
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		user_ID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(user_ID)
		if err != nil {
			response := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}
}
