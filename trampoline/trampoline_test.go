/*
Copyright (c) 2024-2025 Tomasz Paździurek

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package trampoline_test

import (
	"fmt"

	"github.com/tompaz3/fungo/trampoline"
)

type LinkedList[T any] interface {
	Value() T
	Size() int
	Next() LinkedList[T]
	Append(value T) LinkedList[T]
}

type Result[T any] struct {
	Value T
	Valid bool
}

func Example_iterateLinkedListRecursively() {
	size := 100_000_000
	list := LinkedListEmpty[int]()
	for i := range size {
		list = list.Append(i + 1)
	}

	res := iterateToLast(list).Execute()
	fmt.Printf("Valid: %t, Value: %d", res.Valid, res.Value)
	// Output: Valid: true, Value: 1
}

func LinkedListEmpty[T any]() LinkedList[T] {
	return linkedList[T]{}
}

type linkedList[T any] struct {
	value T
	size  int
	next  LinkedList[T]
}

func (l linkedList[T]) Size() int {
	return l.size
}

func (l linkedList[T]) Value() T {
	return l.value
}

func (l linkedList[T]) Next() LinkedList[T] {
	return l.next
}

func (l linkedList[T]) Append(value T) LinkedList[T] {
	if l.Size() == 0 {
		return linkedList[T]{value: value, size: 1}
	}

	return linkedList[T]{value: value, size: l.size + 1, next: l}
}

func iterateToLast[T any](list LinkedList[T]) trampoline.Trampoline[Result[T]] {
	if list.Size() == 0 {
		return trampoline.Complete[Result[T]](Result[T]{})
	}

	if list.Size() == 1 {
		return trampoline.Complete[Result[T]](Result[T]{Valid: true, Value: list.Value()})
	}

	return trampoline.Next[Result[T]](func() trampoline.Trampoline[Result[T]] {
		return iterateToLast[T](list.Next())
	})
}
