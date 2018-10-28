package excel_mgr


import (
	"fmt"
    "github.com/Luxurioust/excelize"
)

type ExcelReader struct {
	xlsx *excelize.File
}

//https://blog.csdn.net/zdyueguanyun/article/details/64588339
//colMapping ：column name
func (this *ExcelReader)ReadFormFile(fileName string, colMapping map[string]string) (string, error) {
	
	/*
	xfile, err := excelize.OpenFile(fileName)
	
	if (err != nil) {
		fmt.Println(err)
	}
	this.xlsx = xfile
	
	cell := this.xlsx.GetCellValue("Sheet1", "A1")
	fmt.Println(cell)
	this.readHeader()
	*/
	return "",nil
}

func (this *ExcelReader)readHeader() {
	fmt.Println("start....")
	rows := this.xlsx.GetRows("订单明细")
	rowLen := len(rows)
	fmt.Println(rowLen)

	for _, row := range rows {
		for _, cell := range row {
			fmt.Print(cell, "\t")
		}
		fmt.Println()
	}
}