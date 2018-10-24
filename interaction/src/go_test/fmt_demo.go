package main

import (
	"fmt"
)

type Per struct {
	N string
	A uint
}

func (a Per) String() string {
	return fmt.Sprintf("--%v --%d",a.N,a.A)
}

func Fmtdemo() {
	a:= Per{
		N:"wins",
		A:20,
	}
	fmt.Println(a)
}