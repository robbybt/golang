package variable

import (
	"fmt"
	//"strconv"
	"encoding/json"
)

func Testvar() {
	var i, j int = 1, 2
	fmt.Println("orang", i, j)
	//varmap := map[string]int{}
	varmap := make(map[string]int)
	varmap["satu"] = 1
	varmap["dua"] = 2
	varmap["tiga"] = 2
	delete(varmap, "tiga")
	v, ok := varmap["dua"]
	fmt.Println("The value:", v, "Present?", ok)
	v, ok = varmap["tiga"]
	fmt.Println("The value:", v, "Present?", ok)
	for k, v := range varmap {
		fmt.Println(k, "=", v)
	}
}

func Printapapun(all ...interface{}) {
	fmt.Println(all...)
}

func Ketiktodos() {
	//fmt.Println("asd")
	var vartodos todos
	vartodos = append(vartodos, todo{1, "asd"})
	vartodos = append(vartodos, todo{2, "asd"})
	vartodos = append(vartodos, todo{3, "asd"})

	//===============================================================
	varjson := map[string]interface{}{}
	varjson["id"] = 4
	varjson["name"] = "cacat"
	fmt.Printf("go data : %+v\n", varjson)
	tole, err := json.Marshal(varjson)
	fmt.Println(string(tole))
	if err != nil {
		fmt.Println(err)
	}
	var cacat interface{}
	json.Unmarshal(tole, &cacat)
	fmt.Printf("hasil unmarshal : %v\n", cacat)
	//================================================================
	for _, v := range vartodos {
		fmt.Printf("%+v\n", v)
	}
}

type todo struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

type todos []todo
