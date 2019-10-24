package ketama

import (
	"strconv"
	"testing"
)    
  
func TestKetama(t *testing.T) {
	k := New(WithReplicas(7))
	k.Add("1.1.1.1", "2.2.2.2", "3.3.3.3")
	for i := 0; i < 10; i++ {
		node, ok := k.Get(strconv.Itoa(i))
		t.Logf("%v : %v %v", i, node, ok)
	}
	t.Logf("========================")
	k.Remove("1.1.1.1")
	for i := 0; i < 10; i++ {
		node, ok := k.Get(strconv.Itoa(i))
		t.Logf("%v : %v %v", i, node, ok)  
	}  
	t.Logf("========================")
	k.Add("4.4.4.4")
	for i := 0; i < 10; i++ {
		node, ok := k.Get(strconv.Itoa(i))
		t.Logf("%v : %v %v", i, node, ok)
	}
}
