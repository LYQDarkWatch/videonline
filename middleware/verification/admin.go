package verification

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"videOnline/pkg/error"
)

//验证当前用户是否管理员
func Verification() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int

		priority := c.Query("priority")
		priorityint, _ := strconv.Atoi(priority)
		if priorityint == 2 {
			code = error.SUCCESS
		} else {
			code = error.ERROR_USER_NOT_ADMIN
		}
		if code != error.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  error.GetMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
