package controllers

import "github.com/labstack/echo/v4"

type GeneralController struct{}

func NewGeneralController() *GeneralController {
	return &GeneralController{}
}

func (c *GeneralController) Register(app *echo.Echo) error {
	app.GET("/", IndexHandler)
	return nil
}

func IndexHandler(c echo.Context) error {
	return c.String(200, "Hello, world")
}
