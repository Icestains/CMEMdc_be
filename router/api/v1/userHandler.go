package v1

import (
	"CMEMdc_be/utils/logging"
	"fmt"
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
	//绑定数据
	newUser := models.User{}
	err := ctx.BindJSON(&newUser)
	if err != nil {
		logging.Info(err)
	}
	//newUser.Email = ctx.Param("email")
	//newUser.Name = ctx.Param("name")
	//newUser.Password = ctx.Param("password")

	newUser.Permission = "viewer"
	valid := validation.Validation{}
	//valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(newUser.Name, "name").Message("用户名不能为空")
	valid.Required(newUser.Password, "password").Message("密码不能为空")
	valid.Required(newUser.Email, "email").Message("邮箱不能为空")
	//valid.Required(createdBy, "created_by").Message("创建人不能为空")
	//valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if hasName, _ := models.FindUserByName(newUser.Name, ""); !hasName {
			err = models.Create(&newUser)
			if err != nil {
				logging.Info(err)
			}
			code = e.SUCCESS
		} else {
			code = e.ERROR_EXIST_USER
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": gin.H{
			"name": newUser.Name,
			"ID":   newUser.ID,
			"err":  err,
		},
	})
}

//
////注销
//func (p *User) Delete(ctx *gin.Context) {
//	code := e.INVALID_PARAMS
//	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
//	if err != nil {
//		log.Println(err)
//		code = e.INVALID_PARAMS
//	} else {
//		fmt.Println("deleted uid:", uint(id))
//		err = p.service.User.Delete(uint(id))
//		if err != nil {
//			panic(err)
//			code = e.ERROR_WRONG_ID
//		} else {
//			code = e.SUCCESS
//		}
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": id,
//	})
//}
//
////更新信息
//func (p *User) Update(c *gin.Context) {
//	code := e.INVALID_PARAMS
//	newUser := models.User{}
//	err := c.BindJSON(&newUser)
//	if err != nil {
//		log.Println(err)
//
//	} else {
//		fmt.Println(
//			c.Params,
//			newUser,
//		)
//		p.service.User.Update(&newUser)
//		code = e.SUCCESS
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": gin.H{
//			"UpdatedTime": newUser.UpdatedAt,
//		}})
//}

//用户登录
//func Login(c *gin.Context) {
//	User := models.UserAuth{}
//	code := e.INVALID_PARAMS
//	err := c.ShouldBindJSON(&User)
//	if err != nil {
//		log.Println(err)
//	} else {
//		fmt.Println(User.Name, User.Password)
//		//查询用户名是否注册
//		if hasName, err := models.FindUserByName(User.Name, User.Password); hasName {
//			//用户名正确
//			fmt.Println(err)
//			if err != nil {
//				//密码错误
//				log.Println(err)
//				code = e.ERROR_WRONG_PASSWORD
//			} else {
//				//验证通过返回 token
//				code = e.SUCCESS
//				return
//			}
//		} else {
//			//用户名不存在
//			code = e.ERROR_NOT_EXIST_USER
//			fmt.Println(err.Error())
//		}
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		//"msg":  err.Error(),
//		"data": User,
//	})
//
//}

// @Summary 退出登录
// @Produce  json
// @Success 200 {object} app.Response
// @Router /v1/user/logout [post]
func Logout(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": gin.H{},
	})
}


//func (u *User) GetName(ctx *gin.Context) {
//	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
//	fmt.Println("claims", claims)
//	if claims != nil {
//		ctx.JSON(http.StatusOK, gin.H{
//			"code": e.SUCCESS,
//			"msg":  e.GetMsg(e.SUCCESS),
//			"data": gin.H{
//				"name": claims.Name,
//			},
//		})
//	}
//
//}
//

// @Summary 查询用户信息
// @Produce  json
// @Success 200 {object} app.Response
// @Router /v1/user/info [get]
func GetUserInfo(c *gin.Context) {
	code := e.INVALID_PARAMS
	data := make(map[string]interface{})
	user := models.User{}

	claims, err := utils.ParseToken(c.Request.Header["Token"][0])

	if err != nil {
		logging.Info(err.Error())
		code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
	} else {
		user, err = models.FindUserInfo(claims.Username)
		if err != nil {
			code = e.INVALID_PARAMS
			fmt.Println("err==========", err.Error())
			logging.Info(err.Error())
			data["error"] = err.Error()
		} else {
			code = e.SUCCESS
			data["name"] = user.Name
			data["email"] = user.Email
			data["permission"] = user.Permission
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
