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
	tasksRepository := repository.NewTaskRepository(db)

	userUseCase := usecase.NewUserUseCase(userRepository)
	tasksUseCase := usecase.NewTaskUseCase(tasksRepository)

	userController := controller.NewUserController(userUseCase)
	tasksController := controller.NewTaskController(tasksUseCase)

	e := router.NewRouter(userController, tasksController)
	e.Logger.Fatal(e.Start(":8080")) // エラーが発生したらログ情報を生成して終了する
}
