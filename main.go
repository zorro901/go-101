package main

import (
	"go-101/controller"
	"go-101/db"
	"go-101/repository"
	"go-101/router"
	"go-101/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080")) // エラーが発生したらログ情報を生成して終了する
}
