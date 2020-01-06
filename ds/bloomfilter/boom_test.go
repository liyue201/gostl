package bloom

import (
	"testing"
)

func TestBoomfilter(t *testing.T) {
	b := New(10000, 7, WithThreadSafe())

	if b.Contains("aa") {
		t.Fatalf("expect false, but get true")
	}

	b.Add("aa")
	if !b.Contains("aa") {
		t.Fatalf("expect true, but get false")
	}
	other := NewFromData(b.Data())

	if !other.Contains("aa") {
		t.Fatalf("expect true, but get false")
	}

	b = NewWithEstimates(100000, 0.0001)
	b.Add("bbbbb")
	if !b.Contains("bbbbb") {
		t.Fatalf("expect true, but get false")
	}
}
