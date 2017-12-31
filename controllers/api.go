package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"gitlab.com/ExpenseApp/models"
)

func Login(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindWith(&user, binding.Form); err == nil {
		context.JSON(http.StatusOK, gin.H{"status": "you are logged in" + user.FirstName + "-" + user.LastName})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func Register(context *gin.Context) {

}
