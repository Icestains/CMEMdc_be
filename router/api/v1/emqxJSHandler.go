package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"CMEMdc_be/models"
	"CMEMdc_be/utils/e"
)

// @Summary 查询所有 emqx 数据信息
// @Produce  json
// @Success 200 {object} app.Response
// @Router /v1/emqx [get]
func FindAllEmqxData(ctx *gin.Context) {
	code := e.SUCCESS

	payload := models.FindAllEmqxData()

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"res":  e.GetMsg(code),
		"data": payload,
	})
}
