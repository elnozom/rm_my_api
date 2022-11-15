package model

type MainGroup struct {
	GroupCode int
	GroupName string

	Icon string
}

type SubGroup struct {
	GroupCode int
	GroupName string
	ImagePath string
}

type DmenuItem struct {
	ItemSerial      int
	ItemPrice       float32
	ItemCode        int
	ItemName        string
	ItemDesc        string
	ImagePath       string
	WithModifier    bool
	Screen          int
	ScreenTimes     int
	OrderItemSerial int
	Qnt             float32
	MainModSerial   int
	AddItems        string
	Printed         bool
}

type UpdateImageReq struct {
	Image         string `json:"image"`
	MainGroupCode int    `json:"mainGroupCode"`
	ProductCode   int    `json:"productCode"`
	SubGroupCode  int    `json:"subGroupCode"`
}
