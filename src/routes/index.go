package router

import (
	"go-crud/src/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	handler.InitItems()
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", handler.Hello)
	e.GET("/all", handler.GetAll)
	e.GET("/detail/:id", handler.GetOne)
	e.POST("/item", handler.CreateItem)
	e.PUT("/item/:id", handler.UpdateItem)
	e.DELETE("/item/:id", handler.RemoveItem)
	return e
}
