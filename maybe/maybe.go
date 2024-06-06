package maybe

import "errors"

type Maybe[T any] interface {
	sealedMaybe()

	IsDefined() bool
	IsEmpty() bool

	Get() (T, error)
	MustGet() T
	OrElse(other T) T
	OrElseGet(other func() T) T
	OrElseTryGet(other func() (T, error)) (T, error)
}

func Map[T, U any](mb Maybe[T], fn func(T) U) Maybe[U] {
	if mb.IsEmpty() {
		return None[U]()
	}

	return Some(fn(mb.MustGet()))
}

func TryMap[T, U any](mb Maybe[T], fn func(T) (U, error)) (Maybe[U], error) {
	if mb.IsEmpty() {
		return None[U](), nil
	}

	val, err := fn(mb.MustGet())
	if err != nil {
		return None[U](), err
	}

	return Some(val), nil
}

func FlatMap[T, U any](mb Maybe[T], fn func(T) Maybe[U]) Maybe[U] {
	if mb.IsEmpty() {
		return None[U]()
	}

	return fn(mb.MustGet())
}

func TryFlatMap[T, U any](mb Maybe[T], fn func(T) (Maybe[U], error)) (Maybe[U], error) {
	if mb.IsEmpty() {
		return None[U](), nil
	}

	return fn(mb.MustGet())
}

var ErrEmptyMaybe = errors.New("empty maybe")

func None[T any]() Maybe[T] {
	return none[T]{}
}

func Some[T any](value T) Maybe[T] {
	return some[T]{value: value}
}

type none[T any] struct{}

type some[T any] struct {
	value T
}

var (
	_ Maybe[any] = (*none[any])(nil)
	_ Maybe[any] = (*some[any])(nil)
)

//nolint:unused // This is a sealed interface
func (none[T]) sealedMaybe() {}

//nolint:unused // This is a sealed interface
func (some[T]) sealedMaybe() {}

func (none[T]) Get() (T, error) {
	var empty T

	return empty, ErrEmptyMaybe
}

func (none[T]) IsDefined() bool {
	return false
}

func (s some[T]) IsDefined() bool {
	return true
}

func (none[T]) IsEmpty() bool {
	return true
}

func (s some[T]) IsEmpty() bool {
	return false
}

func (s some[T]) Get() (T, error) {
	return s.value, nil
}

func (none[T]) MustGet() T {
	panic(ErrEmptyMaybe)
}

func (s some[T]) MustGet() T {
	return s.value
}

func (none[T]) OrElse(other T) T {
	return other
}

func (s some[T]) OrElse(_ T) T {
	return s.value
}

func (none[T]) OrElseGet(other func() T) T {
	return other()
}

func (s some[T]) OrElseGet(_ func() T) T {
	return s.value
}

func (none[T]) OrElseTryGet(other func() (T, error)) (T, error) {
	return other()
}

func (s some[T]) OrElseTryGet(_ func() (T, error)) (T, error) {
	return s.value, nil
}
