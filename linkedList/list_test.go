package linkedList

import "testing"

func TestLinkedList(t *testing.T) {
	list := New()
	if actualValue := list.Empty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
	if actualValue := list.size; actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}

	list.Add(1)
	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.size; actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	list.Add(2)
	if actualValue := list.Empty(); actualValue != false {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
	if actualValue := list.size; actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
}
