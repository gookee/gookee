package excel

import (
	"db"
	"utility"
)

type Excel struct {
	Sheets []Sheet
}

type Sheet struct {
	Name  string
	Table db.DataTable
}

func NewExcel(sheetName string, dt db.DataTable) *Excel {
	st := Sheet{sheetName, dt}
	return &Excel{[]Sheet{st}}
}

func (e *Excel) AddSheet(sheetName string, dt db.DataTable) {
	e.Sheets = append(e.Sheets, Sheet{sheetName, dt})
}

func (e *Excel) ToString() string {
	xmlStr := "<?xml version=\"1.0\"?>\r\n" +
		"<?mso-application progid=\"Excel.Sheet\"?>\r\n" +
		"<Workbook xmlns=\"urn:schemas-microsoft-com:office:spreadsheet\"\r\n" +
		" xmlns:o=\"urn:schemas-microsoft-com:office:office\"\r\n" +
		" xmlns:x=\"urn:schemas-microsoft-com:office:excel\"\r\n" +
		" xmlns:ss=\"urn:schemas-microsoft-com:office:spreadsheet\"\r\n" +
		" xmlns:html=\"http://www.w3.org/TR/REC-html40\">\r\n"
	for _, sheet := range e.Sheets {
		xmlStr += " <Worksheet ss:Name=\"" + sheet.Name + "\">\r\n"
		xmlStr += "  <Table ss:ExpandedColumnCount=\"" + utility.ToStr(len(sheet.Table.Columns)) + "\" ss:ExpandedRowCount=\"" + utility.ToStr(len(sheet.Table.Rows)+1) + "\" x:FullColumns=\"1\"\r\n" +
			"x:FullRows=\"1\" ss:DefaultColumnWidth=\"100\" ss:DefaultRowHeight=\"13.5\">\r\n"
		xmlStr += "   <Row>\r\n"
		for _, cols := range sheet.Table.Columns {
			xmlStr += "    <Cell><Data ss:Type=\"String\">" + cols.Name + "</Data></Cell>\r\n"
		}
		xmlStr += "   </Row>\r\n"
		for _, dr := range sheet.Table.Rows {
			xmlStr += "   <Row>\r\n"
			for _, cols := range sheet.Table.Columns {
				xmlStr += "    <Cell><Data ss:Type=\"String\">" + utility.ToStr(dr.Values[cols.Name]) + "</Data></Cell>\r\n"
			}
			xmlStr += "   </Row>\r\n"
		}
		xmlStr += "  </Table>\r\n"
		xmlStr += " </Worksheet>\r\n"
	}
	xmlStr += "</Workbook>"
	return xmlStr
}
