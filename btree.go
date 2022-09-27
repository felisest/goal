package goal

import (
	"golang.org/x/exp/constraints"
)

type Tree[K constraints.Ordered, V any] struct {
	Key   K
	Value V

	left   *Tree[K, V]
	right  *Tree[K, V]
	parent *Tree[K, V]
}

func NewTree[K constraints.Ordered, V any](k K, v V) (*Tree[K, V], error) {

	return &Tree[K, V]{parent: nil, Key: k, Value: v}, nil
}

func (t *Tree[K, V]) Insert(key K, value V) {

	if t.Key == key {
		t.Value = value

	} else if t.Key > key {

		if t.left == nil {
			t.left = &Tree[K, V]{Key: key, Value: value, parent: t}
		} else {
			t.left.Insert(key, value)
		}

	} else if t.Key < key {

		if t.right == nil {
			t.right = &Tree[K, V]{Key: key, Value: value, parent: t}
		} else {
			t.right.Insert(key, value)
		}
	}
}

func (t *Tree[K, V]) Find(key K) (V, bool) {

	target_tree := find(key, t)

	if target_tree != nil {
		return target_tree.Value, true

	} else {
		var v V
		return v, false
	}
}

func (t *Tree[K, V]) Remove(key K) bool {

	remove_element := find(key, t)
	if remove_element == nil {
		return false
	}

	return remove(remove_element)
}

func remove[K constraints.Ordered, V any](remove_element *Tree[K, V]) bool {

	if remove_element.left == nil && remove_element.right == nil {

		if remove_element.parent.left == remove_element {
			remove_element.parent.left = nil
		} else if remove_element.parent.right == remove_element {
			remove_element.parent.right = nil
		}

	} else if remove_element.left == nil || remove_element.right == nil {

		if remove_element.left != nil {

			if remove_element.parent.left == remove_element {
				remove_element.parent.left = remove_element.left
			} else if remove_element.parent.right == remove_element {
				remove_element.parent.right = remove_element.left
			}
			remove_element.left.parent = remove_element.parent
		}

		if remove_element.right != nil {

			if remove_element.parent.left == remove_element {
				remove_element.parent.left = remove_element.right
			} else if remove_element.parent.right == remove_element {
				remove_element.parent.right = remove_element.right
			}
			remove_element.right.parent = remove_element.parent
		}

	} else {

		if remove_element.right.left == nil {

			remove_element.Key = remove_element.right.Key
			remove_element.Value = remove_element.right.Value
			remove_element.right = remove_element.right.right

		} else {

			min_element := find_min(remove_element.right)

			remove_element.Key = min_element.Key
			remove_element.Value = min_element.Value

			if min_element.right != nil {
				min_element.right.parent = min_element.parent
				min_element.parent.left = min_element.right
			} else {
				min_element.parent.left = nil
			}
		}
	}

	return false
}

func find_min[K constraints.Ordered, V any](root_tree *Tree[K, V]) *Tree[K, V] {

	if root_tree.left != nil {
		return find_min(root_tree.left)
	} else {
		return root_tree
	}
}

func find_max[K constraints.Ordered, V any](root_tree *Tree[K, V]) *Tree[K, V] {

	if root_tree.right != nil {
		return find_max(root_tree.right)
	} else {
		return root_tree
	}
}

func find[K constraints.Ordered, V any](key K, root_tree *Tree[K, V]) *Tree[K, V] {

	if root_tree.Key == key {
		return root_tree
	}

	if root_tree.Key > key && root_tree.left != nil {

		return find(key, root_tree.left)

	} else if root_tree.Key < key && root_tree.right != nil {

		return find(key, root_tree.right)
	}

	return nil
}
