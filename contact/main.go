package main

import (
	"contactApp/app"
	"contactApp/modules"
	"contactApp/repository"
	"contactApp/utils/log"
	"sync"
)

func main() {
	//Logger
	log := log.GetLogger()
	//DB Connections
	db := app.NewDBConnection(log)
	if db == nil {
		log.Error("Db connection failed.")
	}
	defer func() {
		db.Close()
		log.Error("Db closed")
	}()
	var wg *sync.WaitGroup

	repository := repository.NewGormRepositoryMySQL()
	app := app.NewApp("Contact-App", db, log, wg, repository)
	//initilised my router and server
	app.Init()
	//register my routes
	modules.RegisterAllRoutes(app)
	/// models , migrations,
	modules.ConfigureAppTables(app)
	go func() {
		err := app.StartServer()
		if err != nil {
			stopApp(app)
		}
	}()

}
func stopApp(app *app.App) {
	//stop this app
}
