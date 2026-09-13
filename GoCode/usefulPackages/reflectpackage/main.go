package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println(reflect.TypeOf(5.0))
	fmt.Println(reflect.TypeOf("sample_string"))
	fmt.Println(reflect.TypeOf(true))

}