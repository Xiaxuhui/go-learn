package try_test

import "testing"

// func TestSlices(t *testing.T) {
// 	var s0 []int
// 	s0 = append(s0, 1)
// 	t.Log(len(s0), cap(s0))
// 	s1 := []int{1, 2, 3, 4 }
// 	t.Log(len(s1), cap(s1))
// 	s2 := make([]int, 3, 5)
// 	t.Log(len(s2), cap(s2))
// }

// func TestSLiceGrowing(t *testing.T) {
// 	s := []int{}
// 	for i := 0; i < 10; i++ {
// 		s = append(s, i) // 存储空间变化，地址变化后赋值给s
// 		t.Log(len(s), cap(s))
// 	}
// }

func TestSliceShareMemory (t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec" }
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknown"
	t.Log(Q2)
}