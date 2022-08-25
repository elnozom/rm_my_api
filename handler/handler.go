package handler

import (
	"mynozom/repo"

	"github.com/jinzhu/gorm"
)

type Handler struct {
	dbs         []*gorm.DB
	userRepo    repo.UserRepo
	accountRepo repo.AccountRepo
}

func NewHandler(
	databaases []*gorm.DB,
	userRepo repo.UserRepo,
	accountRepo repo.AccountRepo,
) *Handler {
	return &Handler{
		dbs:         databaases,
		userRepo:    userRepo,
		accountRepo: accountRepo,
	}
}
