package main

import (
	"go-echo-htmx-templ/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return renderView(c, views.HelloIndex("www"))
	})
	e.GET("/bt1", func(c echo.Context) error {
		return renderView(c, views.Button1())
	})
	e.GET("/bt2", func(c echo.Context) error {
		return renderView(c, views.Button2())
	})
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":1323"))
}