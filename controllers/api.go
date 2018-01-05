package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"gitlab.com/ExpenseApp/helpers"
	"gitlab.com/ExpenseApp/models"
)

const (
	PARAM_CELLPHONE = "cellphone"
	PARAM_PASSWORD  = "password"
)

func Register(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindWith(&user, binding.Form); err == nil {
		password := helpers.HashPassword(context.PostForm(PARAM_PASSWORD))
		user.Password = password
		err := models.Mgr.AddUser(&user)
		if err == nil {
			generateJsonResponse(context, true, user, "")
		} else {
			log.Print(err.Error())
			generateJsonResponse(context, false, nil, "Server error")
		}
	} else {
		generateJsonResponse(context, false, nil, "Field error")
	}
}

func Login(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindWith(&user, binding.Form); err == nil {
		err := models.Mgr.SelectByCellphone(&user)
		if err != nil {
			generateJsonResponse(context, false, nil, "Server error")
			return
		}
		if user.ID == 0 {
			generateJsonResponse(context, false, nil, "Cellphone not found")
			return
		}
		if helpers.ComparePasswords(user.Password, []byte(context.PostForm(PARAM_PASSWORD))) {
			user.JwtToken, _ = helpers.GenerateJwtToken(user.Cellphone, false)
			generateJsonResponse(context, true, user, "")
		} else {
			generateJsonResponse(context, false, nil, "Password is not correct")
		}
	} else {
		generateJsonResponse(context, false, nil, "Field error")
	}
}

func generateJsonResponse(context *gin.Context, status bool, data interface{}, err string) {
	var response = models.Response{}
	response.Status = status
	response.Data = data
	response.Error = err
	context.JSON(http.StatusOK, response)
}
