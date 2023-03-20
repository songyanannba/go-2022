package main

import "fmt"

func main() {
	m := make(map[string]string)
	fmt.Println(m)

	var m2 map[string]string
	fmt.Println(m2)

	m3 := map[string]string{
		"鲁菜": "大肠",
		"川菜": "豆腐",
	}

	fmt.Println(m3)
	m4 := map[string]map[int]string{
	}
	fmt.Println(m4)

	m5 := map[int]string{}

	m5[1] = "ssss"
	m5[2] = "dd"
	m5[3] = "hh"
	m5[4] = "ff"
	fmt.Println(m5)

	for index ,val := range m5{
		fmt.Println(index,val)
	}

	v2 , ok := m5[22]
	if ok {
		fmt.Println(v2)
	} else {
		fmt.Println("djdj")
	}

	//删除
	delete(m5 , 2)
	fmt.Println(m5)
}
