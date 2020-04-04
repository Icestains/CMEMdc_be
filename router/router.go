package router

import (
	"CMEMdc_be/router/Middlewares"
	myjwt "CMEMdc_be/router/Middlewares/jwt"
	v1 "CMEMdc_be/router/api/v1"
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
)

func Init(host, port string, dbConn *gorm.DB) {
	//CONN = db

	r := gin.Default()
	r.Use(Middlewares.Cors())
	fmt.Println("--------------------------------")
	//r.Use(RequestInfos())
	handlerUserInfo := v1.NewUserinfoHandler(dbConn)
	handlerUser := v1.NewUserHandler(dbConn)
	apiv1 := r.Group("/v1")
	{
		apiv1.GET("/get", handlerUserInfo.Fetch)
		apiv1.POST("/post", handlerUserInfo.Create)
		//apiv1.DELETE("/delete/:uid", handlerUserInfo.Delete)
		apiv1.POST("/update", handlerUserInfo.Update)
		apiv1.GET("/mqtt", handlerUserInfo.Mqtt)
	}

	user := apiv1.Group("")
	{
		//user.POST("/register", handlerUser.Create)
		user.POST("/login", handlerUser.Login)
		user.POST("/sign-up", handlerUser.Create)
		user.POST("/logout", handlerUser.Logout)
		user.DELETE("/delete/:id", handlerUser.Delete)
	}

	userAuth := user.Group("/user", myjwt.JWTAuth())
	{
		userAuth.GET("/data", handlerUser.GetName)
		userAuth.GET("/info", handlerUser.GetUserInfo)
	}

	//定义默认路由404
	r.NoRoute(handleNotFound)

	fmt.Println("host======" + host)
	r.Run(host + ":" + port)
}


func handleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404, page not ready to GO!",
	})
}
