package bean

type OrderRoute struct {
	Oid      int64 `m2s:"oid" json:"oid"`
	Drid     int64 `m2s:"drid" json:"drid"`
	Uid      int64 `m2s:uid" json:"uid"`
	IsEnable bool  `m2s:isEnable" json:"isEnable"`
	Duid     int64 `m2s:"duid" json:"duid"`
}
