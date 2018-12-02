package controler

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"sc.service/business/upload"

	"github.com/gin-gonic/gin"
	"github.com/go-yaml/yaml"
	"sc.service/model"
)

// ImportExcel 导入Excel文件
func ImportExcel(c *gin.Context) {
	// 读取Excel
	file, header, err := c.Request.FormFile("file")
	filename := header.Filename

	// 读取配置
	conf := new(model.Conf)
	yamlFile, err := ioutil.ReadFile("config/conf.yaml")
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatal(err)
	}

	// 保存Excel到上传目录
	path := conf.UploadFolder + filename
	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	// 读取Excel内容
	xlsImpl := new(upload.XlsReader)
	xlsImpl.InitWorkBook(path)
	datas := xlsImpl.ReadData(conf.ColumnMap)

	c.JSON(http.StatusOK, datas)
}
