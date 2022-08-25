package repo

import (
	"fmt"
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
