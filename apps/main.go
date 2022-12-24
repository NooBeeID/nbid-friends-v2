package main

import (
	"nbfriends/apps/config"
	"nbfriends/apps/controller"
	"nbfriends/apps/pkg/token"
	"nbfriends/apps/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	router := gin.New()

	// router.Use(Log())
	router.Use(gin.Logger())

	authController := controller.AuthController{
		Db: db,
	}

	v1 := router.Group("v1")

	router.GET("/ping", Ping)

	auth := v1.Group("auth")
	{
		auth.POST("register", authController.Register)
		auth.POST("login", authController.Login)
		auth.GET("profile", CheckAuth(), authController.Profile)
	}
	v2 := router.Group("v2")
	authV2 := v2.Group("auth")
	{
		authV2.POST("login", authController.Login)

	}

	router.Run(":4444")
}

func CORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Request-Methods", "GET, OPTIONS, POST, PUT, DELETE")
		ctx.Header("Access-Control-Request-Headers", "Authorization, Content-Type")
		ctx.Next()
	}
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		bearerToken := strings.Split(header, "Bearer ")
		if len(bearerToken) != 2 {
			resp := response.ResponseAPI{
				StatusCode: http.StatusUnauthorized,
				Message:    "UNAUTHORIZED",
			}
			ctx.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}

		payload, err := token.ValidateToken(bearerToken[1])
		if err != nil {
			resp := response.ResponseAPI{
				StatusCode: http.StatusUnauthorized,
				Message:    "INVALID TOKEN",
				Payload:    err.Error(),
			}
			ctx.AbortWithStatusJSON(resp.StatusCode, resp)
			return
		}

		ctx.Set("authId", payload.AuthId)

		ctx.Next()
	}
}

// func Log() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		log.Println("method", ctx.Request.Method)
// 		ctx.Next()
// 	}
// }
