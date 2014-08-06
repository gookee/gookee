package db

import (
	"database/sql"
	"fmt"
	_ "sqlite3"
)

var DbTypeName string
var DbConnectionString string

func Init(dbConnectionString string) {
	DbTypeName = "sqlite3"
	DbConnectionString = dbConnectionString
}

func ExecBackId(sqlStr string, values ...interface{}) int64 {
	SqlDb, e := sql.Open(DbTypeName, DbConnectionString)
	checkErr(e)
	defer SqlDb.Close()
	stmt, err := SqlDb.Prepare(sqlStr)
	checkErr(err)
	defer stmt.Close()
	res, errres := stmt.Exec(values...)
	checkErr(errres)
	id, errid := res.LastInsertId()
	checkErr(errid)
	return id
}

func Exec(sqlStr string, values ...interface{}) int64 {
	SqlDb, e := sql.Open(DbTypeName, DbConnectionString)
	checkErr(e)
	defer SqlDb.Close()
	stmt, err := SqlDb.Prepare(sqlStr)
	checkErr(err)
	defer stmt.Close()
	res, errres := stmt.Exec(values...)
	checkErr(errres)
	affect, erraff := res.RowsAffected()
	checkErr(erraff)
	return affect
}

func ExecBatch(sqlStr string) int64 {
	SqlDb, e := sql.Open(DbTypeName, DbConnectionString)
	checkErr(e)
	defer SqlDb.Close()
	res, err := SqlDb.Exec(sqlStr)
	checkErr(err)
	affect, erraff := res.RowsAffected()
	checkErr(erraff)
	return affect
}

func ExecDt(sqlStr string, values ...interface{}) DataTable {
	SqlDb, e := sql.Open(DbTypeName, DbConnectionString)
	checkErr(e)
	defer SqlDb.Close()
	rows, errrow := SqlDb.Query(sqlStr, values...)
	checkErr(errrow)
	defer rows.Close()
	dt := NewDataTable()
	columnNames, _ := rows.Columns()
	fields := make([]interface{}, 0)
	dt.AddColumnByNames(columnNames...)
	for i := 0; i < len(columnNames); i++ {
		var obj interface{}
		fields = append(fields, &obj)
	}
	for rows.Next() {
		dr := dt.NewRow()
		err := rows.Scan(fields...)
		checkErr(err)
		for i := 0; i < len(dt.Columns); i++ {
			dr.Values[dt.Columns[i].Name] = *(fields[i].(*interface{}))
		}
		dt.AddRow(dr)
	}
	return dt
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
