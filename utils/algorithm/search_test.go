/**
 * @Author: Sean
 * @Date: 2022/5/17 16:47
 */

package algorithm

import (
	"github.com/seanxu24/taco/utils/internal"
	"testing"
)

var sortedNumbers = []int{1, 2, 3, 4, 5, 6, 7, 8}

type intComparator struct{}

func (a intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	//ascending order
	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func TestLinearSearch(t *testing.T) {
	assert := internal.NewAssert(t, "TestLinearSearch")

	comparator := &intComparator{}
	assert.Equal(4, LinearSearch(sortedNumbers, 5, comparator))
	assert.Equal(-1, LinearSearch(sortedNumbers, 9, comparator))
}

func TestBinarySearch(t *testing.T) {
	assert := internal.NewAssert(t, "TestBinarySearch")

	comparator := &intComparator{}
	assert.Equal(4, BinarySearch(sortedNumbers, 5, 0, len(sortedNumbers)-1, comparator))
	assert.Equal(-1, BinarySearch(sortedNumbers, 9, 0, len(sortedNumbers)-1, comparator))
}

func TestBinaryIterativeSearch(t *testing.T) {
	assert := internal.NewAssert(t, "TestBinaryIterativeSearch")

	comparator := &intComparator{}
	assert.Equal(4, BinaryIterativeSearch(sortedNumbers, 5, 0, len(sortedNumbers)-1, comparator))
	assert.Equal(-1, BinaryIterativeSearch(sortedNumbers, 9, 0, len(sortedNumbers)-1, comparator))
}
