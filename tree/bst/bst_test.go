package bst

import (
	"testing"
)

func TestBst(t *testing.T) {
	bst := New()

	if actualValue := bst.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := bst.size; actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	bst.Add(3)
	if actualValue := bst.size; actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	if actualValue := bst.root.key; actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	bst.Add(2)
	if actualValue := bst.size; actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue := bst.root.left.key; actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	bst.Add(4)
	if actualValue := bst.size; actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue := bst.root.right.key; actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}

	bst.Add(1)
	if actualValue := bst.size; actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue := bst.root.left.left.key; actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
}
