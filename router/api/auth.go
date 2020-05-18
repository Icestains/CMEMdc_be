package api

import (
	"CMEMdc_be/service/user_service"
	"CMEMdc_be/utils/app"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"CMEMdc_be/utils"
	"CMEMdc_be/utils/e"
	"CMEMdc_be/utils/logging"
)

type auth struct {
	Name     string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary 用户登录获取 token
// @Produce  json
// @Success 200 {object} app.Response
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	data := make(map[string]interface{})

	User := auth{}
	err := c.ShouldBindJSON(&User)
	if err != nil {
		logging.Info(err.Error())
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	fmt.Println(User)
	username := User.Name
	password := User.Password

	a := auth{Name: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return

	}
	userService := user_service.Auth{Username: username, Password: password}
	isExist, err := userService.CheckAuth()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, err.Error())
		return
	}
	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_WRONG_PASSWORD, nil)
		return
	}
	token, err := utils.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, err.Error())
		return
	}

	data["token"] = token
	data["ver"] = "1.0"
	appG.Response(http.StatusOK, e.SUCCESS, data)

}
