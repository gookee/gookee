package dal

import (
	"db"
)

type managerdal struct{}

var Managerdal *managerdal = &managerdal{}

func (this *managerdal) IsExist(username, password string) bool {
	drs := db.ExecDt("select * from manager where username = ? and password = ?", username, password).Rows
	if len(drs) == 0 {
		return false
	} else {
		return true
	}
}

func (this *managerdal) Update(username, password string) int64 {
	return db.Exec("update manager set password=? where username = ?", password, username)
}
