package demo

import (
	"fmt"
	"net/http"
)

//进程间全局变量测试

type DAA struct {
	Url string
}

func (db *DAA) CreateUser(url string) {
	fmt.Println("create user")
	db.Url = url
}

func hello(db DAA) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, db.Url)
	}
}
