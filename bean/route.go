package bean

type Account struct {
	UserId   int64  `m2s:"userid" json:"userId"`
	Account  string `m2s:"account" json:"account"`
	Password string `m2s:"password" json:"password"`
}
