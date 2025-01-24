package try_test

import "testing"

type Programmer interface {
	WriteHelloWorld() string
}

/** duckType */
type GOProgrammer struct {

}

func (p *GOProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GOProgrammer)
	t.Log(p.WriteHelloWorld())
}