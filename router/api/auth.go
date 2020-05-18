package api

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"CMEMdc_be/models"
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
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})

	User := auth{}
	err := c.ShouldBindJSON(&User)
	if err != nil {
		logging.Info(err.Error())
	} else {
		fmt.Println(User)
		username := User.Name
		password := User.Password

		valid := validation.Validation{}
		a := auth{Name: username, Password: password}
		ok, _ := valid.Valid(&a)

		if ok {
			isExist := models.CheckAuth(username, password)
			if isExist {
				token, err := utils.GenerateToken(username, password)
				if err != nil {
					code = e.ERROR_AUTH_TOKEN
				} else {
					data["token"] = token
					data["ver"] = "1.0"
					code = e.SUCCESS
				}
			} else {
				code = e.ERROR_WRONG_PASSWORD
			}
		} else {
			for _, err := range valid.Errors {
				logging.Info(err.Key, err.Message)
			}
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
