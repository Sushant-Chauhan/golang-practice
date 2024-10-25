package modules

import (
	"contactApp/app"
	"contactApp/components/user/controller"
	"contactApp/components/user/service"
)

func registerUserRoutes(appObj *app.App) {

	userService := service.NewUserService(appObj.DB, appObj.Repository, appObj.Log)
	UserController := controller.NewUserController(userService, appObj.Log)
	appObj.RegisterAllControllerRoutes([]app.Controller{UserController})
}
