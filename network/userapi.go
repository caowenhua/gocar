package network

import (
	"github.com/Centny/gwf/routing"
	"me.car/db"
)

func Register(hs *routing.HTTPSession) routing.HResult {
	var userName string
	hobby := ""
	var mobile string
	head := ""
	var gender int64
	gender = 1
	err := hs.ValidCheckVal(`
		userName,R|S,L:0;
		hobby,O|S,L:0;
		mobile,R|S,L:0;
		head,O|S,L:0;
		gender,O|I,R:-1~2;
		`, &userName, &hobby, &mobile, &head, &gender)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else {
		s, err := db.Register(userName, mobile, hobby, head, gender)
		if err != nil {
			return hs.MsgResErr2(1, "", err)
		} else {
			return hs.MsgRes(s)
		}
	}
}

func Login(hs *routing.HTTPSession) routing.HResult {
	var mobile string
	err := hs.ValidCheckVal(`
		mobile,R|S,L:0;
		`, &mobile)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else {
		s, err := db.Login(mobile)
		if err != nil {
			return hs.MsgResErr2(1, "", err)
		} else {
			return hs.MsgRes(s)
		}
	}
}

func FillInfo(hs *routing.HTTPSession) routing.HResult {
	var userName string
	hobby := ""
	var mobile string
	head := ""
	var gender int64
	gender = 1
	var uid int64
	err := hs.ValidCheckVal(`
		userName,R|S,L:0;
		hobby,O|S,L:0;
		mobile,R|S,L:0;
		head,O|S,L:0;
		gender,O|I,R:-1~2;
		uid,R|I,R:0
		`, &userName, &hobby, &mobile, &head, &gender, &uid)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else {
		s, err := db.FillInfo(userName, mobile, hobby, head, gender, uid)
		if err != nil {
			return hs.MsgResErr2(1, "", err)
		} else {
			return hs.MsgRes(s)
		}
	}
}

func AuthDriver(hs *routing.HTTPSession) routing.HResult {
	var uid int64
	err := hs.ValidCheckVal(`
		uid,R|I,R:0
		`, &uid)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else {
		s, err := db.AuthDriver(uid)
		if err != nil {
			return hs.MsgResErr2(1, "", err)
		} else {
			return hs.MsgRes(s)
		}
	}
}

func DeleteUserByMobile(hs *routing.HTTPSession) routing.HResult {
	var mobile string
	err := hs.ValidCheckVal(`
		mobile,R|S,L:0
		`, &mobile)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else {
		s := db.DeleteUserByMobile(mobile)
		return hs.MsgRes(s)
	}
}

func DeleteUserById(hs *routing.HTTPSession) routing.HResult {
	var uid int64
	err := hs.ValidCheckVal(`
		uid,R|I,R:0
		`, &uid)
	if err != nil {
		return hs.MsgResErr(100, "config error", err)
	} else {
		s := db.DeleteUser(uid)
		return hs.MsgRes(s)
	}
}
