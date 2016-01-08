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
	fmt.Println(http.ListenAndServe(":4455", mux))
}

func Close() {
	db.CloseDb()
	fmt.Println("end listen")
}
