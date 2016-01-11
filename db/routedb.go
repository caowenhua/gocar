package db

import (
	"database/sql"
	"errors"
	"github.com/Centny/gwf/dbutil"
	"me.car/bean"
)

const (
	UnitPrice = 1.1
	latValue  = 0.009788
	lngValue  = 0.009
)

func DriverAddRoute(uid int64, startTime int64, sPlace string, ePlace string, sLat float64,
	sLng float64, eLat float64, eLng float64, sCity string, eCity string, distance float64) (string, error) {
	err := checkStartTime(startTime, uid)
	if err != nil {
		return "", err
	}
	rid, err := insertRoute(startTime, sPlace, ePlace, sLat, sLng, eLat, eLng, sCity, eCity)
	if rid == 0 {
		return "", err
	}

	drid, err := dbutil.DbInsert(Db, "INSERT INTO tb_driver_route (uid,rid,unitPrice,distance,price) VALUES (?,?,?,?,?)",
		uid, rid, UnitPrice, distance, UnitPrice*distance)
	PanicErr(err)
	if drid == 0 {
		return "", errors.New("insert driver route error")
	}
	_, err = dbutil.DbInsert(Db, "INSERT INTO tb_trip (uid,drid) VALUES (?,?)", uid, drid)
	if err != nil {
		return "", err
	} else {
		return "success", nil
	}
}

func FindDriverRoute(uid int64, startTime int64, sLat float64,
	sLng float64, eLat float64, eLng float64, sCity string, eCity string) (interface{}, error) {
	err := checkStartTime(startTime, uid)
	if err != nil {
		return "", err
	}
	routes := []bean.FindDriverRoute{}
	err = dbutil.DbQueryS(Db, &routes,
		"SELECT startTime,sPlace,ePlace,sLat,sLng,eLat,eLng,sCity,eCity,price,drid "+
			"FROM tb_route NATURAL JOIN tb_driver_route WHERE"+
			"sCity=? and eCity=? and ABS(sLat-?)<=? and ABS(sLng-?)<=? and ABS(eLat-?)<=? and ABS(eLng-?)<=? and ABS(startTime-?)<1800",
		sCity, eCity, sLat, latValue, sLng, lngValue, eLat, latValue, eLng, lngValue, startTime)
	if err == nil {
		return routes, nil
	} else {
		return "", err
	}

}

func checkStartTime(startTime int64, uid int64) error {
	startTimes := []int64{}
	err := dbutil.DbQueryS(Db, &startTimes, "SELECT startTime FROM tb_driver_route NATURAL JOIN tb_route WHERE uid=?", uid)
	PanicErr(err)
	for _, t := range startTimes {
		if startTime-t < 1800 || t-startTime < 1800 {
			return errors.New("invalid startTime")
		}
	}
	return nil
}

func insertRoute(startTime int64, sPlace string, ePlace string, sLat float64,
	sLng float64, eLat float64, eLng float64, sCity string, eCity string) (int64, error) {
	rid, err := dbutil.DbInsert(Db,
		"INSERT INTO tb_route (startTime,sPlace,ePlace,sLat,sLng,eLat,eLng,sCity,eCity) VALUES (?,?,?,?,?,?,?,?,?)",
		startTime, sPlace, ePlace, sLat, sLng, eLat, eLng, sCity, eCity)
	PanicErr(err)
	if rid == 0 {
		return rid, errors.New("insert route error")
	} else {
		return rid, nil
	}
}

func insertRouteWithTx(tx *sql.Tx, startTime int64, sPlace string, ePlace string, sLat float64,
	sLng float64, eLat float64, eLng float64, sCity string, eCity string) (int64, error) {
	rid, err := dbutil.DbInsert2(tx,
		"INSERT INTO tb_route (startTime,sPlace,ePlace,sLat,sLng,eLat,eLng,sCity,eCity) VALUES (?,?,?,?,?,?,?,?,?)",
		startTime, sPlace, ePlace, sLat, sLng, eLat, eLng, sCity, eCity)
	PanicErr(err)
	if rid == 0 {
		return rid, errors.New("insert route error")
	} else {
		return rid, nil
	}
}

func PassengerJoinRoute(uid int64, startTime int64, sPlace string, ePlace string, sLat float64,
	sLng float64, eLat float64, eLng float64, sCity string, eCity string, price float64, drid int64) (string, error) {
	err := checkStartTime(startTime, uid)
	if err != nil {
		return "", err
	}

	balance := []float64{}
	err = dbutil.DbQueryS(Db, &balance, "SELECT balance FROM tb_user WHERE uid=?", uid)
	if len(balance) == 0 {
		return "", errors.New("no such user")
	} else {
		if balance[0] < price {
			return "", errors.New("your balance is not enough, please withdraw first")
		}
	}

	var txerr error
	txerr = nil

	tx, err := Db.Begin()
	recordTxErr(&txerr, &err)

	rid, err := insertRouteWithTx(tx, startTime, sPlace, ePlace, sLat, sLng, eLat, eLng, sCity, eCity)
	recordTxErr(&txerr, &err)

	_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_passenger_route (uid,rid,price) VALUES (?,?,?)",
		uid, rid, price)
	recordTxErr(&txerr, &err)

	_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_trip (uid,drid) VALUES (?,?)", uid, drid)
	recordTxErr(&txerr, &err)

	code, err := dbutil.DbUpdate2(tx, "UPDATE tb_user SET balance=balance-? WHERE uid=?", price, uid)
	recordTxErr(&txerr, &err)

	if txerr != nil || code == 0 {
		tx.Rollback()
		return "", errors.New("database passenger join error")
	} else {
		tx.Commit()
		return "success", nil
	}
}

func recordTxErr(txerr *error, runningErr *error) {
	if *runningErr != nil {
		txerr = runningErr
	}
}
