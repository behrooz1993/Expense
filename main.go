package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"gitlab.com/ExpenseApp/controllers"
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

func initRoutes() {
	v1 := router.Group("/v1")
	{
		v1.POST("/user/register", controllers.Register)
		v1.POST("/user/login", controllers.Login)
	}
}
