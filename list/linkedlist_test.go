package list

import (
	"reflect"
	"testing"
)

func TestLinkedListString(t *testing.T) {
	list := NewLinkedList[int]()

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

func TestNewLinkedListFromArray(t *testing.T) {
	got := NewLinkedListFromArray([]int{1, 2, 3, 4})

	want := NewLinkedList[int]()

	want.Append(1)
	want.Append(2)
	want.Append(3)
	want.Append(4)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLinkedListPrepend(t *testing.T) {

}

func TestLinkedListAppend(t *testing.T) {

}

func TestLinkedListDelete(t *testing.T) {

}

func TestLinkedListSearch(t *testing.T) {

}
