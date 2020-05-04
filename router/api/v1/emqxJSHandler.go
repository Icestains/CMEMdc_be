package v1

import (
	"CMEMdc_be/Database"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

//NewPostHandler ...
func NewEmqxJSHandler(db *gorm.DB) *EmqxJS {
	return &EmqxJS{
		service: Database.NewService(db),
	}
}

// Post ...
type EmqxJS struct {
	service Database.Service
}

// Fetch all post data
func (p *EmqxJS) Fetch(ctx *gin.Context) {
	payload, _ := p.service.EmqxJS.FindAll()
	fmt.Printf("from database:========== %T",payload)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"res":  "success",
		"data": payload,
	})
}
