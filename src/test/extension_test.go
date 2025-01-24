package try_test

import (
	"fmt"
	"testing"
)

type Pet struct {}

/** 不支持重载 */
type Dog struct {
	Pet
   }

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Print(" ", host)
}


   

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()
}