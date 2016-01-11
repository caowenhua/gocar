package db

import (
	"errors"
	"github.com/Centny/gwf/dbutil"
)

func DriverCancelOrder(uidDriver int64, oid int64, drid int64) (string, error) {
	err := IsUserExist(uidDriver)
	if err != nil {
		return "", err
	}

	var txerr error
	txerr = nil

	tx, err := Db.Begin()
	recordTxErr(&txerr, &err)

	//司机订单置为无效
	_, err = dbutil.DbUpdate2(tx, "UPDATE tb_order SET isEnable=? WHERE duid=? and oid=? and isEnable=true", false, uidDriver, oid)
	recordTxErr(&txerr, &err)

	//司机路线置为无效
	_, err = dbutil.DbUpdate2(tx, "UPDATE tb_driver_route SET isEnable=? WHERE drid=?", false, drid)
	recordTxErr(&txerr, &err)

	//找出受到影响的乘客
	effectedUser := []tmpOrderBean{}
	err = dbutil.DbQueryS2(tx, &effectedUser, "SELECT uid, price FROM tb_order WHERE drid =? and uid !=?", drid, uidDriver)
	recordTxErr(&txerr, &err)

	if len(effectedUser) > 0 {
		//司机减钱
		_, err := dbutil.DbUpdate2(tx, "UPDATE tb_user SET balance=balance-? WHERE uid=?", effectedUser[0].price*float64(len(effectedUser)), uidDriver)
		recordTxErr(&txerr, &err)

		for _, bean := range effectedUser {
			//推送给用户告知订单被司机取消
			//乘客订单置为无效并退回钱
			_, err = dbutil.DbUpdate2(tx, "UPDATE tb_order SET isEnable=? WHERE uid=? and oid=?", false, bean.uid, oid)
			recordTxErr(&txerr, &err)
			_, err = dbutil.DbUpdate2(tx, "UPDATE tb_user SET balance=balance+? WHERE uid=?", bean.price, bean.uid)
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

	tx, err := Db.Begin()
	recordTxErr(&txerr, &err)
	//乘客订单置为无效
	_, err = dbutil.DbUpdate2(tx, "UPDATE tb_order SET isEnable=? WHERE uid=? and oid=? and isEnable=true", false, uidPassenger, oid)
	recordTxErr(&txerr, &err)
	//找出订单的价钱
	prices := []float64{}
	err = dbutil.DbQueryS2(tx, &prices, "SELECT price FROM tb_order WHERE oid =?", oid)
	if len(prices) > 0 {
		//乘客加回钱
		_, err := dbutil.DbUpdate2(tx, "UPDATE tb_user SET balance=balance+? WHERE uid=?", prices[0], uidPassenger)
		recordTxErr(&txerr, &err)
		//找出司机
		uids := []int64{}
		err = dbutil.DbQueryS2(tx, &uids, "SELECT duid FROM tb_driver_route WHERE drid =?", drid)
		recordTxErr(&txerr, &err)
		//司机减钱
		if len(uids) > 0 {
			_, err := dbutil.DbUpdate2(tx, "UPDATE tb_user SET balance=balance-? WHERE uid=?", prices[0], uids[0])
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

func GetPassengerOrderList(uid int64, page int64, pageCount int64) {

}

func GetDriverOrderList(uid int64, page int64, pageCount int64) {

}

type tmpOrderBean struct {
	uid   int64
	price float64
}
