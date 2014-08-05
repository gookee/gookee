package db

import (
	"utility"
)

type DataRow struct {
	Values map[string]interface{}
	Table  *DataTable
}

func (dr DataRow) ToJson() string {
	cols := dr.Table.Columns
	str := "{"
	for i := 0; i < len(cols); i++ {
		if i > 0 {
			str += ","
		}

		name := utility.JsonEncode(cols[i].Name)
		value := utility.JsonEncode(utility.ToStr(dr.Values[cols[i].Name]))
		str += "\"" + name + "\":\"" + value + "\""
	}
	str += "}"

	return str
}
