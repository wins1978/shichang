package dmloperator_test

import (
	"gorm_demo/dmloperator"
	"testing"
	"time"
)

//=======================
//协程多并发测试
//=======================

func TestGoInsertOne(t *testing.T) {
	cntList := [100]int{}

	for idx, _ := range cntList {
		go dmloperator.InsertOne(idx)
	}

	time.Sleep(time.Duration(100) * time.Second)
}
