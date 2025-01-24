package try_test

import "testing"

// func TestMap (t *testing.T) {
// 	m1 := map[int]int{1: 1, 2: 4, 3: 9}
// 	t.Log( m1[2])
// 	t.Logf("len m1=%d", len(m1))
// 	m2 := map[int]int{}
// 	m2[4] = 16
// 	t.Logf("len m2=%d", len(m2))
// 	m3 := make(map[int]int, 10)
// 	t.Logf("len m3=%d", len(m3))
// 	v,ok := m2[1]
// 	t.Log(v, ok)
// 	for k, v := range m1 {
// 		t.Log(k, v)
// 	}
// }

// func TestMaoWithFuncValues(t *testing.T) {
// 	m := map[int]func(op int)int{}
// 	m[1] = func(op int)int{
// 		return op;
// 	}
// 	m[2] = func(op int)int{
// 		return op * op
// 	}
// 	t.Log(m[1](2), m[2](3));
// }

func TestMapSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet));
	delete(mySet, n)
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}