package model

type CashFlow struct {
	DocDate     string
	Customer    float32
	OtherIn     float32
	FromBank    float32
	Supplier    float32
	Expenses    float32
	OtherOut    float32
	ToBank      float32
	TotalDebit  float32
	TotalCredit float32
}

type CashFlowYearReq struct {
	Year      string
	AccSerial int
}

type CashFlowReq struct {
	FromDate  string
	ToDate    string
	AccSerial int
}
