package main

import (
	"fmt"
	"reflect"
)


type User1 struct {
	name string `json:name-field`
	age  int
}
func (u *User1)Sss() {
	fmt.Println("user == sssss")
}

func getStructTag(f reflect.StructField) string {
	fmt.Println("ss")
	return string(f.Tag)
}

func main() {

	user := &User1{"John Doe The Fourth", 20}

	/*field, ok := reflect.TypeOf(user).Elem().FieldByName("name")
	if !ok {
		panic("Field not found")
	}
	fmt.Println(getStructTag(field))*/

	vof := reflect.ValueOf(user)
	var arr []reflect.Value
	vof.MethodByName("Sss").Call(arr)

}


