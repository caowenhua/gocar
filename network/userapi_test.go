package network

import (
	"fmt"
	"github.com/Centny/gwf/util"
	"strconv"
	"testing"
	"time"
)

var host string
var regMobileSlice []string
var uidSlice []interface{}

func TestListen(t *testing.T) {
	// host = "http://127.0.0.1:4455/"
	// db.SetupDb()
	// ts := httptest.NewMuxServer()
	// ts.mux.HFunc("/user/reg", Register)
	// ts.mux.HFunc("/user/login", Login)
	// ts.mux.HFunc("/user/auth", AuthDriver)
	// ts.mux.HFunc("/user/fill", FillInfo)
	host = "http://localhost:4455"
	go Listen()
}

func TestReg(t *testing.T) {
	regMobileSlice = make([]string, 0)
	uidSlice = make([]interface{}, 0)
	fmt.Println("\n\n\n========================================================")
	fmt.Println("Test reg begin\n\n\n")
	ti := time.Now()
	fmt.Println("*************正常参数注册****************")
	head := "http://e.hiphotos.baidu.com/image/h%3D200/sign=bc09df35750e0cf3bff749fb3a46f23d/2fdda3cc7cd98d102db27e16263fb80e7bec90b6.jpg"
	hobby := "coding!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
	gender := 0
	for i := 0; i < 20; i++ {
		str := host + "/user/reg?userName=%v&mobile=%v&head=%v"
		name := "c" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		mobile := "m" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		regMobileSlice = append(regMobileSlice, mobile)
		vmap, err := util.HGet(str, name, mobile, head)
		uidSlice = append(uidSlice, vmap)
		fmt.Println(vmap, err)
	}
	for i := 0; i < 20; i++ {
		str := host + "/user/reg?userName=%v&mobile=%v&head=%v&gender=%v&hobby=%v"
		name := "c" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		mobile := "m" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		regMobileSlice = append(regMobileSlice, mobile)
		vmap, err := util.HGet(str, name, mobile, head, gender, hobby)
		uidSlice = append(uidSlice, vmap)
		fmt.Println(vmap, err)
	}
	fmt.Println("*************正常必要参数only注册****************")
	for i := 0; i < 20; i++ {
		str := host + "/user/reg?userName=%v&mobile=%v"
		name := "d" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		mobile := "d" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		regMobileSlice = append(regMobileSlice, mobile)
		vmap, err := util.HGet(str, name, mobile)
		uidSlice = append(uidSlice, vmap)
		fmt.Println(vmap)
		fmt.Println(vmap, err)
	}
	fmt.Println("*************缺省参数注册****************")
	for i := 0; i < 20; i++ {
		str := host + "/user/reg?userName=%v"
		name := "e" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		vmap, err := util.HGet(str, name)
		fmt.Println(vmap, err)
	}
	for i := 0; i < 20; i++ {
		str := host + "/user/reg?mobile=%v"
		mobile := "f" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		regMobileSlice = append(regMobileSlice, mobile)
		vmap, err := util.HGet(str, mobile)
		fmt.Println(vmap, err)
	}

	fmt.Println("\n\n\nTest reg end")
	fmt.Println("========================================================\n\n\n")
}

func TestLogin(t *testing.T) {

	fmt.Println("\n\n\n========================================================")
	fmt.Println("Test login begin\n\n\n")
	fmt.Println("*************正常保存后登陆****************")
	str := host + "/user/login?mobile=%v"
	for _, m := range regMobileSlice {
		vmap, err := util.HGet(str, m)
		fmt.Println(vmap, err)
	}
	fmt.Println("*************缺省参数登陆****************")
	for i := 0; i < 20; i++ {
		str := host + "/user/login?"
		vmap, err := util.HGet(str)
		fmt.Println(vmap, err)
	}
	for i := 0; i < 20; i++ {
		str := host + "/user/login?uid=1"
		vmap, err := util.HGet(str)
		fmt.Println(vmap, err)
	}
	fmt.Println("\n\n\nTest login end")
	fmt.Println("========================================================\n\n\n")
}

func TestFillInfo(t *testing.T) {
	head := "fill info head"
	hobby := "hobby!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
	gender := 1
	ti := time.Now()

	fmt.Println("\n\n\n========================================================")
	fmt.Println("Test fill info begin\n\n\n")
	fmt.Println("*************正常修改资料(后20个失败)****************")
	for i, m := range regMobileSlice {
		str := host + "/user/fill?userName=%v&head=%v&gender=%v&hobby=%v&mobile=%v"
		name := "ccc" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		vmap, err := util.HGet(str, name, head, gender, hobby, m)
		fmt.Println(vmap, err)
	}

	fmt.Println("*************缺省参数修改资料****************")
	for i := 0; i < 20; i++ {
		str := host + "/user/fill?userName=%v&head=%v&gender=%v&mobile=%v"
		name := "ccc" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		vmap, err := util.HGet(str, name, head, gender, regMobileSlice[i])
		fmt.Println(vmap, err)
	}
	for i := 0; i < 20; i++ {
		str := host + "/user/fill?head=%v&gender=%v&mobile=%v"
		vmap, err := util.HGet(str, head, gender, regMobileSlice[i])
		fmt.Println(vmap, err)
	}
	fmt.Println("\n\n\nTest fill info end")
	fmt.Println("========================================================\n\n\n")
}

func TestAuthDriver(t *testing.T) {
	fmt.Println("\n\n\n========================================================")
	fmt.Println("Test auth begin\n\n\n")
	fmt.Println("*************正常Auth****************")
	str := host + "/user/auth?uid=%v"
	for _, m := range uidSlice {
		vmap, err := util.HGet(str, m)
		fmt.Println(vmap, err)
	}
	fmt.Println("*************缺省参数auth****************")
	str = host + "/user/auth?mobile=123"
	for i := 0; i < 20; i++ {
		vmap, err := util.HGet(str)
		fmt.Println(vmap, err)
	}
	fmt.Println("\n\n\nTest auth end")
	fmt.Println("========================================================\n\n\n")
}

func TestDeleteUser(t *testing.T) {
	fmt.Println("\n\n\n========================================================")
	fmt.Println("Test delete begin\n\n\n")
	fmt.Println("*************正常delete****************")
	str := host + "/user/delbyid?uid=%v"
	for i := 0; i < 20; i++ {
		vmap, err := util.HGet(str, uidSlice[i])
		fmt.Println(vmap, err)
	}
	str = host + "/user/delbym?mobile=%v"
	for _, m := range uidSlice {
		vmap, err := util.HGet(str, m)
		fmt.Println(vmap, err)
	}
	fmt.Println("*************缺省参数delete****************")
	str = host + "/user/delbym?head=112"
	for i := 0; i < 20; i++ {
		vmap, err := util.HGet(str)
		fmt.Println(vmap, err)
	}
	fmt.Println("\n\n\nTest delete end")
	fmt.Println("========================================================\n\n\n")
}

func TestClose(t *testing.T) {
	Close()
}
