package modules

import (
	"contactApp/app"
	"contactApp/components/contact/controller"
	"contactApp/components/contact/service"
)

func RegisterContactRoutes(appObj *app.App) {
	contactService := service.NewContactService(appObj.DB, appObj.Log)
	contactController := controller.NewContactController(*contactService, appObj.Log)

	appObj.RegisterAllControllerRoutes([]app.Controller{contactController})

}
