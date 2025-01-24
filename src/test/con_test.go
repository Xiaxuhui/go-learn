package try_test

import "testing"

func TestCondition(t *testing.T) { 
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not in 0-3")
			
		}
	}
}

