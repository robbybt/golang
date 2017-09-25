package main

import (
	"../golang/robbyfor"
	//"fmt"
	"fmt"
)

func main() {
	fmt.Println("start project")
	robbyfor.Forbiasa()
	robbyfor.Foreachbiasa()

	// c := [3]int{}
	// c[0] = 1
	// c[1] = 2
	// for _, v := range c {
	// 	fmt.Println(v)
	// }
	a := robbyfor.Person{"a", "b"}
	arr := []*robbyfor.Person{}
	arr = append(arr, &a)
	arr = append(arr, &robbyfor.Person{"a1", "b2"})
	robbyfor.ForeachStruct(arr)
	robbyfor.Whilebiasa()
}
