package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"CMEMdc_be/models"
)

// Fetch all post data
// @Summary 查找所有用户信息
// @Produce  json
// @Success 200 {object} app.Response
// @Router /v1/get [get]
//func Fetch(ctx *gin.Context) {
//	payload, _ := p.service.Userinfo.FindAll()
//	//res := Res{
//	//	code: 20000,
//	//	res:  "success",
//	//	data: payload,
//	//}
//	ctx.JSON(http.StatusOK, gin.H{
//		"code": 20000,
//		"res":  "success",
//		"data": payload,
//	})
//	//respoense(ctx, http.StatusOK, payload)
//}


// @Summary 用户注册账号
// @Produce  json
// @Success 200 {object} app.Response
// @Router /register [post]
func Create(ctx *gin.Context) {
	newUser := models.User{}
	err := ctx.BindJSON(&newUser)
	if err != nil {
		panic(err)
	}
	//newUser.Uid = time.Now().Unix()
	err = models.Create(&newUser)
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
			"uid": newUser.ID,
		},
	})
}

//func (p *UserInfo) Delete(ctx *gin.Context) {
//	uid, err := strconv.ParseInt(ctx.Param("uid"), 10, 64)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("deleted uid:", uid)
//	err = p.service.Userinfo.Delete(uid)
//	if err != nil {
//		panic(err)
//	}
//	ctx.JSON(http.StatusOK, gin.H{
//		"code": 20000,
//		"msg":  "delete success",
//		"data": uid,
//	})
//}

//func (p *UserInfo) Update(c *gin.Context) {
//	newUser := models.Userinfo{}
//	err := c.BindJSON(&newUser)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(
//		c.Params,
//		newUser,
//	)
//	p.service.Userinfo.Update(&newUser)
//	c.JSON(http.StatusOK, gin.H{
//		"code": 20000,
//		"msg":  "update success",
//		"data": gin.H{
//			"UpdatedTime": time.Now(),
//		},
//	})
//}
