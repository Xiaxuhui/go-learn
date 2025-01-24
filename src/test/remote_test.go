package try_test

import (
	"go_learn/src/series"
	"testing"
)


func TestPackage(t *testing.T) {
	
	sum := series.GetSum(1, 2)
	t.Logf("Package result %d", sum);
}