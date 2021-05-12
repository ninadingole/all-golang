package datastructures_test

import (
	"testing"

	"github.com/ninadingole/all-golang/pkg/datastructures"
	"github.com/stretchr/testify/assert"
)

func Test_BinaryTree(t *testing.T) {
	t.Run("should insert values to tree", func(t *testing.T) {
		tree := &datastructures.Tree{}

		tree.Insert(10).Insert(8).Insert(25)

		datastructures.PrintTree(tree)
	})

	t.Run("should return true when value exists in tree", func(t *testing.T) {
		tree := &datastructures.Tree{}

		tree.Insert(10).Insert(8).Insert(25)

		assert.True(t, tree.Exists(25))
	})

	t.Run("should return false when value exists not exists in tree", func(t *testing.T) {
		tree := &datastructures.Tree{}

		tree.Insert(10).Insert(8).Insert(25)

		assert.False(t, tree.Exists(100))
	})
}
