package router

import (
	"github.com/labstack/echo/v4"
	"go-101/controller"
)

func NewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	return e
}
