package handler

import (
	"mynozom/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func (h *Handler) MonthlyReport(c echo.Context) error {
	var req model.MonthlyReportReq
	dataParam := c.QueryParam("date")
	parsedData := strings.Split(dataParam, "-")
	req.Year, _ = strconv.Atoi(parsedData[0])
	req.Month, _ = strconv.Atoi(parsedData[1])
	dbIndex := c.Get("dbIndex").(uint)
	resp, err := h.accountRepo.MonthlyReport(&req, dbIndex)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error with procedure : "+err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
