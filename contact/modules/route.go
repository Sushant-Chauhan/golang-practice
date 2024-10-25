package modules

import "contactApp/app"

func RegisterAllRoutes(app *app.App) {
	registerUserRoutes(app)
	RegisterContactRoutes(app)
}
