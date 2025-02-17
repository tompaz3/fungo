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

package trampoline

type TailCall[T any] func() Trampoline[T]

type Trampoline[T any] interface {
	Execute() T

	next() TailCall[T]
	done() bool
	result() T
}

type trampolineImpl[T any] struct {
	nextCall TailCall[T]
	finished bool
	value    T
}

//nolint:unused // it's actually used by the Execute function.
func (tr trampolineImpl[T]) next() TailCall[T] {
	return tr.nextCall
}

//nolint:unused // it's actually used by the Execute function.
func (tr trampolineImpl[T]) done() bool {
	return tr.finished
}

//nolint:unused // it's actually used by the Execute function.
func (tr trampolineImpl[T]) result() T {
	return tr.value
}

func (tr trampolineImpl[T]) Execute() T {
	var trmp Trampoline[T] = tr
	for !trmp.done() {
		trmp = trmp.next()()
	}

	return trmp.result()
}

func Complete[T any](result T) Trampoline[T] {
	return trampolineImpl[T]{finished: true, value: result}
}

func Next[T any](nextCall TailCall[T]) Trampoline[T] {
	return trampolineImpl[T]{nextCall: nextCall}
}
