package handler

import (
	"mynozom/repo"

	"github.com/jinzhu/gorm"
)

type Handler struct {
	dbs      []*gorm.DB
	userRepo repo.UserRepo
}

func NewHandler(
	databaases []*gorm.DB,
	userRepo repo.UserRepo,
) *Handler {
	return &Handler{
		dbs:      databaases,
		userRepo: userRepo,
	}
}
