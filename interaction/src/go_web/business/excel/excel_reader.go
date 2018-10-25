package excel


import (
	"fmt"
    "github.com/Luxurioust/excelize"
)

type ExcelReader struct {
}

//https://blog.csdn.net/zdyueguanyun/article/details/64588339
//colMapping ：column name
func (this *ExcelReader)ReadFormFile(fileName string, colMapping map[string]string) (string, error) {
	xlsx, err := excelize.OpenFile(fileName)

	if (err != nil) {
		fmt.Println(err)
	}

	for k,v := range colMapping {
		fmt.Print(k)
		fmt.Println(v)
	}
	cell := xlsx.GetCellValue("Sheet1", "A1")
	fmt.Println(cell)
	
	return "",nil
}