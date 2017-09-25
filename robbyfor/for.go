package robbyfor

import (
	"fmt"
	"strconv"
)

var tesvar int

func Forbiasa() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(tesvar)
}

func Foreachbiasa() {
	arr := []interface{}{"a1", 2, "a3", 4}
	for _, v := range arr {
		fmt.Println(v)
	}
}

type Person struct {
	Fname string
	Lname string
}

func ForeachStruct(p []*Person) {
	for k, v := range p {
		fmt.Println("orang " + strconv.Itoa(k) + " ->  " + p[k].Fname + " " + v.Lname)
	}
}

func Whilebiasa() {
	var i int
	for i <= 10 {
		i++
		fmt.Println(i)
	}
}
