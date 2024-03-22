package stack

import "testing"

func TestArrayStackEmpty(t *testing.T) {
	emptyStack := NewArrayStack[int]()

	if !emptyStack.IsEmpty() {
		t.Error("Stack should be empty, was found populated.")
	}

	slice := []int{0, 1}
	populatedStack := NewArrayStackFromArray[int](slice)

	if populatedStack.IsEmpty() {
		t.Error("Stack should be populated, was found empty.")
	}
}

func TestArrayStackPeek(t *testing.T) {
	stack := NewArrayStack[int]()

	_, err := stack.Peek()
	if err == nil {
		t.Error("Peek should error if there are no elements, it did not error.")
	}

	slice := []int{0, 1, 2}
	stack = NewArrayStackFromArray[int](slice)

	top, err := stack.Peek()
	if err != nil {
		t.Error("Peek should not error if array is populated.")
	} else if top != 0 {
		t.Errorf("got %q, wanted %q", top, 0)
	}
}

func TestArrayStackPop(t *testing.T) {
	stack := NewArrayStack[int]()

	_, err := stack.Pop()
	if err == nil {
		t.Error("Pop should error if there are no elements, it did not error.")
	}

	slice := []int{0, 1, 2}
	stack = NewArrayStackFromArray[int](slice)

	top, err := stack.Pop()
	if err != nil {
		t.Error("Pop should not error if array is populated.")
	} else if top != 0 {
		t.Errorf("got %q, wanted %q", top, 0)
	}

	top, err = stack.Pop()
	if err != nil {
		t.Error("Pop should not error if array is populated.")
	} else if top != 1 {
		t.Errorf("got %q, wanted %q", top, 1)
	}
}

func TestArrayStackPush(t *testing.T) {
	stack := NewArrayStack[int]()

	err := stack.Push(0)
	if err != nil {
		t.Error("Push should not error when called on empty stack.")
	}

	err = stack.Push(1)
	if err != nil {
		t.Error("Push encountered an error when it should not have.")
	}

	top, err := stack.Peek()
	if err != nil {
		t.Errorf("Peek encountered an error when it should not have.")
	} else if top != 1 {
		t.Errorf("got %q, wanted %q", top, 1)
	}

	top, err = stack.Pop()
	if err != nil {
		t.Errorf("Pop encountered an error when it should not have.")
	} else if top != 1 {
		t.Errorf("got %q, wanted %q", top, 1)
	}
}

func TestArrayStackSearch(t *testing.T) {
	stack := NewArrayStack[int]()

	index, err := stack.IndexOf(0)

	if err == nil || index >= 0 {
		t.Error("Empty stack should not be searchable.")
	}

	slice := []int{0, 1, 2, 3, 4}
	stack = NewArrayStackFromArray[int](slice)

	for i, v := range slice {
		index, err = stack.IndexOf(v)

		if err != nil {
			t.Errorf("Error occured while searching stack for value %q", v)
		} else if index != i {
			t.Errorf("got %q, wanted %q", index, i)
		}
	}
}
