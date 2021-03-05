package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := New("sample_2", "amqp://guest:guest@localhost:5672/")
	err := q.Push([]byte("first sample"))
	if err != nil {
		t.Errorf("not passed")
	}
}
