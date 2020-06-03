package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"CMEMdc_be/models"
	"CMEMdc_be/utils/app"
	"CMEMdc_be/utils/e"
)

// @Summary 查询所有 emqx 数据信息
// @Produce  json
// @Success 200 {object} app.Response
// @Router /v1/emqx [get]
func FindAllEmqxData(c *gin.Context) {
	appG := app.Gin{C: c}

	payload := models.FindAllEmqxData()

	appG.Response(http.StatusOK, e.SUCCESS, payload)
}

func FindAllEmqxClientInfo(c *gin.Context) {
	appG := app.Gin{C: c}

	payload := models.FindAllEmqxClients()

	appG.Response(http.StatusOK, e.SUCCESS, payload)
}

func FindEmqxDataBySender(c *gin.Context) {
	appG := app.Gin{C: c}
	Sender := c.Param("Sender")
	payload := models.FindEmqxDataBySender(Sender)
	appG.Response(http.StatusOK, e.SUCCESS, payload)
}

func FindEmqxDataByTopic(c *gin.Context) {
	appG := app.Gin{C: c}

	topic := c.Query("topic")
	fmt.Println(topic)

	payload := models.FindEmqxDataByTopic(topic)

	appG.Response(http.StatusOK, e.SUCCESS, payload)
}


