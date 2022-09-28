package goal

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBTreeRootInsert(t *testing.T) {

	tree, err := MakeTree(10, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	value, ok := tree.Find(10)
	if value != "Root" && !ok {
		t.Errorf("insert error")
	}

	tree.Insert(10, "New root value")

	value, ok = tree.Find(10)
	if value != "New root value" && !ok {
		t.Errorf("reinsert error")
	}
}

func TestBTreeFind(t *testing.T) {

	tree, err := MakeTree(10, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	tree.Insert(14, "Right tree - level #1")
	tree.Insert(16, "Right tree - right tree - level #2")
	tree.Insert(12, "Right tree - left tree - level #2")

	value, ok := tree.Find(12)
	if value != "Right tree - left tree - level #2" && ok {
		t.Errorf("find error")
	}

	value, ok = tree.Find(14)
	if value != "Right tree - level #1" && ok {
		t.Errorf("find error")
	}

	value, ok = tree.Find(16)
	if value != "Right tree - right tree - level #2" && ok {
		t.Errorf("find error")
	}
}

func TestBTreeFindNonexistent(t *testing.T) {

	tree, err := MakeTree(10, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	tree.Insert(14, "Right tree - level #1")
	tree.Insert(16, "Right tree - right tree - level #2")
	tree.Insert(12, "Right tree - left tree - level #2")

	value, ok := tree.Find(13)
	if value != "" && !ok {
		t.Errorf("find non-existent")
	}
}

func TestMassInsertFind(t *testing.T) {

	const MAX_KEY_VALUE = 10000
	const MAX_ARRAY_LEN = 10000

	rand.Seed(time.Now().UnixNano())

	tree, err := MakeTree(rand.Intn(MAX_KEY_VALUE), "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	keys := make([]int, MAX_ARRAY_LEN)

	for i := 0; i < MAX_ARRAY_LEN; i++ {
		keys[i] = rand.Intn(MAX_KEY_VALUE)
	}

	for _, key := range keys {

		value := fmt.Sprintf("Value of Key = %d", key)
		tree.Insert(key, value)
	}

	// Shuffle keys slice
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(keys), func(i, j int) { keys[i], keys[j] = keys[j], keys[i] })

	// Find test
	for _, key := range keys {
		value := fmt.Sprintf("Value of Key = %d", key)
		tree_value, ok := tree.Find(key)

		if tree_value != value && !ok {
			t.Errorf("find() error key: %d not exists", key)
		}
	}
}

func TestMassInsertRemove(t *testing.T) {

	const MAX_KEY_VALUE = 1000
	const MAX_ARRAY_LEN = 200000
	const CNT_KEYS_REMOVE = 3

	rand.Seed(time.Now().UnixNano())

	tree, err := MakeTree(rand.Intn(MAX_KEY_VALUE), "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	keys_map := make(map[int]struct{}, MAX_ARRAY_LEN)

	for i := 0; i < MAX_ARRAY_LEN; i++ {
		key := rand.Intn(MAX_KEY_VALUE)
		keys_map[key] = struct{}{}
	}

	for key := range keys_map {

		value := fmt.Sprintf("Value of Key = %d", key)
		tree.Insert(key, value)
	}

	//Remove a few keys
	cnt := 0
	deleted_keys := make([]int, CNT_KEYS_REMOVE)
	for key := range keys_map {
		if key == tree.Key {
			continue
		}
		tree.Remove(key)
		delete(keys_map, key)
		deleted_keys[cnt] = key
		fmt.Printf("key = %d", key)
		if cnt >= CNT_KEYS_REMOVE-1 {
			break
		}
		cnt++
	}

	for key := range keys_map {

		value := fmt.Sprintf("Value of Key = %d", key)
		tree_value, ok := tree.Find(key)

		if tree_value != value && !ok {
			t.Errorf("find() error key: %d is not exists", key)
		}
	}

	for _, key := range deleted_keys {
		tree_value, ok := tree.Find(key)
		_ = tree_value
		if ok {
			t.Errorf("find() error key: %d is exists", key)
		}
	}
}

func TestBTreeRemoveLRNil(t *testing.T) {

	tree, err := MakeTree(20, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	tree.Insert(10, "Key = 10")
	tree.Insert(5, "Key = 5")
	tree.Insert(15, "Key = 2")

	tree.Remove(15)
	tree.Remove(5)

	_, ok := tree.Find(5)
	if ok {
		t.Errorf("'right-left nil' error key: %d is exists", 5)
	}

	_, ok = tree.Find(15)
	if ok {
		t.Errorf("'right-left nil' error key: %d is exists", 15)
	}

}

func TestBTreeRemoveLeftLeaf(t *testing.T) {

	tree, err := MakeTree(10, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	tree.Insert(5, "Key = 5")
	tree.Insert(2, "Key = 2")

	tree.Remove(5)

	_, ok := tree.Find(5)
	if ok {
		t.Errorf("'left leaf' error key: %d is exists", 5)
	}
	_, ok = tree.Find(2)
	if !ok {
		t.Errorf("'left leaf' error key: %d is not exists", 2)
	}

	tree, err = MakeTree(10, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	// right nil
	tree.Insert(14, "Key = 14")
	tree.Insert(12, "Key = 12")

	tree.Remove(14)

	_, ok = tree.Find(14)
	if ok {
		t.Errorf("'right nil' error key: %d is exists", 5)
	}
	_, ok = tree.Find(12)
	if !ok {
		t.Errorf("'right nil' error key: %d is not exists", 12)
	}
}

func TestBTreeRemoveRightLeaf(t *testing.T) {

	tree, err := MakeTree(10, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	// right nil
	tree.Insert(5, "Key = 5")
	tree.Insert(7, "Key = 7")

	tree.Remove(5)

	_, ok := tree.Find(5)
	if ok {
		t.Errorf("'right leaf' error key: %d is exists", 5)
	}
	_, ok = tree.Find(7)
	if !ok {
		t.Errorf("'right leaf' error key: %d is not exists", 7)
	}

	tree, err = MakeTree(10, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	// right nil
	tree.Insert(14, "Key = 14")
	tree.Insert(16, "Key = 16")

	tree.Remove(14)

	_, ok = tree.Find(14)
	if ok {
		t.Errorf("'right nil' error key: %d is exists", 14)
	}
	_, ok = tree.Find(16)
	if !ok {
		t.Errorf("'right nil' error key: %d is not exists", 16)
	}
}

func TestBTreeRemoveBothLeafs(t *testing.T) {

	tree, err := MakeTree(10, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	tree.Insert(5, "Key = 5")
	tree.Insert(2, "Key = 2")
	tree.Insert(7, "Key = 7")

	tree.Remove(5)

	_, ok := tree.Find(5)
	if ok {
		t.Errorf("'both leafs' error key: %d is exists", 5)
	}
	_, ok = tree.Find(2)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 2)
	}
	_, ok = tree.Find(7)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 7)
	}

	tree, err = MakeTree(10, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	tree.Insert(14, "Key = 14")
	tree.Insert(12, "Key = 12")
	tree.Insert(16, "Key = 16")

	tree.Remove(14)

	_, ok = tree.Find(14)
	if ok {
		t.Errorf("'right nil' error key: %d is exists", 14)
	}
	_, ok = tree.Find(12)
	if !ok {
		t.Errorf("'right nil' error key: %d is not exists", 16)
	}
	_, ok = tree.Find(16)
	if !ok {
		t.Errorf("'right nil' error key: %d is not exists", 16)
	}
}

func TestBTreeRemoveBothLeafsRightLeft(t *testing.T) {

	tree, err := MakeTree(10, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	tree.Insert(2, "Key = 2")
	tree.Insert(1, "Key = 1")
	tree.Insert(8, "Key = 8")
	tree.Insert(9, "Key = 9")
	tree.Insert(7, "Key = 7")
	tree.Insert(5, "Key = 5")
	tree.Insert(3, "Key = 3")
	tree.Insert(4, "Key = 4")

	tree.Remove(2)

	_, ok := tree.Find(2)
	if ok {
		t.Errorf("'both leafs' error key: %d is exists", 5)
	}
	_, ok = tree.Find(8)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 8)
	}
	_, ok = tree.Find(9)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 9)
	}
	_, ok = tree.Find(7)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 7)
	}
	_, ok = tree.Find(5)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 5)
	}
	_, ok = tree.Find(3)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 3)
	}
	_, ok = tree.Find(4)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 4)
	}
	_, ok = tree.Find(1)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 1)
	}

	tree, err = MakeTree(20, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	tree.Insert(15, "Key = 15")
	tree.Insert(12, "Key = 12")
	tree.Insert(20, "Key = 20")
	tree.Insert(18, "Key = 18")
	tree.Insert(16, "Key = 16")
	tree.Insert(17, "Key = 17")

	tree.Remove(15)

	_, ok = tree.Find(15)
	if ok {
		t.Errorf("'both leafs' error key: %d is exists", 15)
	}
	_, ok = tree.Find(12)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 12)
	}
	_, ok = tree.Find(20)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 20)
	}
	_, ok = tree.Find(18)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 18)
	}
	_, ok = tree.Find(16)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 16)
	}
	_, ok = tree.Find(17)
	if !ok {
		t.Errorf("'both leafs' error key: %d is not exists", 17)
	}

}

func TestBTreeRemoveCustom(t *testing.T) {

	tree, err := MakeTree(723, "Root")
	if err != nil {
		t.Errorf("Can`t create new tree")
	}

	tree.Insert(541, "Key = 541")
	tree.Insert(576, "Key = 576")
	tree.Insert(382, "Key = 382")
	tree.Insert(576, "Key = 576")
	tree.Insert(985, "Key = 985")
	tree.Insert(1000, "Key = 1000")

	tree.Remove(723)

	str, ok := tree.Find(723)
	_, _ = str, ok
}
