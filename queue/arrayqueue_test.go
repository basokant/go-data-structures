package queue

import (
	"testing"
)

func TestArrayQueueDequeue(t *testing.T) {
	emptyQueue := New[int]()

	_, err := emptyQueue.Dequeue()

	if err == nil {
		t.Error("error expected, should not be able to Dequeue an empty queue")
	}

	data := []int{1, 2, 3, 4}
	populatedQueue := NewFromArray[int](data)

	for _, v := range data {
		top, err := populatedQueue.Dequeue()

		if err != nil {
			t.Error(err)
		} else if top != v {
			t.Errorf("got %d, wanted %d", top, v)
		}
	}

	_, err = populatedQueue.Dequeue()

	if err == nil {
		t.Error("error expected, should not be able to Dequeue an empty queue")
	}
}

func TestArrauQueueIsEmpty(t *testing.T) {

}
