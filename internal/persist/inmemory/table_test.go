package inmemory

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func newInMemoryTable[KeyType constraints.Ordered, ItemType any](itemName string) table[KeyType, ItemType] {
	return table[KeyType, ItemType]{
		name:  itemName,
		items: map[KeyType]ItemType{},
	}
}

type testItem struct {
	field1 string
	field2 int
}

func TestInMemory_NewListShouldHaveZeroLength(t *testing.T) {
	items := newInMemoryTable[int16, testItem]("Test")
	want := 0
	got := items.Count()
	if want != got {
		t.Fatalf(`ItemList count = %v. Expected %v`, got, want)
	}
}

func TestInMemory_AddingANewKeyShouldIncreaseCount(t *testing.T) {
	items := newInMemoryTable[int16, testItem]("Test")
	err := items.AddNew(1, &testItem{
		field1: "foo",
		field2: 2,
	})
	want := 1
	got := items.Count()
	if want != got {
		t.Fatalf(`ItemList count = %v. Expected %v`, got, want)
	}
	if err != nil {
		t.Fatalf(`Received error while adding first new item`)
	}

	err = items.AddNew(11, &testItem{
		field1: "bar",
		field2: 22,
	})
	want = 2
	got = items.Count()
	if want != got {
		t.Fatalf(`ItemList count = %v. Expected %v`, got, want)
	}
	if err != nil {
		t.Fatalf(`Received error while adding second new item`)
	}
}

func TestInMemory_AddingTheSameKeyAgainShouldFail(t *testing.T) {
	items := newInMemoryTable[int16, testItem]("Test")
	err := items.AddNew(1, &testItem{
		field1: "foo",
		field2: 2,
	})
	if err != nil {
		t.Fatalf(`Received error while adding first new item`)
	}
	err = items.AddNew(1, &testItem{
		field1: "foo",
		field2: 22,
	})
	want := 1
	got := items.Count()
	if want != got {
		t.Fatalf(`ItemList count = %v. Expected %v`, got, want)
	}
	if err == nil {
		t.Fatalf(`Should have received error while adding item with same key`)
	}
}

func TestInMemory_FindingExistingItemShouldSucceed(t *testing.T) {
	items := newInMemoryTable[int16, testItem]("Test")
	items.AddNew(1, &testItem{
		field1: "foo",
		field2: 2,
	})
	items.AddNew(11, &testItem{
		field1: "bar",
		field2: 22,
	})
	items.AddNew(111, &testItem{
		field1: "baz",
		field2: 222,
	})
	got, err := items.Find(11)
	want := testItem{
		field1: "bar",
		field2: 22,
	}
	if err != nil {
		t.Fatalf(`Received error while finding an existing item`)
	}
	if got.field1 != want.field1 || got.field2 != want.field2 {
		t.Fatalf(`Found wrong item: %+v Expected: %+v`, got, want)
	}
}

func TestInMemory_FindingNonExistingItemShouldReturnWithError(t *testing.T) {
	items := newInMemoryTable[int16, testItem]("Test")
	items.AddNew(1, &testItem{
		field1: "foo",
		field2: 2,
	})
	items.AddNew(11, &testItem{
		field1: "bar",
		field2: 22,
	})
	items.AddNew(111, &testItem{
		field1: "baz",
		field2: 222,
	})
	_, err := items.Find(2)
	if err == nil {
		t.Fatalf(`Should have received error while finding a non-existing item`)
	}
}

func TestInMemory_ShouldStoreCopy(t *testing.T) {
	items := newInMemoryTable[int16, testItem]("Test")
	item := &testItem{
		field1: "foo",
		field2: 2,
	}
	items.AddNew(1, item)
	item.field2 = 3

	storedItem, _ := items.Find(1)
	want := 2
	got := storedItem.field2
	if want != got {
		t.Fatalf(`Wring item value: %v. Expected %v`, got, want)
	}
}

func TestInMemory_FindShouldReturnWithCopy(t *testing.T) {
	items := newInMemoryTable[int16, testItem]("Test")
	items.AddNew(1, &testItem{
		field1: "foo",
		field2: 2,
	})
	storedItem1, _ := items.Find(1)
	storedItem2, _ := items.Find(1)
	storedItem1.field2 = 3
	want := 2
	got := storedItem2.field2
	if want != got {
		t.Fatalf(`Wring item value: %v. Expected %v`, got, want)
	}
}
