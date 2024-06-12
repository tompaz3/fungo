package match

import "errors"

type ErrorPredicate[E error] func(err E) bool

// ErrorType - tests if given error is of the expected type (generic parameter)
// and returns the typed error with a boolean test result.
func ErrorType[E error](err error) (E, bool) {
	var typedErr E
	if errors.As(err, &typedErr) {
		return typedErr, true
	}

	return typedErr, false
}

// ErrorTypeMatches - tests if given error is of the expected type (generic parameter)
// and matches given predicate, returns the typed error and a boolean test result.
func ErrorTypeMatches[E error](err error, pred ErrorPredicate[E]) (E, bool) {
	typedErr, ok := ErrorType[E](err)
	if !ok {
		return typedErr, ok
	}

	if pred(typedErr) {
		return typedErr, true
	}

	var empty E

	return empty, false
}

// ErrorMatches - tests if given error matches given predicate and returns the error with a boolean test result
//
//nolint:revive // this function is used to test error against the predicate and return it in case of success
func ErrorMatches(err error, pred ErrorPredicate[error]) (error, bool) {
	if pred(err) {
		return err, true
	}

	var empty error

	return empty, false
}

func (p ErrorPredicate[E]) And(pred ErrorPredicate[E]) ErrorPredicate[E] {
	return func(err E) bool {
		return p(err) && pred(err)
	}
}

func (p ErrorPredicate[E]) Or(pred ErrorPredicate[E]) ErrorPredicate[E] {
	return func(err E) bool {
		return p(err) || pred(err)
	}
}
