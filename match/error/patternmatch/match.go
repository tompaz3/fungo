package patternmatch

import (
	"errors"
)

var ErrNoMatchFound = errors.New("no match found")

type (
	ResultFunc[R any]              func(err error) (R, error)
	MatchFunc[E error, R any]      func(err E) (R, error)
	ErrPredicateFunc[E error]      func(err error) (E, bool)
	PredicateFunc[E error]         func(err E) bool
	TypedErrPredicateFunc[E error] func(err E) bool
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

func Default[R any](fn ResultFunc[R]) ResultFunc[R] {
	return fn
}

func Case[E error, R any](pred ErrPredicate[E], fn MatchFunc[E, R]) ResultFunc[R] {
	return func(err error) (R, error) {
		theErr, ok := pred.Test(err)
		if ok {
			return fn(theErr)
		}

		var empty R

		return empty, ErrNoMatchFound
	}
}

func Type[E error]() TypedErrPredicate[E] {
	return TypedErrPredicate[E]{
		typePred: RawErrPredicate[E]{
			pred: func(err error) (E, bool) {
				var theErr E
				if errors.As(err, &theErr) {
					return theErr, true
				}

				return theErr, false
			},
		},
		pred: func(_ E) bool {
			return true
		},
	}
}

func CaseType[E error, R any](fn MatchFunc[E, R]) ResultFunc[R] {
	return Case[E, R](Type[E](), fn)
}

func Matches(pred PredicateFunc[error]) ErrPredicate[error] {
	return RawErrPredicate[error]{
		pred: func(err error) (error, bool) {
			return err, pred(err)
		},
	}
}

type ErrPredicate[E error] interface {
	Test(err error) (E, bool)
}

type RawErrPredicate[E error] struct {
	pred ErrPredicateFunc[E]
}

func (p RawErrPredicate[E]) Test(err error) (E, bool) {
	return p.pred(err)
}

func (p RawErrPredicate[E]) And(pred ErrPredicateFunc[E]) RawErrPredicate[E] {
	return RawErrPredicate[E]{
		pred: func(err error) (E, bool) {
			theErr, ok := p.pred(err)
			if !ok {
				return theErr, false
			}

			return pred(err)
		},
	}
}

func (p RawErrPredicate[E]) Or(pred ErrPredicateFunc[E]) RawErrPredicate[E] {
	return RawErrPredicate[E]{
		pred: func(err error) (E, bool) {
			theErr, ok := p.pred(err)
			if ok {
				return theErr, true
			}

			return pred(err)
		},
	}
}

type TypedErrPredicate[E error] struct {
	typePred ErrPredicate[E]
	pred     TypedErrPredicateFunc[E]
}

func (p TypedErrPredicate[E]) Test(err error) (E, bool) {
	typeErr, ok := p.typePred.Test(err)
	if !ok {
		return typeErr, false
	}

	return typeErr, p.pred(typeErr)
}

func (p TypedErrPredicate[E]) And(pred TypedErrPredicateFunc[E]) TypedErrPredicate[E] {
	return TypedErrPredicate[E]{
		pred: func(err E) bool {
			if ok := p.pred(err); !ok {
				return false
			}

			return pred(err)
		},
	}
}

func (p TypedErrPredicate[E]) Or(pred TypedErrPredicateFunc[E]) TypedErrPredicate[E] {
	return TypedErrPredicate[E]{
		pred: func(err E) bool {
			if ok := p.pred(err); ok {
				return true
			}

			return pred(err)
		},
	}
}
