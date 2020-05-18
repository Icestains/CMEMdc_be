package v1

import (
	"CMEMdc_be/utils/app"
	"net/http"

	"github.com/gin-gonic/gin"

	"CMEMdc_be/models"
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
