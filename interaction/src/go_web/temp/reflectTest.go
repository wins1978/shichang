package temp

import (
	"fmt"
	"container/list"
)

func ReflectDemo() {
	nums := list.New()
	nums.PushBack("aaa")
	nums.PushBack("bbb")
	nums.PushBack("ccc")

	for nn := nums.Front(); nn != nil; nn = nn.Next() {
		fmt.Println(nn.Value)
	}

	
}