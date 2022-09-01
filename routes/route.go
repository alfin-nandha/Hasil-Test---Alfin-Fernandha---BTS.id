package routes

import (
	"btsid/test/factory"
	"btsid/test/middlewares"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {

	e := echo.New()
	e.Use(middlewares.CorsMiddleware())
	middlewares.LogMiddlewares(e)

	e.POST("/register", presenter.AuthPresenter.Register)
	e.POST("/login", presenter.AuthPresenter.Login)

	e.GET("/checklist", presenter.ChecklistPresenter.GetData, middlewares.JWTMiddleware())
	e.POST("/checklist", presenter.ChecklistPresenter.InsertData, middlewares.JWTMiddleware())

	return e
}
