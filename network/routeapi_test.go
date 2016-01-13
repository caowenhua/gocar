package network

import (
	"github.com/Centny/gwf/routing/httptest"
	"github.com/Centny/gwf/util"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

const (
	csLng = 112.898254
	csLat = 23.33469
	ceLng = 112.878254
	ceLat = 23.34769
)

var buff *os.File
var ts *httptest.Server
var uidSlice []int64
var t time
var tmpFinds []tmpFind

func init() {
	ts = new_ts()
	uidSlice = []int64{}
	var outputError error
	buff, outputError = os.OpenFile("routeapi_test.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		panic(outputError)
	}
	t = time.Now()
	tmpFind = []tmpFind{}
}

func TestPrepare(t *testing.T) {
	buff.WriteString("\n\n\n========================================================\n")
	buff.WriteString("TestPrepare start\n\n\n")
	buff.WriteString("*************测试前的准备，新建测试用户，10司机账号+10个乘客账号****************\n\n\n")
	head := "test head"
	str := "/user/reg?userName=%v&mobile=%v&head=%v"
	for i := 0; i < 10; i++ {
		name := "driver" + strconv.Itoa(t.Nanosecond()) + strconv.Itoa(i)
		mobile := "dm" + strconv.Itoa(t.Nanosecond()) + strconv.Itoa(i)
		var vmap map[string]interface{}
		s, err := ts.G(str, name, mobile, head)
		vmap, _ = util.Json2Map(s)
		uidSlice = append(uidSlice, vmap["data"])
		writeString(s, err)
	}

	for i := 0; i < 10; i++ {
		name := "passenger" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		mobile := "pm" + strconv.Itoa(ti.Nanosecond()) + strconv.Itoa(i)
		var vmap map[string]interface{}
		s, err := ts.G(str, name, mobile, head)
		vmap, _ = util.Json2Map(s)
		uidSlice = append(uidSlice, vmap["data"])
		writeString(s, err)
	}
	buff.WriteString("TestPrepare end\n\n\n")
	buff.WriteString("\n\n\n========================================================\n")
}

func TestDriverAddRoute(t *testing.T) {
	buff.WriteString("\n\n\n========================================================\n")
	buff.WriteString("Test DriverAddRoute begin\n\n\n")

	buff.WriteString("\n***********************缺失参数加入**********************\n")
	buff.WriteString("########缺失uid#########")
	str := "/route/dadd?startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, unixTime, "sPlace18", "ePlace18", csLat, csLng, ceLat, ceLng, "003", "002", 8.2)
	writeString(s, err)
	buff.WriteString("########缺失startTime#########")
	str := "/route/dadd?uid=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[8], "sPlace28", "ePlace28", csLat, csLng, ceLat, ceLng, "003", "002", 8.2)
	writeString(s, err)
	buff.WriteString("########缺失sPlace#########")
	str := "/route/dadd?uid=%v&startTime=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[8], unixTime, "ePlace38", csLat, csLng, ceLat, ceLng, "003", "002", 8.2)
	writeString(s, err)
	buff.WriteString("########缺失ePlace#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[8], unixTime, "sPlace38", csLat, csLng, ceLat, ceLng, "003", "002", 8.2)
	writeString(s, err)
	buff.WriteString("########缺失sLat#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[8], unixTime, "sPlace48", "ePlace48", csLng, ceLat, ceLng, "003", "002", 8.2)
	writeString(s, err)
	buff.WriteString("########缺失sLng#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[8], unixTime, "sPlace58", "ePlace58", csLat, ceLat, ceLng, "003", "002", 8.2)
	writeString(s, err)
	buff.WriteString("########缺失eLat#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[8], unixTime, "sPlace68", "ePlace68", csLat, csLng, ceLng, "003", "002", 8.2)
	writeString(s, err)
	buff.WriteString("########缺失eLng#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[8], unixTime, "sPlace78", "ePlace78", csLat, csLng, ceLat, "003", "002", 8.2)
	writeString(s, err)
	buff.WriteString("########缺失sCity#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[8], unixTime, "sPlace88", "ePlace88", csLat, csLng, ceLat, ceLng, "002", 8.2)
	writeString(s, err)
	buff.WriteString("########缺失eCity#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[8], unixTime, "sPlace98", "ePlace98", csLat, csLng, ceLat, ceLng, "003", 8.2)
	writeString(s, err)
	buff.WriteString("########缺失distance#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v"
	s, err := ts.G(str, uidSlice[8], unixTime, "sPlace08", "ePlace08", csLat, csLng, ceLat, ceLng, "003", "002")
	writeString(s, err)

	buff.WriteString("\n***********************错误参数格式加入****************\n")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, "userid", "unixTime", "sPlace8", "ePlace8", "csLat", "csLng", "ceLat", "ceLng", "003", "002", "8.2")
	writeString(s, err)

	buff.WriteString("\n***********************正常参数加入****************\n")
	unixTime := t.Unix()
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	for i := 0; i < 8; i++ {
		f := rand.Float64() * 0.0015
		s, err := ts.G(str, uidSlice[i], unixTime, "sPlace"+strconv.Itoa(i), "ePlace"+strconv.Itoa(i), csLat+f, csLng+f/2, ceLat+f, ceLng+f/2, "001", "002", 8.2)
		writeString(s, err)
	}
	//起点城市不一致
	s, err := ts.G(str, uidSlice[8], unixTime, "sPlace8", "ePlace8", csLat, csLng, ceLat, ceLng, "003", "002", 8.2)
	writeString(s, err)
	//终点城市不一致
	s, err := ts.G(str, uidSlice[9], unixTime, "sPlace9", "ePlace9", csLat, csLng, ceLat, ceLng, "001", "003", 8.2)
	writeString(s, err)

}

func TestPassengerFindDriver(t *testing.T) {
	buff.WriteString("\n\n\n========================================================\n")
	buff.WriteString("Test PassengerFindDriver begin\n\n\n")
	// uid int64, startTime int64, sLat float64,
	// sLng float64, eLat float64, eLng float64, sCity string, eCity string
	buff.WriteString("\n***********************缺失参数寻找**********************\n")
	buff.WriteString("########缺失uid#########")
	str := "/route/dadd?startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, unixTime, csLat, csLng, ceLat, ceLng, "003", "002")
	writeString(s, err)
	buff.WriteString("########缺失startTime#########")
	str := "/route/dadd?uid=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[11], csLat, csLng, ceLat, ceLng, "003", "002")
	writeString(s, err)
	buff.WriteString("########缺失sLat#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[11], unixTime, csLng, ceLat, ceLng, "003", "002")
	writeString(s, err)
	buff.WriteString("########缺失sLng#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[11], unixTime, csLat, ceLat, ceLng, "003", "002")
	writeString(s, err)
	buff.WriteString("########缺失eLat#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLng=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[11], unixTime, csLat, csLng, ceLng, "003", "002")
	writeString(s, err)
	buff.WriteString("########缺失eLng#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&sCity=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[11], unixTime, csLat, csLng, ceLat, "003", "002")
	writeString(s, err)
	buff.WriteString("########缺失sCity#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&eCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[11], unixTime, csLat, csLng, ceLat, ceLng, "002")
	writeString(s, err)
	buff.WriteString("########缺失eCity#########")
	str := "/route/dadd?uid=%v&startTime=%v&sPlace=%v&ePlace=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&distance=%v"
	s, err := ts.G(str, uidSlice[11], unixTime, csLat, csLng, ceLat, ceLng, "003")
	writeString(s, err)

	buff.WriteString("\n***********************错误参数格式寻找****************\n")
	str := "/route/dadd?uid=%v&startTime=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v"
	s, err := ts.G(str, "userid", "unixTime", "csLat", "csLng", "ceLat", "ceLng", "003", "002")
	writeString(s, err)

	buff.WriteString("\n***********************正常参数寻找****************\n")
	unixTime := t.Unix()
	str := "/route/dadd?uid=%v&startTime=%v&sLat=%v&sLng=%v&eLat=%v&eLng=%v&sCity=%v&eCity=%v"
	for i := 10; i < 20; i++ {
		s, err := ts.G(str, uidSlice[i], unixTime, csLat, csLng, ceLat, ceLng, "001", "002")
		writeString(s, err)
		routes := []bean.FindDriverRoute{}
		util.J2Ss(s, routes)
		if len(routes) > 0 {
			tmpFinds = append(tmpFinds, tmpFind{routes[0].Drid, routes[0].Duid})
		} else {
			tmpFinds = append(tmpFinds, tmpFind{-1, -1})
		}
	}
}

func TestPassgerJoinRoute(t *testing.T) {

}

func new_ts() *httptest.Server {
	ts := httptest.NewMuxServer()
	ts.mux.HFunc("/route/dadd", DriverAddRoute)
	ts.mux.HFunc("/route/padd", PassengerJoinRoute)
	ts.mux.HFunc("/route/pfind", PassengerFindDriver)
	return ts
}

func writeString(s string, err error) {
	buff.WriteString(s)
	buff.WriteString("\n")
	if err != nil {
		buff.WriteString(err.Error())
		buff.WriteString("\n")
	}

}

type tmpFind struct {
	drid int64
	duid int64
}
