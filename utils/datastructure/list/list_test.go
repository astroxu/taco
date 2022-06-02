/**
 * @Author: Sean
 * @Date: 2022/6/2 09:53
 */

package datastructure

import (
	"github.com/seanxu24/taco/utils/internal"
	"testing"
)

func TestNewList(t *testing.T) {
	assert := internal.NewAssert(t, "TestNewList")

	list := NewList([]int{1, 2, 3})
	assert.Equal([]int{1, 2, 3}, list.Data())
}

func TestList_ValueOf(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_ValueOf")

	list := NewList([]int{1, 2, 3})
	v, ok := list.ValueOf(0)
	assert.Equal(1, *v)
	assert.Equal(true, ok)

	_, ok = list.ValueOf(3)
	assert.Equal(false, ok)
}

func TestList_IndexOf(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_IndexOf")

	list := NewList([]int{1, 2, 3})
	i := list.IndexOf(1)
	assert.Equal(0, i)

	i = list.IndexOf(4)
	assert.Equal(-1, i)
}

func TestList_Contain(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_Contain")

	list := NewList([]int{1, 2, 3})
	assert.Equal(true, list.Contain(1))
	assert.Equal(false, list.Contain(0))
}

func TestList_Push(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_Push")

	list := NewList([]int{1, 2, 3})
	list.Push(4)

	assert.Equal([]int{1, 2, 3, 4}, list.Data())
}

func TestList_InsertAtFirst(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_InsertAtFirst")

	list := NewList([]int{1, 2, 3})
	list.InsertAtFirst(0)

	assert.Equal([]int{0, 1, 2, 3}, list.Data())
}

func TestList_InsertAtLast(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_InsertAtLast")

	list := NewList([]int{1, 2, 3})
	list.InsertAtLast(4)

	assert.Equal([]int{1, 2, 3, 4}, list.Data())
}

func TestList_InsertAt(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_InsertAt")

	list := NewList([]int{1, 2, 3})

	list.InsertAt(-1, 0)
	assert.Equal([]int{1, 2, 3}, list.Data())

	list.InsertAt(4, 0)
	assert.Equal([]int{1, 2, 3}, list.Data())

	list.InsertAt(0, 0)
	assert.Equal([]int{0, 1, 2, 3}, list.Data())

	list.InsertAt(4, 4)
	assert.Equal([]int{0, 1, 2, 3, 4}, list.Data())
}

func TestList_PopFirst(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_PopFirst")

	list := NewList([]int{1, 2, 3})
	v, ok := list.PopFirst()
	assert.Equal(1, *v)
	assert.Equal(true, ok)
	assert.Equal([]int{2, 3}, list.Data())

	list2 := NewList([]int{})
	_, ok = list2.PopFirst()
	assert.Equal(false, ok)
	assert.Equal([]int{}, list2.Data())
}

func TestList_PopLast(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_PopLast")

	list := NewList([]int{1, 2, 3})
	v, ok := list.PopLast()
	assert.Equal(3, *v)
	assert.Equal(true, ok)
	assert.Equal([]int{1, 2}, list.Data())

	list2 := NewList([]int{})
	_, ok = list2.PopLast()
	assert.Equal(false, ok)
	assert.Equal([]int{}, list2.Data())
}

func TestList_DeleteAt(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_DeleteAt")

	list := NewList([]int{1, 2, 3, 4})

	list.DeleteAt(-1)
	assert.Equal([]int{1, 2, 3, 4}, list.Data())

	list.DeleteAt(4)
	assert.Equal([]int{1, 2, 3, 4}, list.Data())

	list.DeleteAt(0)
	assert.Equal([]int{2, 3, 4}, list.Data())

	list.DeleteAt(2)
	assert.Equal([]int{2, 3}, list.Data())
}

func TestList_UpdateAt(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_UpdateAt")

	list := NewList([]int{1, 2, 3, 4})

	list.UpdateAt(-1, 0)
	assert.Equal([]int{1, 2, 3, 4}, list.Data())

	list.UpdateAt(4, 0)
	assert.Equal([]int{1, 2, 3, 4}, list.Data())

	list.UpdateAt(0, 5)
	assert.Equal([]int{5, 2, 3, 4}, list.Data())

	list.UpdateAt(3, 1)
	assert.Equal([]int{5, 2, 3, 1}, list.Data())
}

func TestList_Equal(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_Equal")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{1, 2, 3, 4})
	list3 := NewList([]int{1, 2, 3})

	assert.Equal(true, list1.Equal(list2))
	assert.Equal(false, list1.Equal(list3))
}

func TestList_IsEmpty(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_IsEmpty")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{})

	assert.Equal(false, list1.IsEmpty())
	assert.Equal(true, list2.IsEmpty())
}

func TestList_Clear(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_Clear")

	list1 := NewList([]int{1, 2, 3, 4})
	list1.Clear()
	empty := NewList([]int{})

	assert.Equal(empty, list1)
}

func TestList_Clone(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_Clone")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := list1.Clone()

	assert.Equal(true, list1.Equal(list2))
}

func TestList_Merge(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_Merge")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{4, 5, 6})
	expected := NewList([]int{1, 2, 3, 4, 4, 5, 6})

	list3 := list1.Merge(list2)
	assert.Equal(true, expected.Equal(list3))
}

func TestList_Size(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_Size")

	list1 := NewList([]int{1, 2, 3, 4})
	empty := NewList([]int{})

	assert.Equal(4, list1.Size())
	assert.Equal(0, empty.Size())
}

func TestList_Swap(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_Swap")

	list := NewList([]int{1, 2, 3, 4})
	expected := NewList([]int{4, 2, 3, 1})

	list.Swap(0, 3)

	assert.Equal(true, expected.Equal(list))
}

func TestList_Reverse(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_Reverse")

	list := NewList([]int{1, 2, 3, 4})
	expected := NewList([]int{4, 3, 2, 1})

	list.Reverse()

	assert.Equal(true, expected.Equal(list))
}

func TestList_Unique(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_Unique")

	list := NewList([]int{1, 2, 2, 3, 4})
	expected := NewList([]int{1, 2, 3, 4})

	list.Unique()

	assert.Equal(true, expected.Equal(list))
}

func TestList_Union(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_Union")

	list1 := NewList([]int{1, 2, 2, 3, 4})
	list2 := NewList([]int{4, 5, 6})
	expected := NewList([]int{1, 2, 3, 4, 5, 6})

	list3 := list1.Union(list2)

	assert.Equal(true, expected.Equal(list3))
}

func TestList_Intersection(t *testing.T) {
	assert := internal.NewAssert(t, "TestList_Intersection")

	list1 := NewList([]int{1, 2, 3, 4})
	list2 := NewList([]int{4, 5, 6})
	expected := NewList([]int{4})

	list3 := list1.Intersection(list2)
	assert.Equal(true, expected.Equal(list3))
}
