package stack

import "testing"

func TestListStackEmpty(t *testing.T) {
	emptyStack := NewArrayStack[int]()

	if !emptyStack.Empty() {
		t.Error("Stack should be empty, was found populated.")
	}

	slice := []int{0, 1}
	populatedStack := NewListStackFromArray[int](slice)

	if populatedStack.Empty() {
		t.Error("Stack should be populated, was found empty.")
	}
}

func TestListStackPeek(t *testing.T) {
	st := NewListStack[int]()

	_, err := st.Peek()
	if err == nil {
		t.Error("Peek should error if there are no elements, it did not error.")
	}

	slice := []int{0, 1, 2}
	st = NewListStackFromArray[int](slice)

	top, err := st.Peek()
	if err != nil {
		t.Error("Peek should not error if array is populated.")
	} else if top != 0 {
		t.Errorf("got %q, wanted %q", top, 0)
	}
}

func TestListStackPop(t *testing.T) {
	st := NewListStack[int]()

	_, err := st.Pop()
	if err == nil {
		t.Error("Pop should error if there are no elements, it did not error.")
	}

	slice := []int{0, 1, 2}
	st = NewListStackFromArray[int](slice)

	top, err := st.Pop()
	if err != nil {
		t.Error("Pop should not error if array is populated.")
	} else if top != 0 {
		t.Errorf("got %q, wanted %q", top, 0)
	}

	top, err = st.Pop()
	if err != nil {
		t.Error("Pop should not error if array is populated.")
	} else if top != 1 {
		t.Errorf("got %q, wanted %q", top, 1)
	}
}

func TestListStackPush(t *testing.T) {
	st := NewListStack[int]()

	err := st.Push(0)
	if err != nil {
		t.Error("Push should not error when called on empty stack.")
	}

	err = st.Push(1)
	if err != nil {
		t.Error("Push encountered an error when it should not have.")
	}

	top, err := st.Peek()
	if err != nil {
		t.Errorf("Peek encountered an error when it should not have.")
	} else if top != 1 {
		t.Errorf("got %q, wanted %q", top, 1)
	}

	top, err = st.Pop()
	if err != nil {
		t.Errorf("Pop encountered an error when it should not have.")
	} else if top != 1 {
		t.Errorf("got %q, wanted %q", top, 1)
	}
}

func TestListStackSearch(t *testing.T) {
	st := NewListStack[int]()

	index, err := st.Search(0)

	if err == nil || index >= 0 {
		t.Error("Empty stack should not be searchable.")
	}

	slice := []int{0, 1, 2, 3, 4}
	st = NewListStackFromArray[int](slice)

	for i, v := range slice {
		index, err = st.Search(v)

		if err != nil {
			t.Errorf("Error occured while searching stack for value %q", v)
		} else if index != i {
			t.Errorf("got %q, wanted %q", index, i)
		}
	}
}
