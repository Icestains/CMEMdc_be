package v1

import (
	"CMEMdc_be/utils/app"
	"CMEMdc_be/utils/logging"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"CMEMdc_be/models"
	"CMEMdc_be/utils"
	"CMEMdc_be/utils/e"
)

// 注册
// @Summary 用户注册账号
// @Produce  json
// @Success 200 {object} app.Response
// @Router /register [post]
func Create(ctx *gin.Context) {
	appG := app.Gin{ctx}

	//绑定数据
	newUser := models.User{}
	err := ctx.BindJSON(&newUser)
	if err != nil {
		logging.Info(err)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, err.Error())
	}

	newUser.Permission = "viewer"
	valid := validation.Validation{}
	valid.Required(newUser.Name, "name").Message("用户名不能为空")
	valid.Required(newUser.Password, "password").Message("密码不能为空")
	valid.Required(newUser.Email, "email").Message("邮箱不能为空")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	hasName, _ := models.FindUserByName(newUser.Name, "")
	if hasName {
		logging.Info(err)
		appG.Response(http.StatusOK, e.ERROR_EXIST_USER, nil)
		return
	}

	err = models.Create(&newUser)
	if err != nil {
		logging.Info(err)
		appG.Response(http.StatusOK, e.ERROR, err.Error())
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary 退出登录
// @Produce  json
// @Success 200 {object} app.Response
// @Router /v1/user/logout [post]
func Logout(c *gin.Context) {

	appG := app.Gin{c}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

// @Summary 查询用户信息
// @Produce  json
// @Success 200 {object} app.Response
// @Router /v1/user/info [get]
func GetUserInfo(c *gin.Context) {
	appG := app.Gin{c}

	data := make(map[string]interface{})
	user := models.User{}

	claims, err := utils.ParseToken(c.Request.Header.Get("Authorization"))

	if err != nil {
		logging.Info(err.Error())
		appG.Response(http.StatusOK, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	user, err = models.FindUserInfo(claims.Username)
	if err != nil {
		data["error"] = err.Error()
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		logging.Info(err.Error())
		return
	}

	data["name"] = user.Name
	data["email"] = user.Email
	data["permission"] = user.Permission

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
