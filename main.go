package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gitlab.com/ExpenseApp/controllers"
	"gitlab.com/ExpenseApp/helpers"
)

var router *gin.Engine

func main() {
	file, _ := os.Create("expensify.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	router = gin.Default()

	initRoutes()

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func authMiddleware(context *gin.Context) {
	authorizationHeader := context.GetHeader("Authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			claims, err := helpers.ValidateJwtToken(bearerToken[1])
			if err != nil {
				log.Print(err)
			} else {
				log.Print(claims)
				expiration := time.Now().Add(365 * 24 * time.Hour)
				cookie := http.Cookie{Name: "userId", Value: claims["userId"].(string), Expires: expiration}
				http.SetCookie(context.Writer, &cookie)
				// context.SetCookie("userId", claims["userId"].(string), 1000, "/", "localhost", true, false)
				// log.Print() context.Request.Cookies
				context.Next()
			}
		}
	} else {
		context.JSON(http.StatusOK, gin.H{"error": "An authorization header is required"})
	}
}

func initRoutes() {
	v1 := router.Group("/v1")
	{
		v1.POST("/user/register", controllers.Register)
		v1.POST("/user/login", controllers.Login)

		v1.POST("/user/testJwt", authMiddleware, controllers.TestJwt)
	}
}
