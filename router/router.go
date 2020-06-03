package router

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	_ "CMEMdc_be/docs"
	"CMEMdc_be/middleware/cors"
	"CMEMdc_be/middleware/jwt"
	"CMEMdc_be/router/api"
	"CMEMdc_be/router/api/v1"
	"CMEMdc_be/utils/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.Use(cors.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	gin.SetMode(setting.ServerSetting.RunMode)

	r.POST("/auth", api.GetAuth)
	r.POST("/register", v1.Create)

	apiv1 := r.Group("/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/emqxdata", v1.FindAllEmqxData)
		apiv1.GET("/emqxclient", v1.FindAllEmqxClientInfo)
		apiv1.GET("/sender/:Sender", v1.FindEmqxDataBySender)
		apiv1.GET("/topic", v1.FindEmqxDataByTopic)
	}

	user := apiv1.Group("user")
	{
		user.POST("/logout", v1.Logout)
		//user.DELETE("/delete/:id", handlerUser.Delete)
		user.GET("/info", v1.GetUserInfo)
	}

	//定义默认路由404
	r.NoRoute(handleNotFound)

	return r
}

func handleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404, page not ready to GO!",
	})
}
