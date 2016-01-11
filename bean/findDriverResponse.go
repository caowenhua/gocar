package bean

type FindDriverRoute struct {
	StartTime int64   `m2s:"startTime" json:"startTime"`
	SPlace    string  `m2s:"sPlace" json:"sPlace"`
	EPlace    string  `m2s:"ePlace" json:"ePlace"`
	SLat      float64 `m2s:"sLat" json:"sLat"`
	SLng      float64 `m2s:"sLng" json:"sLng"`
	ELat      float64 `m2s:"eLat" json:"eLat"`
	ELng      float64 `m2s:"eLng" json:"eLng"`
	SCity     string  `m2s:"sCity" json:"sCity"`
	ECity     string  `m2s:"eCity" json:"eCity"`
	Price     float64 `m2s:"price" json:"price"`
	Drid      int64   `m2s:"drid" json:"drid"`
}
