package try_test

import "testing"

func TestImplicit(t *testing.T) {
	var a int32 = 1;
	var b int64 = 2;
	b = int64(a);
	t.Log(a, b);
}