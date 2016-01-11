package bean

type PassengerOrderResponse struct {
	Oid       int64   `m2s:"oid" json:"oid"`
	Price     float64 `m2s:"price" json:"price"`
	OrderTime int64   `m2s:"orderTime" json:"orderTime"`
	OrderNo   string  `m2s:"orderNo" json:"orderNo"`
	Drid      int64   `m2s:"drid" json:"drid"`
	IsEnable  bool    `m2s:isEnable" json:"isEnable"`
	Duid      int64   `m2s:"duid" json:"duid"`

	UserName string `m2s:"userName" json:"userName"`
	Moblie   string `m2s:"mobile" json:"mobile"`
	Head     string `m2s:"head" json:"head"`
	Gender   int    `m2s:"gender" json:"gender"`
}
