package main

import (
	"gorm_demo/conn"
	"gorm_demo/dmloperator"
	"time"
)

func main() {
	cntList := [100]int{}
	conn.InitDB()
	for idx, _ := range cntList {
		go dmloperator.InsertOne(idx)
	}

	time.Sleep(time.Duration(200) * time.Second)
}
