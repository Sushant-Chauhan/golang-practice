package controller

import (
	"contactApp/components/contact/service"
	"contactApp/utils/log"
	"net/http"

	"github.com/gorilla/mux"
)

type ContactController struct {
	contactService *service.ContactService
	log            log.Logger
}

func NewContactController(ContactService service.ContactService,
	log log.Logger) *ContactController {
	return &ContactController{
		log:            log,
		contactService: &ContactService,
	}
}

func (c *ContactController) RegisterRoutes(router *mux.Router) {
	contactRouter := router.PathPrefix("/{userID}/contact").Subrouter()
	contactRouter.HandleFunc("/", c.CreateContact).Methods(http.MethodPost)
}
func (c *ContactController) CreateContact(w http.ResponseWriter, r *http.Request) {
	c.log.Info("CreateContact called")
	//your code here
}
