package set

import "testing" 
 
func TestSet(t *testing.T) {
	cmp := func(a, b interface{}) int {
		if a.(int) == b.(int) {
			return 0  
		}
		if a.(int) < b.(int) {
			return -1
		}
		return 1
	}
	s := New(cmp)
	for i := 10; i >= 1; i-- {
		s.Insert(i) 
	}       
    
	for iter := s.Begin(); !iter.Equal(s.End()); iter = iter.Next() {
		t.Logf("%v\n", iter.Value())
	}  
	t.Logf("=======================" )

	for iter := s.RBegin(); !iter.Equal(s.REnd()); iter = iter.Next() {
		t.Logf("%v\n", iter.Value())
	}
	t.Logf("=======================" )

	if !s.Contains(5) {
		t.Logf("Contains(10) error\n" )
	}
	s.Erase(5)
	if s.Contains(5) {
		t.Logf("s.Erase(10) Contains(10) error\n" )
	}
  
	for iter := s.Begin(); !iter.Equal(s.End()); iter = iter.Next() {
		t.Logf("%v\n", iter.Value())
	}

	if s.Size() != 9{
		t.Logf("s.Size() error: %v\n" , s.Size())
	}
}
