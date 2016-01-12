package db

import (
	"errors"
	"fmt"
	"github.com/Centny/gwf/dbutil"
	"me.car/bean"
	"strconv"
	"time"
)

const (
	UnitPrice = 1.1
	latValue  = 0.009788
	lngValue  = 0.009
)

func DriverAddRoute(uid int64, startTime int64, sPlace string, ePlace string, sLat float64,
	sLng float64, eLat float64, eLng float64, sCity string, eCity string, distance float64) (string, error) {
	//检查用户是否存在，时间是否已经过去，插入到route
	err := IsUserExist(uid)
	if err != nil {
		return "", err
	}
	err = checkStartTime(startTime, uid)
	if err != nil {
		return "", err
	}

	var txerr error
	txerr = nil

	tx, err := Db.Begin()
	recordTxErr(&txerr, &err)
	//插入到route 得 rid
	rid, err := dbutil.DbInsert2(tx,
		"INSERT INTO tb_route (startTime,sPlace,ePlace,sLat,sLng,eLat,eLng,sCity,eCity) VALUES (?,?,?,?,?,?,?,?,?)",
		startTime, sPlace, ePlace, sLat, sLng, eLat, eLng, sCity, eCity)
	recordTxErr(&txerr, &err)

	//单独插入到driver route得到drid
	drid, err := dbutil.DbInsert2(tx, "INSERT INTO tb_driver_route (duid,rid,unitPrice,distance) VALUES (?,?,?,?)",
		uid, rid, UnitPrice, distance)
	recordTxErr(&txerr, &err)

	//drid生成订单插入order 得oid
	t := time.Now()
	unixTime := t.Unix()
	orderNo := fmt.Sprintf("%02d%2d%02d%4d%2d%2d\n", t.Month(), t.Minute(), t.Day(), t.Year(), t.Second(), t.Hour())
	orderNo = orderNo + strconv.FormatInt(startTime, 10) + strconv.FormatInt(uid, 10) + strconv.FormatInt(drid, 10)
	orderComment := "create route: \n" + sPlace + "   to  " + ePlace
	oid, err := dbutil.DbInsert2(tx, "INSERT INTO tb_order (uid,price,orderTime,orderNo,comment) VALUES (?,?,?,?,?)", uid, UnitPrice*distance, unixTime, orderNo, orderComment)
	recordTxErr(&txerr, &err)

	//插入到route order
	_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_order_route (uid,duid,drid,oid) VALUES (?,?,?,?)", uid, uid, drid, oid)
	recordTxErr(&txerr, &err)

	//创建交易记录
	_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_deal (uid,oid,dealTime,dealPrice) VALUES (?,?,?,?)", uid, oid, unixTime, 0.0)
	recordTxErr(&txerr, &err)
	if txerr != nil {
		tx.Rollback()
		return "", errors.New("database DriverAddRoute error")
	} else {
		tx.Commit()
		return "success", nil
	}
}

func FindDriverRoute(uid int64, startTime int64, sLat float64,
	sLng float64, eLat float64, eLng float64, sCity string, eCity string) (interface{}, error) {
	//先判断用户是否存在，时间是否正确
	err := IsUserExist(uid)
	if err != nil {
		return "", err
	}
	//匹配附近的司机路线
	routes := []bean.FindDriverRoute{}
	err = dbutil.DbQueryS(Db, &routes,
		"SELECT startTime,sPlace,ePlace,sLat,sLng,eLat,eLng,sCity,eCity,price,drid,duid "+
			"FROM tb_route NATURAL JOIN tb_driver_route WHERE "+
			"sCity=? and eCity=? and ABS(sLat-?)<=? and ABS(sLng-?)<=? and ABS(eLat-?)<=? and ABS(eLng-?)<=? and ABS(startTime-?)<1800 and isEnable=true",
		sCity, eCity, sLat, latValue, sLng, lngValue, eLat, latValue, eLng, lngValue, startTime)
	if err == nil {
		return routes, nil
	} else {
		return "", err
	}

}

func checkStartTime(startTime int64, driverUid int64) error {
	startTimes := []int64{}
	err := dbutil.DbQueryS(Db, &startTimes, "SELECT startTime FROM tb_driver_route NATURAL JOIN tb_route WHERE duid=?", driverUid)
	PanicErr(err)
	for _, t := range startTimes {
		if startTime-t < 1800 || t-startTime < 1800 {
			return errors.New("invalid startTime")
		}
	}
	return nil
}

func PassengerJoinRoute(uid int64, startTime int64, sPlace string, ePlace string, sLat float64,
	sLng float64, eLat float64, eLng float64, sCity string, eCity string, drid int64, duid int64) (string, error) {
	//先判断用户是否存在，时间是否正确
	err := IsUserExist(uid)
	if err != nil {
		return "", err
	}
	err = IsDriverExist(duid)
	if err != nil {
		return "", err
	}
	err = checkStartTime(startTime, uid)
	if err != nil {
		return "", err
	}

	//用drid找到司机的uid匹配
	duids := []int64{}
	err = dbutil.DbQueryS(Db, &duids, "SELECT duid FROM tb_driver_route WHERE drid=?", drid)
	if err != nil || len(duids) == 0 {
		return "", errors.New("no such driver")
	} else {
		if duids[0] == duid {
			return "", errors.New("invalid driver !!!")
		}
	}
	//从数据库得到价钱
	prices := []float64{}
	err = dbutil.DbQueryS(Db, &duids, "SELECT unitPrice*distance FROM tb_driver_route WHERE drid=?", drid)
	if err != nil || len(prices) == 0 {
		return "", errors.New("no such route")
	}
	price := prices[0]

	//不够钱提示
	balance := []float64{}
	err = dbutil.DbQueryS(Db, &balance, "SELECT balance FROM tb_user WHERE uid=?", uid)
	if len(balance) == 0 {
		return "", errors.New("no such user")
	} else {
		if balance[0] < price {
			return "", errors.New("your balance is not enough, please charge first")
		}
	}

	var txerr error
	txerr = nil
	t := time.Now()
	unixTime := t.Unix()
	tx, err := Db.Begin()
	recordTxErr(&txerr, &err)
	//插入到route
	rid, err := dbutil.DbInsert2(tx,
		"INSERT INTO tb_route (startTime,sPlace,ePlace,sLat,sLng,eLat,eLng,sCity,eCity) VALUES (?,?,?,?,?,?,?,?,?)",
		startTime, sPlace, ePlace, sLat, sLng, eLat, eLng, sCity, eCity)
	recordTxErr(&txerr, &err)
	//插入到passenger route
	_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_passenger_route (uid,rid,price) VALUES (?,?,?)",
		uid, rid, price)
	recordTxErr(&txerr, &err)
	//插入order，生成订单，得oid
	orderNo := fmt.Sprintf("%02d%2d%02d%4d%2d%2d\n", t.Month(), t.Minute(), t.Day(), t.Year(), t.Second(), t.Hour())
	orderNo = orderNo + strconv.FormatInt(startTime, 10) + strconv.FormatInt(uid, 10) + strconv.FormatInt(drid, 10)
	orderComment := "join route: \n" + sPlace + "   to  " + ePlace
	oid, err := dbutil.DbInsert2(tx, "INSERT INTO tb_order (uid,price,orderTime,orderNo,comment) VALUES (?,?,?,?,?)", uid, price, unixTime, orderNo, orderComment)
	recordTxErr(&txerr, &err)
	//插入到route order
	_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_order_route (uid,duid,drid,oid) VALUES (?,?,?,?)", uid, duid, drid, oid)
	recordTxErr(&txerr, &err)
	//扣除金额
	_, err = dbutil.DbUpdate2(tx, "UPDATE tb_user SET balance=balance-? WHERE uid=?", price, uid)
	recordTxErr(&txerr, &err)
	//创建交易记录
	_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_deal (uid,oid,dealTime,dealPrice) VALUES (?,?,?,?)", uid, oid, unixTime, -price)
	recordTxErr(&txerr, &err)
	//司机余额增加
	_, err = dbutil.DbUpdate2(tx, "UPDATE tb_user SET balance=balance+? WHERE uid=?", price, duid)
	recordTxErr(&txerr, &err)
	//创建交易记录
	_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_deal (uid,oid,dealTime,dealPrice) VALUES (?,?,?,?)", duid, oid, unixTime, price)
	recordTxErr(&txerr, &err)

	if txerr != nil {
		tx.Rollback()
		return "", errors.New("database passenger join error")
	} else {
		tx.Commit()
		return "success", nil
	}
}
