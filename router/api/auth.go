package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"CMEMdc_be/models"
	"CMEMdc_be/utils"
	"CMEMdc_be/utils/e"
)

type auth struct {
	Name string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary 用户登录
// @Produce  json
// @Success 200 {object} app.Response
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})

	User := auth{}
	err := c.ShouldBindJSON(&User)
	if err != nil {
		log.Println(err)
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

					code = e.SUCCESS
				}

			} else {
				code = e.ERROR_AUTH
			}
		} else {
			for _, err := range valid.Errors {
				log.Println(err.Key, err.Message)
			}
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
