package v1

import (
	"CMEMdc_be/Database"
	models "CMEMdc_be/Database/model"
	myjwt "CMEMdc_be/router/Middlewares/jwt"
	"CMEMdc_be/utils/e"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

//NewPostHandler ...
func NewUserHandler(db *gorm.DB) *User {
	return &User{
		service: Database.NewService(db),
	}
}

// Post ...
type User struct {
	service Database.Service
}

// 注册
func (p *User) Create(ctx *gin.Context) {
	//绑定数据
	newUser := models.User{}
	err := ctx.BindJSON(&newUser)
	if err != nil {
		panic(err)
	}
	//newUser.Email = ctx.Param("email")
	//newUser.Name = ctx.Param("name")
	//newUser.Password = ctx.Param("password")

	newUser.Promission = "viewer"

	//查询用户名是否注册
	if hasName, _ := p.service.User.FindUserByName(newUser.Name, ""); hasName {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 20000,
			"msg":  "username already exists",
			"data": gin.H{
				"name": newUser.Name,
			},
		})
		return
	}

	err = p.service.User.Create(&newUser)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("newId: ------ ", newUser)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "Successfully Created",
		"data": gin.H{
			"name": newUser.Name,
			"ID":   newUser.ID,
		},
	})
}

//注销
func (p *User) Delete(ctx *gin.Context) {
	code := e.INVALID_PARAMS
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		log.Println(err)
		code = e.INVALID_PARAMS
	} else {
		fmt.Println("deleted uid:", uint(id))
		err = p.service.User.Delete(uint(id))
		if err != nil {
			panic(err)
			code = e.ERROR_WRONG_ID
		} else {
			code = e.SUCCESS
		}
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": id,
	})
}

//更新信息
func (p *User) Update(c *gin.Context) {
	code := e.INVALID_PARAMS
	newUser := models.User{}
	err := c.BindJSON(&newUser)
	if err != nil {
		log.Println(err)

	} else {
		fmt.Println(
			c.Params,
			newUser,
		)
		p.service.User.Update(&newUser)
		code = e.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": gin.H{
			"UpdatedTime": newUser.UpdatedAt,
		}})
}

//用户登录
func (u *User) Login(c *gin.Context) {
	User := models.UserAuth{}
	code := e.INVALID_PARAMS
	err := c.ShouldBindJSON(&User)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(User.Name, User.Password)
		//查询用户名是否注册
		if hasName, err := u.service.User.FindUserByName(User.Name, User.Password); hasName {
			//用户名正确
			fmt.Println(err)
			if err != nil {
				//密码错误
				log.Println(err)
				code = e.ERROR_WRONG_PASSWORD
			} else {
				//验证通过返回 token
				code = e.SUCCESS
				generateToken(c, User)
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

func (u *User) Logout(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": gin.H{},
	})
}

func (u *User) GetName(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	fmt.Println("claims", claims)
	if claims != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  e.GetMsg(e.SUCCESS),
			"data": gin.H{
				"name": claims.Name,
			},
		})
	}

}

func (u *User) GetUserInfo(ctx *gin.Context) {
	code := e.INVALID_PARAMS
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	fmt.Println("claims", claims)
	if claims != nil {
		user, err := u.service.User.FindUserEmail(claims.Name)
		if err != nil {
			code = e.INVALID_PARAMS
		} else {
			code = e.SUCCESS
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": gin.H{
					"name":       user.Name,
					"email":      user.Email,
					"promission": user.Promission,
				},
			})
			return
		}

	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": gin.H{},
	})

}

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	models.UserAuth
}

// 生成令牌
func generateToken(c *gin.Context, user models.UserAuth) {
	code := e.SUCCESS
	j := &myjwt.JWT{
		[]byte(""),
	}
	claims := myjwt.CustomClaims{
		user.Name,
		user.Password,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),    // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600*24), // 过期时间 一天
			Issuer:    "icestains",
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		code = e.ERROR_AUTH
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": gin.H{},
		})
		return
	}
	code = e.SUCCESS
	log.Println(token)
	data := LoginResult{
		UserAuth: user,
		Token:    token,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": gin.H{
			"token": data.Token,
		},
	})
}

// GetDataByTime 一个需要token认证的测试接口
func (u *User) GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ACCESS_TOKEN,
			"msg":  e.GetMsg(e.ACCESS_TOKEN),
			"data": claims,
		})
	}
}
