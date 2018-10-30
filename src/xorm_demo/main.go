package main

import (
	"fmt"
	"xorm_demo/conn"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	conn.InitDB()
}

func main() {
	fmt.Println("starting main")
}
