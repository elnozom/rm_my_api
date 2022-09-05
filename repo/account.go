package repo

import (
	"fmt"
	"math"
	"mynozom/model"

	"github.com/jinzhu/gorm"
)

type AccountRepo struct {
	dbs []*gorm.DB
}

func NewAccountRepo(dbs []*gorm.DB) AccountRepo {
	return AccountRepo{
		dbs: dbs,
	}
}

func (ur *AccountRepo) BalanceOfTrade(req *model.GetBalanceOfTradeRequest, dbIndex uint) (*[]model.GetBalanceOfTradeResponse, error) {
	var resp []model.GetBalanceOfTradeResponse
	rows, err := ur.dbs[dbIndex].Raw("EXEC balanceoftrade @AccountType = ? , @DateFrom = ? , @DateTo = ?", req.AccType, req.FromDate, req.ToDate).Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.GetBalanceOfTradeResponse
		if err := rows.Scan(&rec.AccountCode, &rec.AccountName, &rec.BBD, &rec.BBC, &rec.BAD, &rec.BAC); err != nil {
			return nil, err
		}
		val := (rec.BBC + rec.BAC) - (rec.BBD + rec.BAD)
		if val > 0 {
			rec.Credit = math.Abs(val)
		} else {
			rec.Debit = math.Abs(val)
		}
		resp = append(resp, rec)
	}
	return &resp, nil
}

func (ur *AccountRepo) GetAccountBalance(req *model.GetAccountBalanceRequest, dbIndex uint) (*[]model.GetAccountBalanceData, error) {
	var data []model.GetAccountBalanceData
	rows, err := ur.dbs[dbIndex].Raw("EXEC AccTr01CashFlow @DateFrom = ?, @DateTo = ? , @AccSerial = ? ;", req.FromDate, req.ToDate, req.AccSerial).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// var raseed = math.Abs(resp.Raseed)
	for rows.Next() {
		var balance float64
		var serial int
		var rec model.GetAccountBalanceData
		err := rows.Scan(&rec.DocNo, &rec.AccMoveName, &rec.DocDate, &rec.Dbt, &rec.Crdt, &balance, &rec.Note, &rec.Account2, &serial)
		if err != nil {
			return nil, err
		}
		if balance > 0 {
			rec.RaseedDbt = balance
		} else {
			rec.RaseedCrdt = math.Abs(balance)
		}
		data = append(data, rec)
	}
	return &data, nil
}
func (ur *AccountRepo) MonthlyReport(req *model.MonthlyReportReq, dbIndex uint) (*[]model.MonthlyReportResp, error) {
	var resp []model.MonthlyReportResp
	fmt.Print(resp)
	rows, err := ur.dbs[dbIndex].Raw("EXEC Rpt_MonthlyReport @Year = ? , @Month = ?;", req.Year, req.Month).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rec model.MonthlyReportResp
		err := rows.Scan(
			&rec.DocDay,
			&rec.SessionNo,
			&rec.TotalCash,
			&rec.CashierMoney,
			&rec.AmountIn,
			&rec.Supplier,
			&rec.Expenses,
			&rec.BankIn,
			&rec.TotalVisa,
			&rec.Net,
		)
		if err != nil {
			return nil, err
		}
		resp = append(resp, rec)
	}

	return &resp, nil
}
