package xxl

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ginMux(e *gin.Engine, exec Executor) {
	xxlTokenValid := tokenValid(exec.GetAccessToken())
	//注册的gin的路由
	e.POST("run", xxlTokenValid, gin.WrapF(exec.RunTask))
	e.POST("kill", xxlTokenValid, gin.WrapF(exec.KillTask))
	e.POST("log", xxlTokenValid, gin.WrapF(exec.TaskLog))
	e.POST("beat", xxlTokenValid, gin.WrapF(exec.Beat))
	e.POST("idleBeat", xxlTokenValid, gin.WrapF(exec.IdleBeat))
}

func tokenValid(accessToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("XXL-JOB-ACCESS-TOKEN")
		if accessToken != token {
			//上面是handler执行之前执行
			c.JSON(http.StatusForbidden, gin.H{
				"code": http.StatusForbidden,
				"msg":  "Token认证失败",
			})
			c.Abort()
			return
		}

		c.Next()
		//下面是之后执行

	}
}
