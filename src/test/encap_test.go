package try_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
	Age int
}

// func (e *Employee) String() string {
// 	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
// }

/**
结构会被复制，地址会改变
 */
func (e Employee) String() string {
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}
func TestEncap(t *testing.T) {
	e := Employee{"1", "John", 18 }
	e1 := Employee{ Name: "John", Age: 17, Id: "2" }
	e2 := new(Employee) // 返回的是指针
	e2.Name = "John"
	e2.Age = 16
	e2.Id = "3"
	t.Log(e1.Name)
	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
	fmt.Printf("Address is %x", unsafe.Pointer(&e1.Name))
	t.Log(e1.String())
}
