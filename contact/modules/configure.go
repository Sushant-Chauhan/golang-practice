package modules

import (
	"contactApp/app"
	"contactApp/models/user"
)

func ConfigureAppTables(appObj *app.App) {
	userModuleConfig := user.NewUserModuleConfig(appObj.DB, appObj.Log)
	appObj.TableMigration([]app.ModuleConfig{userModuleConfig})
}
