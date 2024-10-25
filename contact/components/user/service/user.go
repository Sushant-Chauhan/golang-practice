package service

import (
	"contactApp/models/user"
	"contactApp/repository"
	"contactApp/utils/log"
	"contactApp/utils/web"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	DB         *gorm.DB
	repository repository.Repository
	log        log.Logger
}

func NewUserService(db *gorm.DB, repository repository.Repository, log log.Logger) *UserService {
	return &UserService{
		DB:         db,
		repository: repository,
		log:        log,
	}
}

// limit?=12 && offset?=23 && name?=yash
func (u *UserService) CreateUser(newUser *user.User) error {
	//transaction
	uow := repository.NewUnitOfWork(u.DB)
	defer uow.RollBack()
	err := u.repository.Add(uow, newUser)
	if err != nil {
		return err
	}

	uow.Commit()
	return nil

}
func (u *UserService) GetAllUsers(allUsers []*user.User, totalCount *int, parser web.Parser) error {

	uow := repository.NewUnitOfWork(u.DB)
	defer uow.RollBack()
	//
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors,
		u.repository.Filter("name=?", parser.Form.Get("name")),
		u.repository.Filter("name=?", parser.Form.Get("name")),
		u.repository.Filter("name=?", parser.Form.Get("name")),
		u.repository.Count(parser.Form.Get("limit"), parser.Form.Get("offset"), totalCount),
		u.repository.Limit(parser.Form.Get("limit")),
		u.repository.Offset(parser.Form.Get("offset")),
	)
	u.repository.GetAll(uow, allUsers, queryProcessors...)

	//.where.where.where.Count.limit.offset .find

	uow.Commit()
	return nil

}
