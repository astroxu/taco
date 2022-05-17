/**
 * @Author: Sean
 * @Date: 2022/5/17 16:14
 */

// Package constraints contain some customer constraints.
package constraints

// Comparator is for comparing two value
type Comparator interface {
	// Compare v1 and v2
	// Ascending order: should return 1 -> v1 > v2, 0 -> v1 = v2,-1 -> v1 < v2
	// Descending order: should return 1 -> v1 < v2, 0 -> v1 = v2,-1 -> v1 > v2
	Compare(v1, v2 any) int
}

// Number contains all types of number and uintptr, used for generics constraint
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}
