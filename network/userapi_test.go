package network

import (
	"fmt"
	"github.com/Centny/gwf/util"
	"os"
	"strconv"
	"testing"
	"time"
)

var host string
var regMobileSlice []string
var uidSlice []interface{}
var buff *os.File

func TestListen(t *testing.T) {
	// host = "http://127.0.0.1:4455/"
	// db.SetupDb()
	// ts := httptest.NewMuxServer()
	// ts.mux.HFunc("/user/reg", Register)
	// ts.mux.HFunc("/user/login", Login)
	// ts.mux.HFunc("/user/auth", AuthDriver)
	// ts.mux.HFunc("/user/fill", FillInfo)
	host = "http://localhost:4455"
	regMobileSlice = make([]string, 0)
	uidSlice = make([]interface{}, 0)
	// buff = make([]byte, 0)
	var outputError error
	buff, outputError = os.OpenFile("userapi_test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		panic(outputError)
	}
	go Listen()
}

func TestReg(t *testing.T) {
	buff.WriteString("\n\n\n========================================================\n")
	buff.WriteString("Test reg begin\n\n\n\n")
	fmt.Println("\n\n\n========================================================")
	fmt.Println("Test reg begin\n\n\n")
	ti := time.Now()
	buff.WriteString("*************正常参数注册****************\n")
	fmt.Println("*************正常参数注册****************")
	head := "http://e.hiphotos.baidu.com/image/h%3D200/sign=bc09df35750e0cf3bff749fb3a46f23d/2fdda3cc7cd98d102db27e16263fb80e7bec90b6.jpg"
	hobby := "coding!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
	gender := 0
	for i := 0; i < 10; i++ {
		str := host + "/user/reg?userName=%v&mobile=%v&head=%v"
		name := "c" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		mobile := "m" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		regMobileSlice = append(regMobileSlice, mobile)
		var vmap map[string]interface{}
		s, err := util.HGet(str, name, mobile, head)
		vmap, _ = util.Json2Map(s)
		uidSlice = append(uidSlice, vmap["data"])
		writeString(s, err)
		fmt.Println(vmap, err)
	}
	for i := 0; i < 10; i++ {
		str := host + "/user/reg?userName=%v&mobile=%v&head=%v&gender=%d&hobby=%v"
		name := "d" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		mobile := "n" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		regMobileSlice = append(regMobileSlice, mobile)
		var vmap map[string]interface{}
		s, err := util.HGet(str, name, mobile, head, gender, hobby)
		vmap, _ = util.Json2Map(s)
		uidSlice = append(uidSlice, vmap["data"])
		writeString(s, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("*************正常必要参数only注册****************\n")
	fmt.Println("*************正常必要参数only注册****************")
	for i := 0; i < 10; i++ {
		str := host + "/user/reg?userName=%v&mobile=%v"
		name := "e" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		mobile := "o" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		regMobileSlice = append(regMobileSlice, mobile)
		var vmap map[string]interface{}
		s, err := util.HGet(str, name, mobile)
		vmap, _ = util.Json2Map(s)
		uidSlice = append(uidSlice, vmap["data"])
		writeString(s, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("*************缺省参数注册****************\n")
	fmt.Println("*************缺省参数注册****************")
	for i := 0; i < 10; i++ {
		str := host + "/user/reg?userName=%v"
		name := "f" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		vmap, err := util.HGet(str, name)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	for i := 0; i < 10; i++ {
		str := host + "/user/reg?mobile=%v"
		mobile := "g" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		regMobileSlice = append(regMobileSlice, mobile)
		vmap, err := util.HGet(str, mobile)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}

	buff.WriteString("\n\n\nTest reg end\n")
	buff.WriteString("========================================================\n\n\n\n")
	fmt.Println("\n\n\nTest reg end")
	fmt.Println("========================================================\n\n\n")
}

func TestLogin(t *testing.T) {
	buff.WriteString("\n\n\n========================================================\n")
	buff.WriteString("Test login begin\n\n\n\n")
	buff.WriteString("*************正常保存后登陆（后10个mobile无效）****************\n")
	fmt.Println("\n\n\n========================================================")
	fmt.Println("Test login begin\n\n\n")
	fmt.Println("*************正常保存后登陆（后10个mobile无效）****************")
	str := host + "/user/login?mobile=%v"
	for _, m := range regMobileSlice {
		vmap, err := util.HGet(str, m)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("*************缺省参数登陆****************\n")
	fmt.Println("*************缺省参数登陆****************")
	for i := 0; i < 10; i++ {
		str := host + "/user/login?"
		vmap, err := util.HGet(str)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	for i := 0; i < 10; i++ {
		str := host + "/user/login?uid=1"
		vmap, err := util.HGet(str)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("\n\n\nTest login end\n")
	buff.WriteString("========================================================\n\n\n\n")
	fmt.Println("\n\n\nTest login end")
	fmt.Println("========================================================\n\n\n")
}

func TestFillInfo(t *testing.T) {
	head := "fill_info_head"
	hobby := "hobby!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
	gender := 1
	ti := time.Now()

	buff.WriteString("\n\n\n========================================================\n")
	buff.WriteString("Test fill info begin\n\n\n\n")
	buff.WriteString("*************正常修改资料****************\n")
	fmt.Println("\n\n\n========================================================")
	fmt.Println("Test fill info begin\n\n\n")
	fmt.Println("*************正常修改资料****************")
	for i, m := range uidSlice {
		str := host + "/user/fill?userName=%v&head=%v&gender=%v&hobby=%v&uid=%v"
		name := "ccc" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		vmap, err := util.HGet(str, name, head, gender, hobby, m)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}

	buff.WriteString("*************缺省参数修改资料****************\n")
	fmt.Println("*************缺省参数修改资料****************")
	for i := 0; i < 10; i++ {
		str := host + "/user/fill?userName=%v&head=%v&gender=%v&uid=%v"
		name := "ccc" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		vmap, err := util.HGet(str, name, head, gender, uidSlice[i])
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	for i := 0; i < 10; i++ {
		str := host + "/user/fill?head=%v&gender=%v&uid=%v"
		vmap, err := util.HGet(str, head, gender, uidSlice[i])
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("\n\n\nTest fill info end\n")
	buff.WriteString("========================================================\n\n\n\n")
	fmt.Println("\n\n\nTest fill info end")
	fmt.Println("========================================================\n\n\n")
}

func TestChargeBalance(t *testing.T) {
	buff.WriteString("\n\n\n========================================================\n")
	buff.WriteString("Test ChargeBalance begin\n\n\n\n")
	buff.WriteString("*************正常ChargeBalance(后20个uid失败)****************\n")
	fmt.Println("\n\n\n========================================================")
	fmt.Println("Test ChargeBalance begin\n\n\n")
	fmt.Println("*************正常ChargeBalance(后20个uid失败)****************")
	str := host + "/user/charge?uid=%v&money=%v"
	for _, m := range uidSlice {
		vmap, err := util.HGet(str, m, 10)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("*************错误参数ChargeBalance****************\n")
	fmt.Println("*************错误参数ChargeBalance****************")
	str = host + "/user/charge?uid=%v&money=%v"
	for _, m := range uidSlice {
		vmap, err := util.HGet(str, m, -10)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("*************缺省参数ChargeBalance****************\n")
	fmt.Println("*************缺省参数ChargeBalance****************")
	str = host + "/user/charge?uid=555"
	for i := 0; i < 10; i++ {
		vmap, err := util.HGet(str)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("\n\n\nTest ChargeBalance end\n")
	buff.WriteString("========================================================\n\n\n\n")
	fmt.Println("\n\n\nTest ChargeBalance end")
	fmt.Println("========================================================\n\n\n")
}

func TestWithDraw(t *testing.T) {
	buff.WriteString("\n\n\n========================================================\n")
	buff.WriteString("Test WithDraw begin\n\n\n\n")
	buff.WriteString("*************正常WithDraw****************\n")
	fmt.Println("\n\n\n========================================================")
	fmt.Println("Test WithDraw begin\n\n\n")
	fmt.Println("*************正常WithDraw****************")
	str := host + "/user/withdraw?uid=%v&money=%v"
	for _, m := range uidSlice {
		vmap, err := util.HGet(str, m, 9)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("*************正常WithDraw，余额不足****************\n")
	fmt.Println("*************正常WithDraw，余额不足****************")
	str = host + "/user/withdraw?uid=%v&money=%v"
	for _, m := range uidSlice {
		vmap, err := util.HGet(str, m, 20)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("*************错误参数WithDraw****************\n")
	fmt.Println("*************错误参数WithDraw****************")
	str = host + "/user/withdraw?uid=%v&money=%v"
	for _, m := range uidSlice {
		vmap, err := util.HGet(str, m, -10)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("*************缺省参数WithDraw****************\n")
	fmt.Println("*************缺省参数WithDraw****************")
	str = host + "/user/withdraw?uid=555"
	for i := 0; i < 10; i++ {
		vmap, err := util.HGet(str)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("\n\n\nTest WithDraw end\n")
	buff.WriteString("========================================================\n\n\n\n")
	fmt.Println("\n\n\nTest WithDraw end")
	fmt.Println("========================================================\n\n\n")
}

func TestAuthDriver(t *testing.T) {
	buff.WriteString("\n\n\n========================================================\n")
	buff.WriteString("Test auth begin\n\n\n\n")
	buff.WriteString("*************正常Auth****************\n")
	fmt.Println("\n\n\n========================================================")
	fmt.Println("Test auth begin\n\n\n")
	fmt.Println("*************正常Auth****************")
	str := host + "/user/auth?uid=%v"
	for _, m := range uidSlice {
		vmap, err := util.HGet(str, m)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("*************缺省参数auth****************\n")
	fmt.Println("*************缺省参数auth****************")
	str = host + "/user/auth?mobile=123"
	for i := 0; i < 10; i++ {
		vmap, err := util.HGet(str)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("\n\n\nTest auth end\n")
	buff.WriteString("========================================================\n\n\n\n")
	fmt.Println("\n\n\nTest auth end")
	fmt.Println("========================================================\n\n\n")
}

func TestDeleteUser(t *testing.T) {
	buff.WriteString("\n\n\n========================================================\n")
	buff.WriteString("Test delete begin\n\n\n\n")
	buff.WriteString("*************正常delete****************\n")
	fmt.Println("\n\n\n========================================================")
	fmt.Println("Test delete begin\n\n\n")
	fmt.Println("*************正常delete****************")
	str := host + "/user/delbyid?uid=%v"
	for i := 0; i < 10; i++ {
		vmap, err := util.HGet(str, uidSlice[i])
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	str = host + "/user/delbym?mobile=%v"
	for _, m := range regMobileSlice {
		vmap, err := util.HGet(str, m)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("*************缺省参数delete****************\n")
	fmt.Println("*************缺省参数delete****************")
	str = host + "/user/delbym?head=112"
	for i := 0; i < 10; i++ {
		vmap, err := util.HGet(str)
		writeString(vmap, err)
		fmt.Println(vmap, err)
	}
	buff.WriteString("\n\n\nTest delete end\n")
	buff.WriteString("========================================================\n\n\n\n")
	fmt.Println("\n\n\nTest delete end")
	fmt.Println("========================================================\n\n\n")
}

func TestClose(t *testing.T) {
	Close()
	buff.Close()
}

func writeString(s string, err error) {
	buff.WriteString(s)
	buff.WriteString("\n")
	if err != nil {
		buff.WriteString(err.Error())
		buff.WriteString("\n")
	}

}
