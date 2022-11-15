package handler

import (
	"mynozom/router/middleware"
	"mynozom/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)

	v1.GET("/me", h.Login, jwtMiddleware)
	v1.POST("/upload", h.Upload)

	v1.POST("/login", h.Login)
	v1.PUT("/store", h.DefaultStoreUpdate, jwtMiddleware)

	cashtry := v1.Group("/cashtray", jwtMiddleware)
	cashtry.GET("", h.CashTryAnalysis)
	cashtry.GET("/open", h.OpenCashTry)
	cashtry.GET("/paused", h.PausedCashTry)
	cashtry.GET("/stores", h.CashTryStores)

	v1.GET("/groups", h.GetGroups, jwtMiddleware)
	v1.GET("/stock", h.GetStock, jwtMiddleware)
	v1.GET("/top", h.GetTopSalesItem, jwtMiddleware)
	v1.GET("/cash-flow", h.GetCashFlow, jwtMiddleware)
	v1.GET("/cash-flow-year", h.GetCashFlowYear, jwtMiddleware)
	v1.GET("/balance-of-trade", h.GetBalnaceOfTrade, jwtMiddleware)
	v1.GET("/branches-sales", h.GetBranchesSales, jwtMiddleware)
	v1.GET("/branches-profit", h.GetBranchesProfit, jwtMiddleware)
	v1.GET("/monthly-sales", h.GetMonthlySales, jwtMiddleware)
	v1.GET("/daily-sales", h.GetDailySales, jwtMiddleware)
	v1.GET("/get-emp-totals", h.EmpTotals, jwtMiddleware)
	v1.GET("/get-drivers", h.GetDrivers, jwtMiddleware)
	v1.GET("/trans-cycle-acc", h.GetTransCycleAcc, jwtMiddleware)
	v1.GET("/get-item", h.GetItem, jwtMiddleware)
	v1.GET("/revenue", h.GetRevenuePerTime, jwtMiddleware)
	v1.GET("/expenses", h.GetExpnsesByMonth, jwtMiddleware)
	v1.POST("/pay", h.AccTr01Insert, jwtMiddleware)
	v1.POST("/get-doc", h.GetDocNo, jwtMiddleware)
	v1.POST("/get-doc-items", h.GetDocItems, jwtMiddleware)
	v1.POST("/insert-item", h.InsertItem, jwtMiddleware)
	v1.POST("/delete-item", h.DeleteItem, jwtMiddleware)
	v1.POST("/get-docs", h.GetOpenDocs, jwtMiddleware)
	v1.GET("/cashtray/data/:serial", h.GetCashtrayData, jwtMiddleware)
	// account routes
	v1.GET("/supplier-balance", h.GetSupplierBalance, jwtMiddleware)
	v1.GET("/get-account-balance", h.GetAccountBalance, jwtMiddleware)
	v1.GET("/balance/before", h.GetAccountBalanceBefore, jwtMiddleware)
	v1.GET("/get-account", h.GetAccount, jwtMiddleware)
	v1.GET("/monthly-report", h.MonthlyReport, jwtMiddleware)

	groupG := v1.Group("/group", jwtMiddleware)
	groupG.GET("", h.MainGroupsList)
	groupG.GET("/:group", h.SubGroupsList)

	itemsG := v1.Group("/item", jwtMiddleware)
	itemsG.GET("/:group/:tableSerial", h.ItemsList)
	itemsG.PUT("/image", h.UpdateImage)

}
