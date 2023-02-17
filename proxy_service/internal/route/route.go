package route

import (
	"github.com/gin-gonic/gin"
	"github.com/roxyash/kmf_testtask/proxy_service/internal/handler"
	"net/http"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/roxyash/kmf_testtask/proxy_service/docs"
)

type Route struct {
	handler *handler.Handler
}

func NewRoute(handler *handler.Handler) *Route {
	return &Route{
		handler: handler,
	}
}

func (r *Route) InitRoutes() *gin.Engine {
	router := gin.New()
	//Swagger docs . . .
	router.GET("/swagger", gin.WrapH(http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently)))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := router.Group("api/v1")

	proxy := apiV1.Group("proxy")
	{
		proxy.POST("/", r.handler.Proxy)
	}

	return router
}
