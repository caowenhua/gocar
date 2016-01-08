package db

import (
	"errors"
	"github.com/Centny/gwf/dbutil"
	"me.car/bean"
)

func Register(userName string, mobile string, hobby string, head string, gender int64) (string, error) {
	user, err := findUserByMobile(mobile)
	if err != nil {
		return "", err
	} else if user.Uid != 0 {
		return "", errors.New("user is existed")
	}
	_, err = dbutil.DbInsert(Db, "INSERT INTO tb_user (userName,mobile,hobby,head,gender) VALUES (?,?,?,?,?)", userName, mobile, hobby, head, gender)
	PanicErr(err)

	if err != nil {
		return "", err
	} else {
		return "success", nil
	}
}

func FillInfo(userName string, mobile string, hobby string, head string, gender int64, uid int64) (string, error) {
	flag, err := dbutil.DbUpdate(Db, "UPDATE tb_user SET userName=?,mobile=?,hobby=?,head=?,gender=? WHERE uid=?", userName, mobile, hobby, head, gender, uid)
	PanicErr(err)
	if err != nil {
		return "database error", err
	} else {
		if flag == 0 {
			return "no such user", nil
		} else {
			return "success", nil
		}
	}
}

func AuthDriver(uid int64) (string, error) {
	flag, err := dbutil.DbUpdate(Db, "UPDATE tb_user SET isDriver=? WHERE uid=?", true)
	PanicErr(err)
	if err != nil {
		return "database error", err
	} else {
		if flag == 0 {
			return "no such user", nil
		} else {
			return "success", nil
		}
	}
}

func Login(mobile string) (bean.User, error) {
	user, err := findUserByMobile(mobile)
	if err != nil {
		return user, err
	} else if user.Uid != 0 {
		return user, errors.New("user is not existed")
	} else {
		return user, nil
	}
}

func DeleteUser(userid int64) string {
	i, err := dbutil.DbUpdate(Db, "delete from tb_user where userid = ?", userid)
	PanicErr(err)
	if err == nil {
		if i == 0 {
			return "no such user"
		}
		d, e := dbutil.DbUpdate(Db, "delete from tb_account where userid = ?", userid)
		PanicErr(e)
		if d == 0 {
			return "no such user"
		} else {
			return "success"
		}
	}
	return "config error"
}

func DeleteUserByMobile(mobile string) string {
	i, err := dbutil.DbUpdate(Db, "delete from tb_user where mobile = ?", mobile)
	PanicErr(err)
	if err == nil {
		if i == 0 {
			return "no such user"
		}
		d, e := dbutil.DbUpdate(Db, "delete from tb_account where mobile = ?", mobile)
		PanicErr(e)
		if d == 0 {
			return "no such user"
		} else {
			return "success"
		}
	}
	return "config error"
}

func findUserByMobile(mobile string) (bean.User, error) {
	user := bean.User{}
	slice := []bean.User{}
	err := dbutil.DbQueryS(Db, &slice, "SELECT * FROM tb_user WHERE mobile=?", mobile)
	PanicErr(err)
	if err != nil {
		return user, err
	} else {
		if len(slice) > 0 {
			return slice[0], nil
		} else {
			return user, errors.New("user is not existed")
		}
	}
}
