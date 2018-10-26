package excel_mgr

import "github.com/extrame/xls"
import "fmt"
import "go_web/models"

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

	if sheet1 := this.xlFile.GetSheet(0); sheet1 != nil {
	}
	
	return nil
}

func (this *XlsReader)validate() error {
	return nil
}


func (this *XlsReader)ExampleWorkBook_NumberSheets(filePath string) {
	if xlFile, err := xls.Open(filePath, "utf-8"); err == nil {
		for i := 0; i < xlFile.NumSheets(); i++ {
			sheet := xlFile.GetSheet(i)
			fmt.Println(sheet.Name)
		}
	}
}

//Output: read the content of first two cols in each row
func ExampleWorkBook_GetSheet(filePath string) {
	if xlFile, err := xls.Open(filePath, "utf-8"); err == nil {
		if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
			fmt.Print("Total Lines ", sheet1.MaxRow, sheet1.Name)
			col1 := sheet1.Row(0).Col(0)
			col2 := sheet1.Row(0).Col(0)
			for i := 0; i <= (int(sheet1.MaxRow)); i++ {
				row1 := sheet1.Row(i)
				col1 = row1.Col(0)
				col2 = row1.Col(1)
				fmt.Print("\n", col1, ",", col2)
			}
		}
	}
}