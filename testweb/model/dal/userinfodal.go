package dal

import (
	"db"
	"utility"
)

type userinfodal struct{}

var Userinfodal *userinfodal = &userinfodal{}

func (this *userinfodal) GetAll(f_username string, pid int) string {
	dt := db.ExecDt("select * from userinfo where username like ? and pid = ?", "%"+f_username+"%", pid)
	return dt.ToJson()
}

func (this *userinfodal) GetRow(id int) string {
	drs := db.ExecDt("select * from userinfo where id = ?", id).Rows
	if len(drs) == 0 {
		return ""
	} else {
		return drs[0].ToJson()
	}
}

func (this *userinfodal) Del(id int) int64 {
	return db.Exec("delete from userinfo where id = ?", id)
}

func (this *userinfodal) DelAll(ids string) int64 {
	return db.Exec("delete from userinfo where id in (" + ids + ")")
}

func (this *userinfodal) Insert(username, address, phone string, isspend, pid int) int64 {
	if pid >= 0 && this.CheckUserNum(pid) < 20 {
		return db.Exec("insert into userinfo(username,address,phone,isspend,pid) values(?,?,?,?,?)", username, address, phone, isspend, pid)
	} else {
		return 0
	}
}

func (this *userinfodal) CheckUserNum(pid int) int {
	return utility.ToInt(db.ExecDt("select count(0) as c from userinfo where pid = ?", pid).Rows[0].Values["c"])
}

func (this *userinfodal) Update(username, address, phone string, isspend, pid, id int) int64 {
	return db.Exec("update userinfo set username=?,address=?,phone=?,isspend=?,pid=? where id = ?", username, address, phone, isspend, pid, id)
}
