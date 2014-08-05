package db

import (
	"regexp"
	"strings"
	"utility"
)

type DataTable struct {
	Rows    []DataRow
	Columns []DataColumn
}

func NewDataTable() DataTable {
	return DataTable{make([]DataRow, 0), make([]DataColumn, 0)}
}

func JsonToDataTable(str string) DataTable {
	dt := DataTable{}
	drP := regexp.MustCompile("\\{.*?[^\\\\:]('|\\\")\\}")
	drs := drP.FindAllStringSubmatch(str, -1)
	if drs != nil {
		for i := 0; i < len(drs); i++ {
			p := regexp.MustCompile("([^{:,]+?):\\\"(|.*?[^\\\\])\\\"")
			m := p.FindAllStringSubmatch(drs[i][0], -1)
			dr := dt.NewRow()
			for j := 0; j < len(m); j++ {
				name := utility.JsonDecode(strings.Trim(m[j][1], "\""))
				value := utility.JsonDecode(m[j][2])

				if i == 0 {
					dt.AddColumnByNames(name)
				}

				dr.Values[name] = value
			}
			dt.AddRow(dr)
		}
	}
	return dt
}

func (t *DataTable) Clone() DataTable {
	dt := DataTable{}
	dt.Columns = t.Columns
	return dt
}

func (t *DataTable) NewRow() DataRow {
	row := DataRow{}
	row.Values = make(map[string]interface{})
	row.Table = t
	return row
}

func (t *DataTable) NewColumn(name string) DataColumn {
	column := DataColumn{name, t}
	return column
}

func (t *DataTable) GetRow(index int) DataRow {
	return t.Rows[index]
}

func (t *DataTable) GetColumn(index int) DataColumn {
	return t.Columns[index]
}

func (t *DataTable) AddColumn(column DataColumn) {
	t.Columns = append(t.Columns, column)
}

func (t *DataTable) AddColumnByNames(arr ...string) {
	if arr == nil {
		return
	}

	for i := 0; i < len(arr); i++ {
		column := t.NewColumn(arr[i])
		t.Columns = append(t.Columns, column)
	}
}

func (t *DataTable) InsertColumn(index int, column DataColumn) {
	result := make([]DataColumn, len(t.Columns)+1)
	at := copy(result, t.Columns[:index]) + 1
	result = append(result, column)
	copy(result[at:], t.Columns[index:])
	t.Columns = result
}

func (t *DataTable) InsertColumnByNames(index int, arr ...string) {
	if arr == nil {
		return
	}

	result := make([]DataColumn, len(t.Columns)+len(arr))
	at := copy(result, t.Columns[:index]) + len(arr)
	for i := 0; i < len(arr); i++ {
		column := t.NewColumn(arr[i])
		result = append(result, column)
	}
	copy(result[at:], t.Columns[index:])
	t.Columns = result
}

func (t *DataTable) AddRow(row DataRow) {
	t.Rows = append(t.Rows, row)
}

func (t *DataTable) AddRowByValues(arr ...interface{}) {
	if arr == nil {
		return
	}

	row := t.NewRow()
	for i := 0; i < len(t.Columns); i++ {
		if len(arr) <= i {
			row.Values[t.Columns[i].Name] = nil
		} else {
			row.Values[t.Columns[i].Name] = arr[i]
		}
	}

	t.Rows = append(t.Rows, row)
}

func (t *DataTable) InsertRow(index int, row DataRow) {
	result := make([]DataRow, len(t.Rows)+1)
	at := copy(result, t.Rows[:index]) + 1
	result = append(result, row)
	copy(result[at:], t.Rows[index:])
	t.Rows = result
}

func (t *DataTable) InsertRowByValues(index int, arr ...interface{}) {
	if arr == nil {
		return
	}

	row := t.NewRow()
	for i := 0; i < len(t.Columns); i++ {
		if len(arr) <= i {
			row.Values[t.Columns[i].Name] = nil
		} else {
			row.Values[t.Columns[i].Name] = arr[i]
		}
	}
	result := make([]DataRow, len(t.Rows)+1)
	at := copy(result, t.Rows[:index]) + 1
	result = append(result, row)
	copy(result[at:], t.Rows[index:])
	t.Rows = result
}

func (dt DataTable) ToJson() string {
	str := "["
	rows := dt.Rows
	for i := 0; i < len(rows); i++ {
		if i > 0 {
			str += ","
		}

		str += rows[i].ToJson()
	}
	str += "]"

	return str
}
