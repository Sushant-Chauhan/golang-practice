package controller

import (
	"contactApp/components/user/service"
	"contactApp/models/user"
	"contactApp/utils/log"
	"contactApp/utils/web"
	"net/http"

	"github.com/gorilla/mux"
)

type UserController struct {
	UserService *service.UserService
	log         log.Logger
}

func NewUserController(UserService *service.UserService,
	log log.Logger) *UserController {
	return &UserController{
		UserService: UserService,
		log:         log,
	}
}

func (u *UserController) RegisterRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", u.CreateUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/", u.GetAllUsers).Methods(http.MethodGet)

}
func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	u.log.Info("CreateUser called")
	newUser := user.User{}
	_ = web.UnMarshalJSON(r, &newUser)
	//newUser validation
	u.UserService.CreateUser(&newUser)
	//your code here

}
func (u *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	u.log.Info("GetAllUsers called")
	parser := web.NewParser(r)
	//validation
	allUsers := []*user.User{}
	var totalCount int
	u.UserService.GetAllUsers(allUsers, &totalCount, *parser)
	//your code here

}
