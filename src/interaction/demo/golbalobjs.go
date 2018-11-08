package demo

import (
	"fmt"
	"interaction/common"
	"interaction/models"
)

var idx int

func CreateUser() {
	fmt.Println("create user")

	var user models.User

	common.Db.First(&user)

	fmt.Println(user.Name)

	if idx == 0 {
		idx = 1
	} else if idx == 1 {
		idx = 2
		aa := common.Db.Where("ssss", "ddd").First(&user)
		if aa.Error != nil {
			fmt.Println(aa.Error)
		}
		panic("err------------------")
	} else {
		idx = 2
	}

}
