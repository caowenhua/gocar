package bean

type DriverRoute struct {
	Uid int64 `m2s:"uid" json:"uid"`
	Rid int64 `m2s:"rid" json:"rid"`
	// StartTime int64   `m2s:"startTime" json:"startTime"`
	// SPlace    string  `m2s:"sPlace" json:"sPlace"`
	// EPlace    string  `m2s:"ePlace" json:"ePlace"`
	// SLat      float64 `m2s:"sLat" json:"sLat"`
	// SLng      float64 `m2s:"sLng" json:"sLng"`
	// ELat      float64 `m2s:"eLat" json:"eLat"`
	// ELng      float64 `m2s:"eLng" json:"eLng"`
	// SCity     string  `m2s:"sCity" json:"sCity"`
	// ECity     string  `m2s:"eCity" json:"eCity"`
	UnitPrice float64 `m2s:"unitPrice" json:"unitPrice"`
	Distance  float64 `m2s:"distance" json:"distance"`
	Drid      int64   `m2s:"drid" json:"drid"`
}
