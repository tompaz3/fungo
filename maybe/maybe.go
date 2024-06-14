package maybe

import (
	"errors"
	"reflect"
)

var ErrEmptyMaybe = errors.New("empty maybe")

type Maybe[T any] struct {
	defined bool
	value   T
}

//nolint:unused // sealed interface
func (m *Maybe[T]) sealedMaybe() {}

func (m *Maybe[T]) IsDefined() bool {
	return !m.IsEmpty()
}

func (m *Maybe[T]) IsEmpty() bool {
	return isNil(m) || !m.defined
}

func (m *Maybe[T]) Get() (T, error) {
	if m.IsEmpty() {
		var zero T
		return zero, ErrEmptyMaybe
	}

	return m.value, nil
}

func (m *Maybe[T]) OrZero() T {
	if m.IsEmpty() {
		var zero T
		return zero
	}

	return m.value
}

func (m *Maybe[T]) OrElse(other T) T {
	if m.IsEmpty() {
		return other
	}

	return m.value
}

func (m *Maybe[T]) OrElseGet(other func() T) T {
	if m.IsEmpty() {
		return other()
	}

	return m.value
}

func (m *Maybe[T]) OrElseTryGet(other func() (T, error)) (T, error) {
	if m.IsEmpty() {
		return other()
	}

	return m.value, nil
}

func (m *Maybe[T]) Filter(pred func(T) bool) *Maybe[T] {
	if m.IsEmpty() || !pred(m.OrZero()) {
		return None[T]()
	}
	return m
}
func (m *Maybe[T]) Map(fn func(T) T) *Maybe[T] {
	if m.IsEmpty() {
		return m
	}
	return Some(fn(m.OrZero()))
}
func (m *Maybe[T]) FlatMap(fn func(T) *Maybe[T]) *Maybe[T] {
	if m.IsEmpty() {
		return m
	}
	return fn(m.OrZero())
}

func None[T any]() *Maybe[T] {
	return &Maybe[T]{}
}

func Some[T any](value T) *Maybe[T] {
	return &Maybe[T]{defined: true, value: value}
}

func OfNillable[T any](value T) *Maybe[T] {
	if isNil(value) {
		return None[T]()
	}

	return Some(value)
}

func Map[T, U any](mb *Maybe[T], fn func(T) U) *Maybe[U] {
	if mb.IsEmpty() {
		return None[U]()
	}

	return Some(fn(mb.OrZero()))
}

func TryMap[T, U any](mb *Maybe[T], fn func(T) (U, error)) (*Maybe[U], error) {
	if mb.IsEmpty() {
		return None[U](), nil
	}

	val, err := fn(mb.OrZero())
	if err != nil {
		return None[U](), err
	}

	return Some(val), nil
}

func FlatMap[T, U any](mb *Maybe[T], fn func(T) *Maybe[U]) *Maybe[U] {
	if mb.IsEmpty() {
		return None[U]()
	}

	return fn(mb.OrZero())
}

func TryFlatMap[T, U any](mb *Maybe[T], fn func(T) (*Maybe[U], error)) (*Maybe[U], error) {
	if mb.IsEmpty() {
		return None[U](), nil
	}

	return fn(mb.OrZero())
}

func isNil(value any) bool {
	if value == nil {
		return true
	}

	tp := reflect.TypeOf(value)
	//nolint:exhaustive // we are only interested in the following types and default handles the rest
	switch tp.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(value).IsNil()
	default:
		return false
	}
}
