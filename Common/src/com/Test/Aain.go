package main

import (
	"fmt"
	"testing"
	"com/pdool/Core"
)
type TestI struct {

}
func (t *TestI)Handle(obj int, scenId int, groupId int, args interface{}){
	fmt.Println(t)
}
func TestStack(t *testing.T) {
	t1 := new(TestI)
	t2 := new(TestI)
	fmt.Println(t1 == t2)

}
func main() {
	t1 := Core.NewGuid()
	t2 := Core.NewGuid()
	fmt.Println(t1 == t2)
}
