package db

import (
	"errors"
	"fmt"
	"github.com/Centny/gwf/dbutil"
	"me.car/bean"
	"strconv"
	"time"
)

func ChargeBalance(uid int64, money float64) (string, error) {
	//检查用户是否存在
	err := IsUserExist(uid)
	if err != nil {
		return "", err
	}
	var txerr error
	txerr = nil
	t := time.Now()

	tx, err := Db.Begin()
	recordTxErr(&txerr, &err)
	_, err = dbutil.DbUpdate(Db, "UPDATE tb_user SET balance=balance+? WHERE uid=?", money, uid)
	recordTxErr(&txerr, &err)
	unixTime := t.Unix()
	//插入order，生成订单，得oid
	orderNo := fmt.Sprintf("%02d%2d%02d%4d%2d%2d\n", t.Month(), t.Minute(), t.Day(), t.Year(), t.Second(), t.Hour())
	orderNo = orderNo + strconv.FormatInt(unixTime, 10) + strconv.FormatInt(uid, 10) + strconv.FormatInt(0, 10)
	oid, err := dbutil.DbInsert2(tx, "INSERT INTO tb_order (uid,price,orderTime,orderNo) VALUES (?,?,?,?)", uid, money, unixTime, orderNo)
	recordTxErr(&txerr, &err)
	//创建交易记录
	_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_deal (uid,oid,dealTime,dealPrice) VALUES (?,?,?)", uid, oid, unixTime, money)
	recordTxErr(&txerr, &err)
	if txerr != nil {
		tx.Rollback()
		return "", errors.New("ChargeBalance error")
	} else {
		tx.Commit()
		return "success", nil
	}
}

//司机提现限制小于余额-冻结金额，订单未完成，金额冻结
func WithDrawBalance(uid int64, money float64) (string, error) {
	//检查用户是否存在
	err := IsUserExist(uid)
	if err != nil {
		return "", err
	}
	var txerr error
	txerr = nil
	t := time.Now()
	unixTime := t.Unix()
	//查出所有进行的订单的总额
	priceSlice := []float64{}
	err = dbutil.DbQueryS(Db, &priceSlice, "SELECT sum(price) FROM tb_order JOIN tb_route_order ON tb_order.oid = tb_route_order.oid WHERE uid=? and duid !=?, orderTime>?",
		uid, uid, unixTime)
	if err != nil || len(priceSlice) == 0 {
		return "", err
	}

	icePrice := priceSlice[0]

	tx, err := Db.Begin()
	recordTxErr(&txerr, &err)
	slice := []float64{}
	err = dbutil.DbQueryS2(tx, &slice, "SELECT * FROM tb_user WHERE uid=?", uid)
	recordTxErr(&txerr, &err)
	if len(slice) > 0 {
		if slice[0]-icePrice < money {
			return "no enough money", errors.New("no enough money to withdraw or some order had not finished")
		} else {
			_, err := dbutil.DbUpdate2(tx, "UPDATE tb_user SET balance=balance-? WHERE uid=?", money, uid)
			recordTxErr(&txerr, &err)
			if txerr != nil {
				tx.Rollback()
				return "", txerr
			} else {
				tx.Commit()
				return "success", nil
			}
		}
	} else {
		tx.Rollback()
		return "", errors.New("WithDrawBalance error")
	}
}

func DriverCancelOrder(uidDriver int64, oid int64, drid int64) (string, error) {
	err := IsUserExist(uidDriver)
	if err != nil {
		return "", err
	}

	var txerr error
	txerr = nil
	t := time.Now()
	unixTime := t.Unix()

	tx, err := Db.Begin()
	recordTxErr(&txerr, &err)

	//司机订单置为无效
	_, err = dbutil.DbUpdate2(tx, "UPDATE tb_route_order SET isEnable=? WHERE duid=? and oid=? and isEnable=true", false, uidDriver, oid)
	recordTxErr(&txerr, &err)

	//司机路线置为无效
	_, err = dbutil.DbUpdate2(tx, "UPDATE tb_driver_route SET isEnable=? WHERE drid=?", false, drid)
	recordTxErr(&txerr, &err)

	//找出受到影响的乘客
	effectedUser := []tmpOrderBean{}
	err = dbutil.DbQueryS2(tx, &effectedUser, "SELECT uid, price,oid FROM tb_order NATURAL JOIN tb_route_order WHERE drid =? and uid !=?", drid, uidDriver)
	recordTxErr(&txerr, &err)

	if len(effectedUser) > 0 {
		//司机减钱
		_, err := dbutil.DbUpdate2(tx, "UPDATE tb_user SET balance=balance-? WHERE uid=?", effectedUser[0].price*float64(len(effectedUser)), uidDriver)
		recordTxErr(&txerr, &err)
		_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_deal (uid,oid,dealTime,dealPrice) VALUES (?,?,?)", uidDriver, oid, unixTime, -effectedUser[0].price*float64(len(effectedUser)))
		recordTxErr(&txerr, &err)

		for _, bean := range effectedUser {
			//推送给用户告知订单被司机取消
			//乘客订单置为无效并退回钱
			_, err = dbutil.DbUpdate2(tx, "UPDATE tb_route_order SET isEnable=? WHERE uid=? and oid=?", false, bean.uid, oid)
			recordTxErr(&txerr, &err)
			_, err = dbutil.DbUpdate2(tx, "UPDATE tb_user SET balance=balance+? WHERE uid=?", bean.price, bean.uid)
			recordTxErr(&txerr, &err)
			//创建交易记录
			_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_deal (uid,oid,dealTime,dealPrice) VALUES (?,?,?)", bean.uid, bean.oid, unixTime, bean.price)
			recordTxErr(&txerr, &err)
		}
	}

	if txerr != nil {
		tx.Rollback()
		return "", errors.New("database DriverCancelRoute error")
	} else {
		tx.Commit()
		return "success", nil
	}
}

func PassengerCancelOrder(uidPassenger int64, oid int64, drid int64) (string, error) {
	err := IsUserExist(uidPassenger)
	if err != nil {
		return "", err
	}

	var txerr error
	txerr = nil
	t := time.Now()
	unixTime := t.Unix()

	tx, err := Db.Begin()
	recordTxErr(&txerr, &err)
	//乘客订单置为无效
	_, err = dbutil.DbUpdate2(tx, "UPDATE tb_route_order SET isEnable=? WHERE uid=? and oid=? and isEnable=true", false, uidPassenger, oid)
	recordTxErr(&txerr, &err)
	//找出订单的价钱
	prices := []float64{}
	err = dbutil.DbQueryS2(tx, &prices, "SELECT price FROM tb_order NATURAL JOIN tb_route_order WHERE oid =?", oid)
	if len(prices) > 0 {
		//乘客加回钱
		_, err := dbutil.DbUpdate2(tx, "UPDATE tb_user SET balance=balance+? WHERE uid=?", prices[0], uidPassenger)
		recordTxErr(&txerr, &err)
		//创建交易记录
		_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_deal (uid,oid,dealTime,dealPrice) VALUES (?,?,?)", uidPassenger, oid, unixTime, prices[0])
		recordTxErr(&txerr, &err)
		//找出司机
		tmps := []tmpOrderBean{}
		err = dbutil.DbQueryS2(tx, &tmps, "SELECT uid, price,oid FROM tb_order NATURAL JOIN tb_route_order WHERE drid =?", drid)
		recordTxErr(&txerr, &err)
		//司机减钱
		if len(tmps) > 0 {
			_, err := dbutil.DbUpdate2(tx, "UPDATE tb_user SET balance=balance-? WHERE uid=?", prices[0], tmps[0].uid)
			recordTxErr(&txerr, &err)
			//创建交易记录
			_, err = dbutil.DbInsert2(tx, "INSERT INTO tb_deal (uid,oid,dealTime,dealPrice) VALUES (?,?,?)", tmps[0].uid, tmps[0].oid, unixTime, -prices[0])
			recordTxErr(&txerr, &err)
		} else {
			e := errors.New("can not find the driver")
			recordTxErr(&txerr, &e)
		}
	}

	if txerr != nil {
		tx.Rollback()
		return "", errors.New("database PassengerCancelRoute error")
	} else {
		tx.Commit()
		return "success", nil
	}
}

func GetPassengerOrderList(uid int64, page int64, pageCount int64) ([]bean.PassengerOrderResponse, error) {
	response := []bean.PassengerOrderResponse{}
	err := IsUserExist(uid)
	if err != nil {
		return response, err
	}
	err = dbutil.DbQueryS(Db, &response,
		"SELECT oid,price,orderTime,orderNo,drid,isEnable,duid, userName,head,mobile,gender "+
			"FROM (tb_order NATURAL JOIN tb_route_order) JOIN tb_user ON tb_order.duid = tb_user.uid "+
			"WHERE tb_order.duid = ? LIMIT ?,?",
		uid, (page-1)*pageCount, pageCount)
	return response, err
}

// func GetDriverOrderList(uid int64, page int64, pageCount int64) ([]bean.DriverOrderResponse, error) {
// 	response := []bean.DriverOrderResponse{}
// 	err := IsUserExist(uid)
// 	if err != nil {
// 		return response, err
// 	}
// 	// err = dbutil.DbQueryS(Db, &response,
// 	// 	"SELECT oid,price,orderTime,orderNo,drid,isEnable,duid, userName,head,mobile,gender "+
// 	// 		"FROM tb_order JOIN tb_user ON tb_order.duid = tb_user.uid "+
// 	// 		"WHERE tb_order.duid = ? LIMIT ?,?",
// 	// 	uid, (page-1)*pageCount, pageCount)

// }

type tmpOrderBean struct {
	uid   int64
	price float64
	oid   int64
}
