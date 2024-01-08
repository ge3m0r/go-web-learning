package bootstrap

import (
	"golearning/app/http/middlewares"
	"golearning/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)
func SetupRoute(router *gin.Engine) {

    // 注册全局中间件
    registerGlobalMiddleWare(router)

    //  注册 API 路由
    routes.RegisterAPIRoutes(router)

    //  配置 404 路由
    setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine){
	router.Use(
		middlewares.Logger(),
		middlewares.Recovery(),
	)
}

func setup404Handler(router *gin.Engine){
	router.NoRoute(func(c *gin.Context) {
		accpetString := c.Request.Header.Get("Accept")
		if strings.Contains(accpetString,"text/html") {
			c.String(http.StatusNotFound,"页面返回 404")
		}else{
			c.JSON(http.StatusNotFound,gin.H{
				"error_code" : 404,
				"error_message":"路由为定义，请确定url和请求方法是否正确",
			})
		}
	})
}