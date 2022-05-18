/**
 * @Author: Sean
 * @Date: 2022/5/18 16:16
 */

package algorithm

import (
	"fmt"
	"github.com/seanxu24/taco/utils/internal"
	"testing"
)

// People test mock data
type people struct {
	Name string
	Age  int
}

// PeopleAgeComparator sort people slice by age field
type PeopleAgeComparator struct{}

// Compare implements github.com/seanxu24/tacor/utils/constraints/constraints.go/Comparator
func (pc *PeopleAgeComparator) Compare(v1, v2 any) int {
	p1, _ := v1.(people)
	p2, _ := v2.(people)

	// ascending order
	if p1.Age < p2.Age {
		return -1
	} else if p1.Age > p2.Age {
		return 1
	}
	return 0
}

//var peoples = []people{
//	{Name:"a",Age: 20},
//	{Name:"b",Age: 10},
//	{Name:"c",Age: 17},
//	{Name:"d",Age: 8},
//	{Name:"e",Age: 28},
//}

func TestBubbleSortForStructSlice(t *testing.T) {
	assert := internal.NewAssert(t, "TestBubbleSortForStructSlice")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &PeopleAgeComparator{}
	BubbleSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	assert.Equal(expected, actual)
}

func TestBubbleSortForIntSlices(t *testing.T) {
	assert := internal.NewAssert(t, "TestBubbleSortForIntSlices")
	numbers := []int{2, 1, 5, 3, 6, 4}
	comparator := &intComparator{}
	BubbleSort(numbers, comparator)

	assert.Equal([]int{1, 2, 3, 4, 5, 6}, numbers)
}

func TestInsertionSort(t *testing.T) {
	assert := internal.NewAssert(t, "TestInsertionSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &PeopleAgeComparator{}
	InsertionSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	assert.Equal(expected, actual)
}

func TestSelectionSort(t *testing.T) {
	assert := internal.NewAssert(t, "TestSelectionSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &PeopleAgeComparator{}
	SelectionSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	assert.Equal(expected, actual)
}

func TestShellSort(t *testing.T) {
	assert := internal.NewAssert(t, "TestShellSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &PeopleAgeComparator{}
	ShellSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	assert.Equal(expected, actual)
}

func TestQuickSort(t *testing.T) {
	assert := internal.NewAssert(t, "TestQuickSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &PeopleAgeComparator{}
	QuickSort(peoples, 0, len(peoples)-1, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	assert.Equal(expected, actual)
}

func TestHeapSort(t *testing.T) {
	assert := internal.NewAssert(t, "TestHeapSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &PeopleAgeComparator{}
	HeapSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	assert.Equal(expected, actual)
}

func TestMergeSort(t *testing.T) {
	assert := internal.NewAssert(t, "TestMergeSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &PeopleAgeComparator{}
	MergeSort(peoples, comparator)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", peoples)

	assert.Equal(expected, actual)
}

func TestCountSort(t *testing.T) {
	assert := internal.NewAssert(t, "TestCountSort")

	peoples := []people{
		{Name: "a", Age: 20},
		{Name: "b", Age: 10},
		{Name: "c", Age: 17},
		{Name: "d", Age: 8},
		{Name: "e", Age: 28},
	}
	comparator := &PeopleAgeComparator{}
	sortedPeopleByAge := CountSort(peoples, comparator)
	t.Log(sortedPeopleByAge)

	expected := "[{d 8} {b 10} {c 17} {a 20} {e 28}]"
	actual := fmt.Sprintf("%v", sortedPeopleByAge)

	assert.Equal(expected, actual)
}
