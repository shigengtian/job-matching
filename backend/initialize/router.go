package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"job-maching/api/v1"
	"job-maching/global"
)

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func Routers() *gin.Engine {
	Router := gin.Default()
	//docs.SwaggerInfo.Title = "Swagger Example API"
	//docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	//docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	//docs.SwaggerInfo.BasePath = "/v2"
	//docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//docs.SwaggerInfo.BasePath = global.CONFIG.System.RouterPrefix
	Router.GET(global.CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")
	// programmatically set swagger info
	//Router.GET("/", api.Helloworld)

	PublicGroup := Router.Group(global.CONFIG.System.RouterPrefix)
	{
		PublicGroup.GET("/", v1.Helloworld)
	}
	//{
	//	systemRouter.InitBaseRouter(PublicGroup)
	//	systemRouter.InitInitRouter(PublicGroup)
	//}

	//PrivateGroup := Router.Group(global.GB_CONFIG.System.RouterPrefix)
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//{
	//}

	global.LOG.Info("router register success")
	return Router
}
