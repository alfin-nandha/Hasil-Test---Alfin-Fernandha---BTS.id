package routes

import (
	"btsid/test/factory"
	"btsid/test/middlewares"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {

	e := echo.New()

	e.Use(middlewares.CorsMiddleware())

	e.POST("/register", presenter.AuthPresenter.Register)
	e.POST("/login", presenter.AuthPresenter.Login)

	e.GET("/checklist", presenter.ChecklistPresenter.GetData, middlewares.JWTMiddleware())

	return e
}
