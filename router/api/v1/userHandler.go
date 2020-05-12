package v1

import (
	"CMEMdc_be/models"
	"CMEMdc_be/utils"
	"CMEMdc_be/utils/e"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//NewPostHandler ...
//func NewUserHandler(db *gorm.DB) *User {
//	return &User{
//		service: Database.NewService(db),
//	}
//}
//
//// Post ...
//type User struct {
//	service Database.Service
//}

// 注册
//func (p *User) Create(ctx *gin.Context) {
//	//绑定数据
//	newUser := models.User{}
//	err := ctx.BindJSON(&newUser)
//	if err != nil {
//		log.Println(err)
//	}
//	//newUser.Email = ctx.Param("email")
//	//newUser.Name = ctx.Param("name")
//	//newUser.Password = ctx.Param("password")
//
//	newUser.Permission = "viewer"
//
//	//查询用户名是否注册
//	if hasName, _ := p.service.User.FindUserByName(newUser.Name, ""); hasName {
//		ctx.JSON(http.StatusOK, gin.H{
//			"code": 20000,
//			"msg":  "username already exists",
//			"data": gin.H{
//				"name": newUser.Name,
//			},
//		})
//		return
//	}
//
//	err = p.service.User.Create(&newUser)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	fmt.Println("newId: ------ ", newUser)
//
//	ctx.JSON(http.StatusOK, gin.H{
//		"code": 20000,
//		"msg":  "Successfully Created",
//		"data": gin.H{
//			"name": newUser.Name,
//			"ID":   newUser.ID,
//		},
//	})
//}
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
func Login(c *gin.Context) {
	User := models.UserAuth{}
	code := e.INVALID_PARAMS
	err := c.ShouldBindJSON(&User)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(User.Name, User.Password)
		//查询用户名是否注册
		if hasName, err := models.FindUserByName(User.Name, User.Password); hasName {
			//用户名正确
			fmt.Println(err)
			if err != nil {
				//密码错误
				log.Println(err)
				code = e.ERROR_WRONG_PASSWORD
			} else {
				//验证通过返回 token
				code = e.SUCCESS
				return
			}
		} else {
			//用户名不存在
			code = e.ERROR_NOT_EXIST_USER
			fmt.Println(err.Error())
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		//"msg":  err.Error(),
		"data": User,
	})

}

//func (u *User) Logout(c *gin.Context) {
//	code := e.SUCCESS
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": gin.H{},
//	})
//}
//
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
	claims,err := utils.ParseToken(c.Request.Header["Token"][0])

	//fmt.Println("claims", claims)
	user, err := models.FindUserInfo(claims.Username)
	if err != nil {
		code = e.INVALID_PARAMS
		fmt.Println("err==========",err.Error())
	} else {
		code = e.SUCCESS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": gin.H{
				"name":       user.Name,
				"email":      user.Email,
				"permission": user.Permission,
			},
		})
		return
	}

	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": code,
	//	"msg":  e.GetMsg(code),
	//	"data": gin.H{},
	//})

}

// GetDataByTime 一个需要token认证的测试接口
//func (u *User) GetDataByTime(c *gin.Context) {
//	claims := c.MustGet("claims").(*myjwt.CustomClaims)
//	if claims != nil {
//		c.JSON(http.StatusOK, gin.H{
//			"code": e.ACCESS_TOKEN,
//			"msg":  e.GetMsg(e.ACCESS_TOKEN),
//			"data": claims,
//		})
//	}
//}
