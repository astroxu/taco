/**
 * @Author: Sean
 * @Date: 2022/5/17 16:12
 */

// Package algorithm contain some basic algorithm function. eg. sort, search, list,linklist, stack, queue, tree, graph. TODO
package algorithm

import "github.com/seanxu24/taco/utils/constraints"

// Search algorithm see https://github.com/TheAlgorithms/Go/tree/master/search

// LinearSearch Simple linear search algorithm that iterates over all elements of a slice
// If a target is found, the index of the target is returned. Else the function return -1
func LinearSearch[T any](slice []T, target T, comparator constraints.Comparator) int {
	for i, v := range slice {
		if comparator.Compare(v, target) == 0 {
			return i
		}
	}
	return -1
}

// BinarySearch search form target within a sorted slice,recursive call itself.
// If a target is found, the index of the target is returned. Else the function return -1
func BinarySearch[T any](sortedSlice []T, target T, lowIndex, highIndex int, comparator constraints.Comparator) int {
	if highIndex < lowIndex || len(sortedSlice) == 0 {
		return -1
	}

	midIndex := lowIndex + (highIndex-lowIndex)/2
	isMidValGreatTarget := comparator.Compare(sortedSlice[midIndex], target) == 1
	isMidValLessTarget := comparator.Compare(sortedSlice[midIndex], target) == -1

	if isMidValGreatTarget {
		return BinarySearch(sortedSlice, target, lowIndex, midIndex-1, comparator)
	} else if isMidValLessTarget {
		return BinarySearch(sortedSlice, target, midIndex+1, highIndex, comparator)
	}

	return midIndex
}

// BinaryIterativeSearch search for target within a sorted slice.
// If a target is found, the index of the target is returned.Else the function return -1
func BinaryIterativeSearch[T any](sortedSlice []T, target T, lowIndex, highIndex int, comparator constraints.Comparator) int {
	startIndex := lowIndex
	endIndex := highIndex

	var midIndex int
	for startIndex <= endIndex {
		midIndex = startIndex + (endIndex-startIndex)/2
		isMidValGreatTarget := comparator.Compare(sortedSlice[midIndex], target) == 1
		isMidValLessTarget := comparator.Compare(sortedSlice[midIndex], target) == -1

		if isMidValGreatTarget {
			endIndex = midIndex - 1
		} else if isMidValLessTarget {
			startIndex = midIndex + 1
		} else {
			return midIndex
		}
	}

	return -1
}
