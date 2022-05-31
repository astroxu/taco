/**
 * @Author: Sean
 * @Date: 2022/5/31 10:21
 */

// Package datastructure implements some data structure. eg. list, linklist, stack, queue, tree, graph.
package datastructure

// LinkNode is a linkedlist node, which have a Value and Pre points to previous node, Next points to a next node of the link.
type LinkNode[T any] struct {
	Value T
	pre   *LinkNode[T]
	Next  *LinkNode[T]
}

// NewLinkNode return a LinkNode pointer
func NewLinkNode[T any](value T) *LinkNode[T] {
	return &LinkNode[T]{value, nil, nil}
}

// StackNode is node in stack,which have a Value and Next pointer points to next node in the stack.
type StackNode[T any] struct {
	Value T
	Next  *StackNode[T]
}

// NewStackNode return a StackNode pointer
func NewStackNode[T any](value T) *StackNode[T] {
	return &StackNode[T]{value, nil}
}

// QueueNode is a node in a queue,which a have a Value and Next pointer points to next node in the queue.
type QueueNode[T any] struct {
	Value T
	Next  *QueueNode[T]
}

// NewQueueNode return a QueueNode pointer
func NewQueueNode[T any](value T) *QueueNode[T] {
	return &QueueNode[T]{value, nil}
}

// TreeNode is node of tree
type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// NewTreeNode return a TreeNode pointer
func NewTreeNode[T any](value T) *TreeNode[T] {
	return &TreeNode[T]{value, nil, nil}
}