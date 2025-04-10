package linkedlist

import (
	"testing"
)

func TestLinkedList_AddPrependAppend(t *testing.T) {
	ll := New[int]()

	// Test append to an empty list
	ll.append(1)
	if ll.len() != 1 || ll.head.val != 1 || ll.tail.val != 1 {
		t.Errorf("Append failed: expected head and tail to be 1, got head = %v, tail = %v", ll.head.val, ll.tail.val)
	}

	// Test prepend to a list
	ll.prepend(0)
	if ll.len() != 2 || ll.head.val != 0 || ll.tail.val != 1 {
		t.Errorf("Prepend failed: expected head 0 and tail 1, got head = %v, tail = %v", ll.head.val, ll.tail.val)
	}

	// Test add in the middle
	ll.add(2, 1) // List should be 0, 2, 1
	if ll.get(1) != 2 {
		t.Errorf("Add failed: expected index 1 to be 2, got %v", ll.get(1))
	}
}

func TestLinkedList_RemoveShiftPop(t *testing.T) {
	ll := New[int]()
	ll.append(1)
	ll.append(2)
	ll.append(3) // List is 1, 2, 3

	// Test shift
	if val := ll.shift(); val != 1 || ll.len() != 2 {
		t.Errorf("Shift failed: expected return 1 and length 2, got return %v and length %v", val, ll.len())
	}

	// Test pop
	if val := ll.pop(); val != 3 || ll.len() != 1 {
		t.Errorf("Pop failed: expected return 3 and length 1, got return %v and length %v", val, ll.len())
	}

	// Test remove middle
	ll.prepend(1) // List is 1, 2
	if val := ll.remove(1); val != 2 || ll.len() != 1 || ll.head.val != 1 || ll.tail.val != 1 {
		t.Errorf("Remove failed: expected return 2 and remaining list [1], got return %v and head %v tail %v", val, ll.head.val, ll.tail.val)
	}
}

func TestLinkedList_Get(t *testing.T) {
	ll := New[int]()
	ll.append(10)
	ll.append(20)
	ll.append(30) // List is 10, 20, 30

	if val := ll.get(1); val != 20 {
		t.Errorf("Get failed: expected 20, got %v", val)
	}

	if val := ll.get(0); val != 10 {
		t.Errorf("Get failed: expected 10, got %v", val)
	}

	if val := ll.get(2); val != 30 {
		t.Errorf("Get failed: expected 30, got %v", val)
	}
}

func TestLinkedList_BoundaryConditions(t *testing.T) {
	ll := New[int]()

	// Test removing from an empty list
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when removing from empty list, no panic occurred")
		}
	}()
	ll.remove(0)

	// More boundary conditions can be added here
}
