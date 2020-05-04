package v1

import (
	"CMEMdc_be/Database"
	models "CMEMdc_be/Database/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

//NewPostHandler ...
func NewUserinfoHandler(db *gorm.DB) *UserInfo {
	return &UserInfo{
		service: Database.NewService(db),
	}
}

// Post ...
type UserInfo struct {
	service Database.Service
}

type Res struct {
	code int
	res  string
	data interface{}
}

// Fetch all post data
func (p *UserInfo) Fetch(ctx *gin.Context) {
	payload, _ := p.service.Userinfo.FindAll()
	//res := Res{
	//	code: 20000,
	//	res:  "success",
	//	data: payload,
	//}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"res":  "success",
		"data": payload,
	})
	//respoense(ctx, http.StatusOK, payload)
}

// Create a new post
func (p *UserInfo) Create(ctx *gin.Context) {
	newUser := models.Userinfo{}
	err := ctx.BindJSON(&newUser)
	if err != nil {
		panic(err)
	}
	newUser.Uid = time.Now().Unix()
	err = p.service.Userinfo.Create(&newUser)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("newId: ------ ", newUser)

	//respondwithJSON(ctx.Writer, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "Successfully Created",
		"data": gin.H{
			"uid": newUser.Uid,
		},
	})
}

func (p *UserInfo) Delete(ctx *gin.Context) {
	uid, err := strconv.ParseInt(ctx.Param("uid"), 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("deleted uid:", uid)
	err = p.service.Userinfo.Delete(uid)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "delete success",
		"data": uid,
	})
}

func (p *UserInfo) Update(c *gin.Context) {
	newUser := models.Userinfo{}
	err := c.BindJSON(&newUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(
		c.Params,
		newUser,
	)
	p.service.Userinfo.Update(&newUser)
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"msg":  "update success",
		"data": gin.H{
			"UpdatedTime":time.Now(),
		},
	})
}

