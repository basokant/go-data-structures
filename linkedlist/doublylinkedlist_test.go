package list

import (
	"reflect"
	"testing"
)

func TestDoublyLinkedListString(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	got := list.String()
	want := "[1 2 3 4]"

	if want != got {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewDoublyLinkedListFromArray(t *testing.T) {
	got := NewDoublyLinkedListFromArray([]int{1, 2, 3, 4})

	want := NewDoublyLinkedList[int]()

	want.Append(1)
	want.Append(2)
	want.Append(3)
	want.Append(4)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDoublyLinkedListPrepend(t *testing.T) {
	got := NewDoublyLinkedList[int]()

	got.Prepend(1)
	got.Prepend(2)
	got.Prepend(3)
	got.Prepend(4)

	want := NewDoublyLinkedListFromArray([]int{4, 3, 2, 1})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got.Prepend(5)
	want = NewDoublyLinkedListFromArray([]int{5, 4, 3, 2, 1})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDoublyLinkedListAppend(t *testing.T) {
	got := NewDoublyLinkedList[int]()

	got.Append(1)
	got.Append(2)
	got.Append(3)
	got.Append(4)

	want := NewDoublyLinkedListFromArray([]int{1, 2, 3, 4})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got.Append(5)
	want = NewDoublyLinkedListFromArray([]int{1, 2, 3, 4, 5})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDoublyLinkedListDelete(t *testing.T) {
	node := &DoublyLinkedNode[int]{}

	empty_list := NewDoublyLinkedList[int]()
	_, err := empty_list.Delete(node)

	if err == nil {
		t.Errorf("deleting from an empty list %q did not error", empty_list)
	}

	got := NewDoublyLinkedListFromArray([]int{1, 2, 3, 4, 5})
	data, err := got.Delete(got.head)

	want := NewDoublyLinkedListFromArray([]int{2, 3, 4, 5})

	if err != nil || data != 1 || got.String() != want.String() {
		t.Errorf("got %q, wanted %q", got, want)
	}

	_, err = got.Delete(node)

	if err == nil {
		t.Errorf("deleting node %q that is not in the list %q did not error", node, got)
	}
}

func TestDoublyLinkedListSearch(t *testing.T) {
	empty_list := NewDoublyLinkedList[int]()
	_, err := empty_list.Search(5)

	if err == nil {
		t.Errorf("searching from an empty list %q did not error", empty_list)
	}

	list := NewDoublyLinkedListFromArray([]int{1, 2, 3, 4, 5})
	got, err := list.Search(3)

	want := list.head.next.next

	if err != nil || !reflect.DeepEqual(want, got) {
		t.Errorf("got %q, wanted %q", got, want)
	}

	_, err = list.Search(6)

	if err == nil {
		t.Errorf("searching for value %q that is not in the list %q did not error", 6, list)
	}
}
