package errmatch

import (
	"errors"
)

var ErrNoMatchFound = errors.New("no match found")

type (
	ResultFunc[R any]         func(err error) (R, error)
	MatchFunc[E error, R any] func(err E) (R, error)
	ErrPredicateFunc[E error] func(err error) (E, bool)
)

func Switch[R any](err error, fns ...ResultFunc[R]) (R, error) {
	for _, fn := range fns {
		if fn != nil {
			res, resErr := fn(err)
			if errors.Is(resErr, ErrNoMatchFound) {
				continue
			}

			return res, resErr
		}
	}

	var empty R

	return empty, ErrNoMatchFound
}

func Case[E error, R any](fn MatchFunc[E, R]) ResultFunc[R] {
	return func(err error) (R, error) {
		var theErr E
		if errors.As(err, &theErr) {
			return fn(theErr)
		}

		var empty R

		return empty, ErrNoMatchFound
	}
}

func Default[R any](fn ResultFunc[R]) ResultFunc[R] {
	return fn
}

func CaseMatches[E error, R any](pred ErrPredicate[E], fn MatchFunc[E, R]) ResultFunc[R] {
	return func(err error) (R, error) {
		theErr, ok := pred.Test(err)
		if ok {
			return fn(theErr)
		}

		var empty R

		return empty, ErrNoMatchFound
	}
}

func Type[E error]() ErrPredicate[E] {
	return ErrPredicate[E]{
		pred: func(err error) (E, bool) {
			var theErr E
			if errors.As(err, &theErr) {
				return theErr, true
			}

			return theErr, false
		},
	}
}

type ErrPredicate[E error] struct {
	pred ErrPredicateFunc[E]
}

func (p ErrPredicate[E]) Test(err error) (E, bool) {
	return p.pred(err)
}

func (p ErrPredicate[E]) And(pred ErrPredicateFunc[E]) ErrPredicate[E] {
	return ErrPredicate[E]{
		pred: func(err error) (E, bool) {
			theErr, ok := p.pred(err)
			if !ok {
				return theErr, false
			}

			return pred(err)
		},
	}
}

func (p ErrPredicate[E]) Or(pred ErrPredicateFunc[E]) ErrPredicate[E] {
	return ErrPredicate[E]{
		pred: func(err error) (E, bool) {
			theErr, ok := p.pred(err)
			if ok {
				return theErr, true
			}

			return pred(err)
		},
	}
}
