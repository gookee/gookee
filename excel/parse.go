package excel

import (
	"db"
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"utility"
)

type Workbook struct {
	Worksheet []Worksheet `xml:"Worksheet"`
}

type Worksheet struct {
	Name  string `xml:"ss:Name,attr"`
	Table Table  `xml:"Table"`
}

type Table struct {
	ExpandedColumnCount string `xml:"ss:ExpandedColumnCount,attr"`
	ExpandedRowCount    string `xml:"ss:ExpandedRowCount,attr"`
	FullColumns         string `xml:"ss:FullColumns,attr"`
	FullRows            string `xml:"ss:FullRows,attr"`
	DefaultColumnWidth  string `xml:"ss:DefaultColumnWidth,attr"`
	DefaultRowHeight    string `xml:"ss:DefaultRowHeight,attr"`
	Row                 []Row  `xml:"Row"`
}

type Row struct {
	Cell []Cell `xml:"Cell"`
}

type Cell struct {
	Type string `xml:"ss:Type"`
	Data string `xml:"Data"`
}

func XMLToExcel(r io.Reader) *Excel {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Println(err)
	}
	var workbook Workbook
	err = xml.Unmarshal(content, &workbook)
	if err != nil {
		fmt.Println(err)
	}

	excel := &Excel{}
	for _, sheet := range workbook.Worksheet {
		sheetName := sheet.Name
		dt := db.NewDataTable()
		for i, row := range sheet.Table.Row {
			if i == 0 {
				for _, cell := range row.Cell {
					dt.AddColumnByNames(cell.Data)
				}
			} else {
				dr := dt.NewRow()
				for j, cell := range row.Cell {
					dr.Values[dt.Columns[j].Name] = cell.Data
				}
				dt.AddRow(dr)
			}
		}
		excel.AddSheet(sheetName, dt)
	}
	return excel
}

func CSVToDataTable(r io.Reader) db.DataTable {
	reader := csv.NewReader(r)
	dt := db.NewDataTable()
	i := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return dt
		}
		if i == 0 {
			for _, v := range record {
				dt.AddColumnByNames(strings.Trim(v, "\xEF\xBB\xBF"))
			}
		} else {
			dr := dt.NewRow()
			for x, v := range record {
				dr.Values[dt.Columns[x].Name] = v
			}
			dt.AddRow(dr)
		}
		i++
	}
	return dt
}

func DataTableToCSV(r io.Writer, dt db.DataTable) {
	r.Write([]byte("\xEF\xBB\xBF"))
	w := csv.NewWriter(r)
	colsCount := len(dt.Columns)
	cols := make([]string, colsCount, colsCount)
	for i := 0; i < colsCount; i++ {
		cols[i] = dt.Columns[i].Name
	}
	w.Write(cols)
	for _, v := range dt.Rows {
		arr := make([]string, colsCount, colsCount)
		for i := 0; i < colsCount; i++ {
			arr[i] = utility.ToStr(v.Values[dt.Columns[i].Name])
		}
		w.Write(arr)
	}
	w.Flush()
}
