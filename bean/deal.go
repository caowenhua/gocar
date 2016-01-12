package bean

type Deal struct {
	Dealid    int64   `m2s:"dealid" json:"dealid"`
	DealTime  int64   `m2s:"dealTime" json:"dealTime"`
	DealPrice float64 `m2s:"dealPrice" json:"dealPrice"`
	Oid       int64   `m2s:"oid" json:"oid"`
	Uid       int64   `m2s:"uid" json:"uid"`
}
