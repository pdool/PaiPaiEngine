package main

import (
	"fmt"
	"reflect"
)

type NetModule struct {

}
func(net *NetModule)printS(name string){
	fmt.Println(name)
}
func main() {
	t := reflect.ValueOf("NetModule").Type()
	v := reflect.New(t).Elem()
	name := v.MethodByName("printS")
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf("reflection test")
	name.Call(params)
	fmt.Print(name)
}
