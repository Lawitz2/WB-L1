package main

import (
	"fmt"
	"reflect"
)

func Ex14() {
	var unknown interface{}
	unknown = "kek"
	fmt.Println(reflect.TypeOf(unknown))

	unknown = 1
	fmt.Println(reflect.TypeOf(unknown))

	unknown = true
	fmt.Println(reflect.TypeOf(unknown))

	unknown = make(chan struct{})
	fmt.Println(reflect.TypeOf(unknown))
}
