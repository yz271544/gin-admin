package router

import (
	"github.com/LyricTian/gin-admin/v6/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")

	g.Use(middleware.RateLimiterMiddleware())

	v1 := g.Group("/v1")
	{

		gDemo := v1.Group("demos")
		{
			gDemo.GET("", a.DemoAPI.Query)
			gDemo.GET(":id", a.DemoAPI.Get)
			gDemo.POST("", a.DemoAPI.Create)
			gDemo.PUT(":id", a.DemoAPI.Update)
			gDemo.DELETE(":id", a.DemoAPI.Delete)
			gDemo.PATCH(":id/enable", a.DemoAPI.Enable)
			gDemo.PATCH(":id/disable", a.DemoAPI.Disable)
		}
	}
}
