package datastructures_test

import (
	"testing"

	"github.com/ninadingole/all-golang/pkg/datastructures"
	"github.com/stretchr/testify/assert"
)

func Test_LinkedList(t *testing.T) {
	t.Run("should push data to linked list", func(t *testing.T) {
		ll := &datastructures.LinkedList{}

		ll.Push(1)
		ll.Push(2)
		ll.Push(3)

		n := ll.First()

		for {
			t.Log(n.Value())
			n = n.Next()
			if n == nil {
				break
			}
		}
	})

	t.Run("should return first element in the list", func(t *testing.T) {
		ll := &datastructures.LinkedList{}

		ll.Push(1)
		ll.Push(2)
		ll.Push(3)

		n := ll.First()

		assert.Equal(t, 1, n.Value())
	})

	t.Run("should return last element in the list", func(t *testing.T) {
		ll := &datastructures.LinkedList{}

		ll.Push(1)
		ll.Push(2)
		ll.Push(3)

		n := ll.Last()

		assert.Equal(t, 3, n.Value())
	})
}
