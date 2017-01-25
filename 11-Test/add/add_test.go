package add

import "testing"

func TestAddResult(t *testing.T) {
	actual := Add(1, 2)
	expected := 3
	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
