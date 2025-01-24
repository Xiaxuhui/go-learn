package try_test

import "testing"

func TestStrainBytes(t *testing.T) {
	s := "中"
	c := []rune(s) // 代表str的unicode
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}