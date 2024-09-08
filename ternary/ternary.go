package ternary

type Ternary[T any] interface {
	Then(then T) Then[T]
	ThenF(f func() T) Then[T]
}

type Then[T any] interface {
	Else(other T) T
	ElseF(f func() T) T
}

type ternary[T any] struct {
	condition func() bool
	then      func() T
}

var (
	_ Ternary[any] = (*ternary[any])(nil)
	_ Then[any]    = (*ternary[any])(nil)
)

func (t *ternary[T]) Then(then T) Then[T] {
	t.then = func() T { return then }
	return t
}

func (t *ternary[T]) ThenF(fn func() T) Then[T] {
	t.then = fn
	return t
}

func (t *ternary[T]) Else(other T) T {
	if t.condition() {
		return t.then()
	}
	return other
}

func (t *ternary[T]) ElseF(fn func() T) T {
	if t.condition() {
		return t.then()
	}
	return fn()
}

func If[T any](condition bool) Ternary[T] {
	return &ternary[T]{
		condition: func() bool { return condition },
	}
}

func IfF[T any](condition func() bool) Ternary[T] {
	return &ternary[T]{
		condition: condition,
	}
}
