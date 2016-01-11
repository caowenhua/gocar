package db

import (
	"errors"
	"github.com/Centny/gwf/dbutil"
	"me.car/bean"
)

func Register(userName string, mobile string, hobby string, head string, gender int64) (int64, error) {
	user, err := findUserByMobile(mobile)
	if err == nil {
		return user.Uid, errors.New("user is existed")
	}
	id, err := dbutil.DbInsert(Db, "INSERT INTO tb_user (userName,mobile,hobby,head,gender) VALUES (?,?,?,?,?)", userName, mobile, hobby, head, gender)
	PanicErr(err)

	if err != nil {
		return user.Uid, err
	} else {
		user, err := FindUserById(id)
		return user.Uid, err
	}
}

func Login(mobile string) (bean.User, error) {
	user, err := findUserByMobile(mobile)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func FillInfo(userName string, hobby string, head string, gender int64, uid int64) (string, error) {
	flag, err := dbutil.DbUpdate(Db, "UPDATE tb_user SET userName=?,hobby=?,head=?,gender=? WHERE uid=?", userName, hobby, head, gender, uid)
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
	flag, err := dbutil.DbUpdate(Db, "UPDATE tb_user SET isDriver=? WHERE uid=?", true, uid)
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

func DeleteUser(userid int64) (string, error) {
	i, err := dbutil.DbUpdate(Db, "delete from tb_user where uid = ?", userid)
	PanicErr(err)
	if err == nil {
		if i == 0 {
			return "no such user", errors.New("no such user")
		} else {
			return "success", nil
		}
	}
	return "config error", errors.New("config error")
}

func DeleteUserByMobile(mobile string) (string, error) {
	i, err := dbutil.DbUpdate(Db, "delete from tb_user where mobile = ?", mobile)
	PanicErr(err)
	if err == nil {
		if i == 0 {
			return "no such user", errors.New("no such user")
		} else {
			return "success", nil
		}
	}
	return "config error", errors.New("config error")
}

func ChargeBalance(uid int64, money float64) (string, error) {
	flag, err := dbutil.DbUpdate(Db, "UPDATE tb_user SET balance=balance+? WHERE uid=?", money, uid)
	PanicErr(err)
	if err != nil {
		return "database error", err
	} else {
		if flag == 0 {
			return "no such user", errors.New("no such user")
		} else {
			return "success", nil
		}
	}
}

//司机的提现限制
func WithDrawBalance(uid int64, money float64) (string, error) {
	slice := []bean.User{}
	err := dbutil.DbQueryS(Db, &slice, "SELECT * FROM tb_user WHERE uid=?", uid)
	PanicErr(err)
	if err != nil {
		return "database error", err
	} else {
		if len(slice) > 0 {
			if slice[0].Balance < money {
				return "no enough money", errors.New("no enough money")
			} else {
				flag, err := dbutil.DbUpdate(Db, "UPDATE tb_user SET balance=balance-? WHERE uid=?", money, uid)
				PanicErr(err)
				if err != nil {
					return "database error", err
				} else {
					if flag == 0 {
						return "no such user", errors.New("no such user")
					} else {
						return "success", nil
					}
				}
			}
		} else {
			return "no such user", errors.New("no such user")
		}
	}
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

func FindUserById(uid int64) (bean.User, error) {
	user := bean.User{}
	slice := []bean.User{}
	err := dbutil.DbQueryS(Db, &slice, "SELECT * FROM tb_user WHERE uid=?", uid)
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

func IsUserExist(uid int64) error {
	params := []bool{}
	err := dbutil.DbQueryS(Db, &params, "SELECT isDriver FROM tb_user WHERE uid=?", uid)
	if err != nil {
		return err
	} else {
		if len(params) > 0 {
			return nil
		} else {
			return errors.New("no such user")
		}
	}
}

func IsDriverExist(uid int64) error {
	params := []bool{}
	err := dbutil.DbQueryS(Db, &params, "SELECT isDriver FROM tb_user WHERE uid=? and isDriver=true", uid)
	if err != nil {
		return err
	} else {
		if len(params) > 0 {
			return nil
		} else {
			return errors.New("no such user")
		}
	}
}
