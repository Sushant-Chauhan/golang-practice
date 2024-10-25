package app

import (
	"contactApp/repository"
	"contactApp/utils/log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type App struct {
	sync.Mutex
	Name       string
	Router     *mux.Router
	DB         *gorm.DB
	Log        log.Logger
	Server     *http.Server
	WG         *sync.WaitGroup
	Repository repository.Repository
}

func NewApp(name string, db *gorm.DB, log log.Logger,
	wg *sync.WaitGroup, repo repository.Repository) *App {
	return &App{
		Name:       name,
		DB:         db,
		Log:        log,
		WG:         wg,
		Repository: repo,
	}
}
func NewDBConnection(log log.Logger) *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/user_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	db.LogMode(true)
	return db
}
func (app *App) Init() {
	app.initializeRouter()
	app.initializeServer()
}
func (app *App) initializeRouter() {
	app.Log.Info(app.Name + " App Route initializing")
	app.Router = mux.NewRouter().StrictSlash(true)
	app.Router = app.Router.PathPrefix("/api/v1/contactapp").Subrouter()
}
func (app *App) initializeServer() {
	headers := handlers.AllowedHeaders([]string{
		"Content-Type", "X-Total-Count", "token",
	})
	methods := handlers.AllowedMethods([]string{
		http.MethodPost, http.MethodPut, http.MethodGet, http.MethodDelete, http.MethodOptions,
	})
	originOption := handlers.AllowedOriginValidator(app.checkOrigin)
	app.Server = &http.Server{
		Addr:         "0.0.0.0:4000",
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS(headers, methods, originOption)(app.Router),
	}
	app.Log.Info("Server Exposed On 4000")
}
func (app *App) checkOrigin(origin string) bool {
	// origin will be the actual origin from which the request is made.

	return true
}

type Controller interface {
	RegisterRoutes(router *mux.Router)
}

func (app *App) RegisterAllControllerRoutes(controllers []Controller) {
	for i := 0; i < len(controllers); i++ {
		//router are created here
		controllers[i].RegisterRoutes(app.Router)
	}
}

func (app *App) StartServer() error {
	err := app.Server.ListenAndServe()
	if err != nil {
		app.Log.Error(err.Error())
		return err
	}
	return nil
}

type ModuleConfig interface {
	TableMigration()
}

func (app *App) TableMigration(moduleConfigs []ModuleConfig) {
	for i := 0; i < len(moduleConfigs); i++ {
		moduleConfigs[i].TableMigration()
	}
}
