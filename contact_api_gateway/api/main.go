package api

import (
	"contact_api_gateway/config"
	"contact_api_gateway/pkg/logger"
	"contact_api_gateway/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// @Summary 登录
	// @Description 登录
	// @Produce json
	// @Param body body controllers.LoginParams true "body参数"
	// @Success 200 {string} string "ok" "返回用户信息"
	// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
	// @Failure 401 {string} string "err_code：10001 登录失败"
	// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
	// @Router /user/person/login [post]
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "contact_api_gateway/api/docs"
	v1 "contact_api_gateway/api/handlers/v1"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServiceManager
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	// router.GET("/config", handlerV1.GetConfig)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handlerV1.SignUp)
		auth.POST("/sign-in", handlerV1.SignIn)
	}

	apiV1 := router.Group("/v1", handlerV1.UserIdentify)

	apiV1.POST("/contact", handlerV1.CreateContact)
	// apiV1.GET("/", handlerV1.getAllContact)
	// apiV1.GET("/:id", handlerV1.getContact)
	// apiV1.PUT("/:id", handlerV1.updateContact)
	// apiV1.DELETE("/:id", handlerV1.deleteContact)

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
