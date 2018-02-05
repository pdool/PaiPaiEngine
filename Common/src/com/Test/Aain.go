package main

import (
	"fmt"
	"com/pdool/Core"
	"testing"
)

func TestStack(t *testing.T) {
	//s := Stack{}
	//
	//s.Push("queue")
	//if s.IsEmpty() {
	//	t.Error("Stack is not Empty")
	//}
	//s.Push("stack")
	//
	//if s.Peek() != "stack" {
	//	t.Error("Stack peek is not right")
	//}
	//
	//if s.Pop() != "stack" {
	//	t.Error("Stack pop is not right")
	//}
	//
	//if s.Size() != 1 {
	//	t.Error("Stack size is wrong")
	//}
	//
	//s.Clear()
	//
	//if !s.IsEmpty() {
	//	s.Print()
	//	t.Error("Stack Clear is fail")
	//}

}
func main() {
	r := Core.NewDataRow()
	r.SetValue(0,"吃年糕")

	fmt.Println(r.GetString(0))
}
