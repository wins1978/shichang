package upload

import (
	"reflect"
	"strconv"

	"sc.service/model"

	"github.com/extrame/xls"
)

// https://github.com/extrame/xls/blob/master/example_test.go

// XlsReader WorkBook Reader
type XlsReader struct {
	xlFile *xls.WorkBook
}

// InitWorkBook 初始化
func (c *XlsReader) InitWorkBook(filePath string) {
	//load exc to instance
	if xlFile, err := xls.Open(filePath, "utf-8"); err == nil {
		c.xlFile = xlFile
	}
	//load style config
}

// ReadData 读取Excel数据并转化为实体
func (c *XlsReader) ReadData(colMap map[string]string) []model.InputInfo {
	err := c.validate()
	if err != nil {
		return nil
	}

	sheet1 := c.xlFile.GetSheet(0)
	if sheet1 == nil {
		return nil
	}

	// 匹配第一行标题
	var colIndexMap map[string]int
	colIndexMap = make(map[string]int)
	captionRow := sheet1.Row(0)
	for idx := captionRow.FirstCol(); idx < captionRow.LastCol(); idx++ {
		cellValue := captionRow.Col(idx)
		for prop, value := range colMap {
			if value == cellValue {
				colIndexMap[prop] = idx
			}
		}
	}

	// 读取数据（从第二行开始）
	var infoList []model.InputInfo
	for i := 1; i <= int(sheet1.MaxRow); i++ {
		row := sheet1.Row(i)
		info := fillDataByRow(row, colIndexMap, i)
		infoList = append(infoList, info)
	}

	return infoList
}

func fillDataByRow(row *xls.Row, colMap map[string]int, index int) model.InputInfo {
	var info model.InputInfo
	pt := reflect.ValueOf(&info).Elem()
	for colN, colIdx := range colMap {
		val := row.Col(colIdx)
		pt.FieldByName(colN).SetString(val)
	}
	pt.FieldByName("RowIndex").SetString(strconv.Itoa(index))

	return info
}

func (c *XlsReader) validate() error {
	return nil
}
