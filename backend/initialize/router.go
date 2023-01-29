package initialize

import (
	"github.com/gin-gonic/gin"
	"job-maching/global"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	//docs.SwaggerInfo.BasePath = global.GB_CONFIG.System.RouterPrefix
	//Router.GET(global.GB_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//global.GB_LOG.Info("register swagger handler")

	PublicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	//{
	//	systemRouter.InitBaseRouter(PublicGroup)
	//	systemRouter.InitInitRouter(PublicGroup)
	//}
	//PrivateGroup := Router.Group(global.GB_CONFIG.System.RouterPrefix)
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//{
	//
	//}

	global.LOG.Info("router register success")
	return Router
}
