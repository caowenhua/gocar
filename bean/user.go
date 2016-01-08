package bean

type User struct {
	UserName string  `m2s:"userName" json:"userName"`
	Moblie   string  `m2s:"mobile" json:"mobile"`
	Uid      int64   `m2s:"uid" json:"uid"`
	Hobby    string  `m2s:"hobby" json:"hobby"`
	Head     string  `m2s:"head" json:"head"`
	IsDriver bool    `m2s:"isDriver" json:"isDriver"`
	Balance  float64 `m2s:"balance" json:"balance"`
	Gender   int     `m2s:"gender" json:"gender"`
}
