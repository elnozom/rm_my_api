package handler

import (
	"mynozom/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) MainGroupsList(c echo.Context) error {
	dbIndex := c.Get("dbIndex").(uint)
	resp, err := h.dmenuRepo.MainGroupList(dbIndex)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) SubGroupsList(c echo.Context) error {
	dbIndex := c.Get("dbIndex").(uint)
	id, _ := strconv.Atoi(c.Param("group"))
	resp, err := h.dmenuRepo.SubGroupList(id, dbIndex)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) ItemsList(c echo.Context) error {
	dbIndex := c.Get("dbIndex").(uint)
	group, _ := strconv.Atoi(c.Param("group"))
	tableSerial, _ := strconv.Atoi(c.Param("tableSerial"))
	resp, err := h.dmenuRepo.ItemsList(group, tableSerial, dbIndex)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateImage(c echo.Context) error {
	dbIndex := c.Get("dbIndex").(uint)
	req := new(model.UpdateImageReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	resp, err := h.dmenuRepo.UpdateImage(req, dbIndex)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
