package network

import (
	"github.com/Centny/gwf/routing"
	"me.car/db"
)

func ChargeBalance(hs *routing.HTTPSession) routing.HResult {
	var uid int64
	var money float64
	err := hs.ValidCheckVal(`
		uid,R|I,R:0;
		money,R|F,R:0;
		`, &uid, &money)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else {
		s, err := db.ChargeBalance(uid, money)
		if err != nil {
			return hs.MsgResErr2(1, "", err)
		} else {
			return hs.MsgRes(s)
		}
	}
}

func WithDrawBalance(hs *routing.HTTPSession) routing.HResult {
	var uid int64
	var money float64
	err := hs.ValidCheckVal(`
		uid,R|I,R:0;
		money,R|F,R:0;
		`, &uid, &money)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else {
		s, err := db.WithDrawBalance(uid, money)
		if err != nil {
			return hs.MsgResErr2(1, "", err)
		} else {
			return hs.MsgRes(s)
		}
	}
}

func DriverCancelOrder(hs *routing.HTTPSession) routing.HResult {
	var uid int64
	var oid int64
	var drid int64
	err := hs.ValidCheckVal(`
		uid,R|I,R:0;
		oid,R|I,R:0;
		drid,R|I,R:0;
		`, &uid, &oid, &drid)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else {
		s, err := db.DriverCancelOrder(uid, oid, drid)
		if err != nil {
			return hs.MsgResErr2(1, "", err)
		} else {
			return hs.MsgRes(s)
		}
	}
}

func PassengerCancelOrder(hs *routing.HTTPSession) routing.HResult {
	var uid int64
	var oid int64
	var drid int64
	err := hs.ValidCheckVal(`
		uid,R|I,R:0;
		oid,R|I,R:0;
		drid,R|I,R:0;
		`, &uid, &oid, &drid)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else {
		s, err := db.PassengerCancelOrder(uid, oid, drid)
		if err != nil {
			return hs.MsgResErr2(1, "", err)
		} else {
			return hs.MsgRes(s)
		}
	}
}

func GetPassengerOrderList(hs *routing.HTTPSession) routing.HResult {
	var uid int64
	var page int64
	var pageCount int64
	err := hs.ValidCheckVal(`
		uid,R|I,R:0;
		page,R|I,R:0;
		pageCount,R|I,R:0;
		`, &uid, &page, &pageCount)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else {
		s, err := db.GetPassengerOrderList(uid, page, pageCount)
		if err != nil {
			return hs.MsgResErr2(1, "", err)
		} else {
			return hs.MsgRes(s)
		}
	}
}

func GetDriverOrderList(hs *routing.HTTPSession) routing.HResult {
	var uid int64
	var page int64
	var pageCount int64
	err := hs.ValidCheckVal(`
		uid,R|I,R:0;
		page,R|I,R:0;
		pageCount,R|I,R:0;
		`, &uid, &page, &pageCount)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else {
		s, err := db.GetDriverOrderList(uid, page, pageCount)
		if err != nil {
			return hs.MsgResErr2(1, "", err)
		} else {
			return hs.MsgRes(s)
		}
	}
}
