package network

import (
	"fmt"
	"net/http"

	"github.com/Centny/gwf/routing"

	"me.car/db"
)

func Listen() {
	db.SetupDb()
	fmt.Println("begin listen")
	sb := routing.NewSrvSessionBuilder("", "/", "example", 60*60*1000, 10000)
	mux := routing.NewSessionMux("", sb)
	// mux.URL = "127.0.0.1"
	mux.HFunc("/user/reg", Register)
	mux.HFunc("/user/login", Login)
	mux.HFunc("/user/auth", AuthDriver)
	mux.HFunc("/user/fill", FillInfo)
	mux.HFunc("/user/delbyid", DeleteUserById)
	mux.HFunc("/user/delbym", DeleteUserByMobile)

	mux.HFunc("/route/dadd", DriverAddRoute)
	mux.HFunc("/route/padd", PassengerJoinRoute)
	mux.HFunc("/route/pfind", PassengerFindDriver)

	mux.HFunc("/order/charge", ChargeBalance)
	mux.HFunc("/order/withdraw", WithDrawBalance)
	mux.HFunc("/order/dcancel", DriverCancelOrder)
	mux.HFunc("/order/pcancel", PassengerCancelOrder)
	mux.HFunc("/order/dlist", GetDriverOrderList)
	mux.HFunc("/order/plist", GetPassengerOrderList)
	// mux.HFunc("/user/t", TestMethod)
	fmt.Println(http.ListenAndServe(":4455", mux))
}

func Close() {
	db.CloseDb()
	fmt.Println("end listen")
}

func TestMethod(hs *routing.HTTPSession) routing.HResult {
	balance, err := db.Method()
	fmt.Println(balance, err)
	return hs.MsgRes(balance)
}
