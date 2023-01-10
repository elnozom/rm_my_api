package handler

import (
	"fmt"
	"mynozom/model"
	"mynozom/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CurrentUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "valid")
}

func (h *Handler) DefaultStoreUpdate(c echo.Context) error {

	return c.JSON(http.StatusOK, "r")
}

func (h *Handler) Login(c echo.Context) error {
	req := new(model.UserLoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	r := new(model.UserResponse)
	u, err := h.userRepo.GetByCode(&req.Username)
	if err != nil || u == nil {
		return c.JSON(http.StatusBadRequest, "incorrect_uname"+err.Error())
	}
	// return c.JSON(http.StatusOK, u)
	if u.EmpPassword != req.Password {
		return c.JSON(http.StatusBadRequest, "wrong_password")
	}
	r.User = *u
	r.Token = utils.GenerateJWT(uint(u.EmpCode), req.Store)
	return c.JSON(http.StatusOK, r)
}

func (h *Handler) GetUser(c echo.Context) error {
	id := c.Get("user").(uint)

	strId := strconv.FormatUint(uint64(id), 10)
	fmt.Println(id)
	u, err := h.userRepo.GetByCode(&strId)
	if err != nil || u == nil {
		return c.JSON(http.StatusBadRequest, "incorrect_uname"+err.Error())
	}

	return c.JSON(http.StatusOK, u)
}
