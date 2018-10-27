package excel_mgr

import (
	"reflect"
	"github.com/extrame/xls"
	"go_web/models"
)

//https://github.com/extrame/xls/blob/master/example_test.go
type XlsReader struct {
	xlFile *xls.WorkBook
}

//Init
func (this *XlsReader)InitWorkBook(filePath string) {
	//load exc to instance
	if xlFile, err := xls.Open(filePath, "utf-8"); err == nil {
		this.xlFile = xlFile
	}
	//load style config
}

func (this *XlsReader)ReadData(colMap map[string]string) []models.InputInfo {
	err := this.validate()
	if (err != nil) {
		return nil
	}

	sheet1 := this.xlFile.GetSheet(0)
	if (sheet1 == nil) {
		return nil
	}

	//Map FirstHeader
	var colIndexMap map[string]int
	colIndexMap = make(map[string]int)
	captionRow := sheet1.Row(0)
	for idx := captionRow.FirstCol(); idx < captionRow.LastCol(); idx++ {
		cellValue := captionRow.Col(idx)
		prop,ok := colMap[cellValue]
		if (ok == true) {
			colIndexMap[prop] = idx
		}
	}

	// build data
	var infoList []models.InputInfo

	for i := 1; i<= int(sheet1.MaxRow); i++ {
		row := sheet1.Row(i)
		info := fillDataByRow(row,colIndexMap)
		infoList = append(infoList,info)
	}	
	
	return infoList
}

func fillDataByRow(row *xls.Row, colMap map[string]int) models.InputInfo {
	var info models.InputInfo
	pt := reflect.ValueOf(&info).Elem()
	for colN,colIdx := range colMap {
		val := row.Col(colIdx)
		pt.FieldByName(colN).SetString(val)
	}

	return info
}

func (this *XlsReader)validate() error {
	return nil
}