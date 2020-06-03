package jwt

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"CMEMdc_be/utils"
	"CMEMdc_be/utils/e"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS

		fmt.Println(c.Request.Header.Get("Token"))
		fmt.Println(c.Request.Header.Get("Authorization"))
		token := c.Request.Header.Get("Token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
