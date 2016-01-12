package bean

type Order struct {
	Oid       int64   `m2s:"oid" json:"oid"`
	Price     float64 `m2s:"price" json:"price"`
	OrderTime int64   `m2s:"orderTime" json:"orderTime"`
	OrderNo   string  `m2s:"orderNo" json:"orderNo"`
	Uid       int64   `m2s:uid" json:"uid"`
	IsEnable  bool    `m2s:isEnable" json:"isEnable"`
}
