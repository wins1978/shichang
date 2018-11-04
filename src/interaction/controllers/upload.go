package controllers

/*
1) Upload Exc file to server path
2) Read from Exc plugin
3) Map with Fix DB column
4) Validate column -- not null
5) Show UI with warning

6) Save sourcedata to DB -- EXCEL_IMPORT_DAILY
7) Create Daily Report
8) Create Daily WX Report

EXCEL======
1) Read from Exc --Ignore DATA, params--column mapping
2) Write to Exc -- Ignore DATA
*/

import (
	"interaction/business/excel_mgr"
	"interaction/config"

	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

//https://www.cnblogs.com/goloving/p/8967865.html 上传文件示例
//http://localhost:8080/Upload/UpFile
func (this *UploadController) UpFile() {
	f, h, _ := this.GetFile("file")
	defer f.Close()

	path := config.UPLOAD_FOLDER + h.Filename
	this.SaveToFile("file", path)

	var colMap = initClumnMap()

	xlsImpl := new(excel_mgr.XlsReader)
	xlsImpl.InitWorkBook(path)
	datas := xlsImpl.ReadData(colMap)

	this.Data["datas"] = datas
	this.Data["msg"] = ""
	this.Ctx.Output.JSON(this.Data["datas"], false, false)
}

//定义表头字段对于关系
func initClumnMap() map[string]string {
	var colMap map[string]string
	colMap = make(map[string]string)

	colMap["首次称重操作员"] = "Date"
	colMap["磅单编号"] = "NumberFlag"
	colMap["首次称重日期"] = "Date"
	colMap["毛重时间"] = "Time"
	colMap["车号"] = "CarNumber"
	colMap["货物名称"] = "ShopNumber"
	colMap["毛重"] = "GrossWeight"
	colMap["皮重"] = "TareWeight"
	colMap["净重"] = "NetWeight"
	colMap["收货单位"] = "TakeDept"
	colMap["发货单位"] = "SendOutDept"
	return colMap
}
