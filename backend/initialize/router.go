package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "job-maching/docs"
	"job-maching/global"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	docs.SwaggerInfo.BasePath = global.CONFIG.System.RouterPrefix
	Router.GET(global.CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")

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
