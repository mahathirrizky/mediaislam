package main

import (
	"fmt"
	"log"
	"mediaislam/auth"
	"mediaislam/handler"
	"mediaislam/helper"
	"mediaislam/materi"
	"mediaislam/submateri"
	"mediaislam/subscribe"
	"mediaislam/user"
	"mediaislam/ustadz"
	"mediaislam/video"
	"mediaislam/videomateri"
	"mediaislam/watched"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
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
	submateriRepository := submateri.NewRepository(db)
	videomateriRepository := videomateri.NewRepository(db)
	videoRepository := video.NewRepository(db)
	watchedRepository := watched.NewRepository(db)
	ustadzRepository := ustadz.NewRepository(db)
	subscribeRepository := subscribe.NewRepository(db)

	userService := user.NewService(userRepository)
	materiService := materi.NewService(materiRepository)
	submateriService := submateri.NewService(submateriRepository)
	videomateriService := videomateri.NewService(videomateriRepository)
	videoService := video.NewService(videoRepository)
	watchedService := watched.NewService(watchedRepository)
	ustadzService := ustadz.NewService(ustadzRepository)
	subscribeService := subscribe.NewService(subscribeRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	materiHandler := handler.NewMateriHandler(materiService)
	submateriHandler := handler.NewSubmateriHandler(submateriService)
	videomateriHandler := handler.NewVideomateriHandler(videomateriService)
	videoHandler := handler.NewVideoHandler(videoService)
	watchedHandler := handler.NewWatchedHandler(watchedService)
	ustadzHandler := handler.NewUstadzHandler(ustadzService)
	subscribeHandler := handler.NewSubscribeHandler(subscribeService)
	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	api.POST("/ustadz", authMiddleware(authService, userService), ustadzHandler.RegisterUstadz)
	api.PUT("/ustadz/:id", authMiddleware(authService, userService), ustadzHandler.UpdateUstadz)
	api.GET("/ustadz", ustadzHandler.GetUstadzList)
	api.GET("/ustadz/:id", authMiddleware(authService, userService), ustadzHandler.GetUstadz)

	api.POST("/subscribe", authMiddleware(authService, userService), subscribeHandler.CreateSubscribe)
	api.GET("/subscribe", authMiddleware(authService, userService), subscribeHandler.GetSubscribe)

	api.POST("/materi", authMiddleware(authService, userService), materiHandler.CreateMateri)
	api.POST("/materiimage", authMiddleware(authService, userService), materiHandler.UploadImage)
	api.PUT("/materi/:id", authMiddleware(authService, userService), materiHandler.UpdateMateri)
	api.GET("/materi", materiHandler.GetMateriList)
	api.GET("/materi/:id", materiHandler.GetMateri)
	api.GET("/materiall/:id", authMiddleware(authService, userService), materiHandler.GetMateriSubandVideo)
	
	api.POST("/submateri", authMiddleware(authService, userService), submateriHandler.CreateSubmateri)
	api.PUT("/submateri/:id", authMiddleware(authService, userService), submateriHandler.UpdateSubmateri)

	api.POST("/videomateri", authMiddleware(authService, userService), videomateriHandler.CreateVideomateri)
	api.PUT("/videomateri/:id", authMiddleware(authService, userService), videomateriHandler.UpdateVideomateri)

	api.POST("/videotematik", authMiddleware(authService, userService), videoHandler.CreateVideoTematik)
	api.POST("/videotematikimage", authMiddleware(authService, userService), videoHandler.UploadImage)
	api.PUT("/videotematik/:id", authMiddleware(authService, userService), videoHandler.GetVideo)
	api.GET("/videotematik", videoHandler.GetTematikList)
	api.GET("/videotematik/:id", videoHandler.GetVideo)

	api.POST("/videoshort", authMiddleware(authService, userService), videoHandler.CreateVideoShort)
	api.POST("/videoshortimage", authMiddleware(authService, userService), videoHandler.UploadImage)
	api.PUT("/videoshort/:id", authMiddleware(authService, userService), videoHandler.GetVideo)
	api.GET("/videoshort", videoHandler.GetShortList)
	api.GET("/videoshort/:id", videoHandler.GetVideo)

	api.POST("/watched", authMiddleware(authService, userService), watchedHandler.CreateWatched)
	api.GET("/watched", authMiddleware(authService, userService), watchedHandler.GetWatched)

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
