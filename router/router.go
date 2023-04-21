package router

import (
	"github.com/labstack/echo/v4"
	"go-101/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	// tasksのエンドポイントをまとめる
	t := e.Group("/tasks")
	// tasksのエンドポイントにミドルウェアを設定する
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	// tasks以下のエンドポイントの挙動を設定する
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	return e
}
