package ternary_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tompaz3/fungo/ternary"
)

const (
	elseString = "else"
	thenString = "then"
)

func Test_Ternary_ConditionTrue_ConditionEagerlyEvaluated(t *testing.T) {
	t.Parallel()

	t.Run(`and then eagerly evaluated`, func(t *testing.T) {
		t.Parallel()

		t.Run(`and else eagerly evaluated`, func(t *testing.T) {
			t.Parallel()

			res := ternary.If[string](true).
				Then(thenString).
				Else(elseString)

			assert.Equal(t, thenString, res)
		})

		t.Run(`and else lazily evaluated`, func(t *testing.T) {
			t.Parallel()

			elseEvaluated := false

			res := ternary.If[string](true).
				Then(thenString).
				ElseF(func() string {
					elseEvaluated = true
					return elseString
				})

			assert.False(t, elseEvaluated)
			assert.Equal(t, thenString, res)
		})
	})

	t.Run(`and then lazily evaluated`, func(t *testing.T) {
		t.Parallel()

		t.Run(`and else eagerly evaluated`, func(t *testing.T) {
			t.Parallel()

			thenEvaluated := false

			res := ternary.If[string](true).
				ThenF(func() string {
					thenEvaluated = true
					return thenString
				}).
				Else(elseString)

			assert.True(t, thenEvaluated)
			assert.Equal(t, thenString, res)
		})

		t.Run(`and else lazily evaluated`, func(t *testing.T) {
			t.Parallel()

			thenEvaluated := false
			elseEvaulated := false

			res := ternary.If[string](true).
				ThenF(func() string {
					thenEvaluated = true
					return thenString
				}).
				ElseF(func() string {
					elseEvaulated = true
					return elseString
				})

			assert.True(t, thenEvaluated)
			assert.False(t, elseEvaulated)
			assert.Equal(t, thenString, res)
		})
	})
}

func Test_Ternary_ConditionTrue_ConditionLazilyEvaluated(t *testing.T) {
	t.Parallel()

	t.Run(`and then eagerly evaluated`, func(t *testing.T) {
		t.Parallel()

		t.Run(`and else eagerly evaluated`, func(t *testing.T) {
			t.Parallel()

			conditionEvaluated := false

			res := ternary.IfF[string](func() bool {
				conditionEvaluated = true
				return true
			}).
				Then(thenString).
				Else(elseString)

			assert.True(t, conditionEvaluated)
			assert.Equal(t, thenString, res)
		})

		t.Run(`and else lazily evaluated`, func(t *testing.T) {
			t.Parallel()

			conditionEvaluated := false
			elseEvaluated := false

			res := ternary.IfF[string](func() bool {
				conditionEvaluated = true
				return true
			}).
				Then(thenString).
				ElseF(func() string {
					elseEvaluated = true
					return elseString
				})

			assert.True(t, conditionEvaluated)
			assert.False(t, elseEvaluated)
			assert.Equal(t, thenString, res)
		})
	})

	t.Run(`and then lazily evaluated`, func(t *testing.T) {
		t.Parallel()

		t.Run(`and else eagerly evaluated`, func(t *testing.T) {
			t.Parallel()

			conditionEvaluated := false
			thenEvaluated := false

			res := ternary.IfF[string](func() bool {
				conditionEvaluated = true
				return true
			}).
				ThenF(func() string {
					thenEvaluated = true
					return thenString
				}).
				Else(elseString)

			assert.True(t, conditionEvaluated)
			assert.True(t, thenEvaluated)
			assert.Equal(t, thenString, res)
		})

		t.Run(`and else lazily evaluated`, func(t *testing.T) {
			t.Parallel()

			conditionEvaluated := false
			thenEvaluated := false
			elseEvaluated := false

			res := ternary.IfF[string](func() bool {
				conditionEvaluated = true
				return true
			}).
				ThenF(func() string {
					thenEvaluated = true
					return thenString
				}).
				ElseF(func() string {
					elseEvaluated = true
					return elseString
				})

			assert.True(t, conditionEvaluated)
			assert.True(t, thenEvaluated)
			assert.False(t, elseEvaluated)
			assert.Equal(t, thenString, res)
		})
	})
}

func Test_Ternary_ConditionFalse_ConditionEagerlyEvaluated(t *testing.T) {
	t.Parallel()

	t.Run(`and then eagerly evaluated`, func(t *testing.T) {
		t.Parallel()

		t.Run(`and else eagerly evaluated`, func(t *testing.T) {
			t.Parallel()

			res := ternary.If[string](false).
				Then(thenString).
				Else(elseString)

			assert.Equal(t, elseString, res)
		})

		t.Run(`and else lazily evaluated`, func(t *testing.T) {
			t.Parallel()

			elseEvaluated := false

			res := ternary.If[string](false).
				Then(thenString).
				ElseF(func() string {
					elseEvaluated = true
					return elseString
				})

			assert.True(t, elseEvaluated)
			assert.Equal(t, elseString, res)
		})
	})

	t.Run(`and then lazily evaluated`, func(t *testing.T) {
		t.Parallel()

		t.Run(`and else eagerly evaluated`, func(t *testing.T) {
			t.Parallel()

			thenEvaluated := false

			res := ternary.If[string](false).
				ThenF(func() string {
					thenEvaluated = true
					return thenString
				}).
				Else(elseString)

			assert.False(t, thenEvaluated)
			assert.Equal(t, elseString, res)
		})

		t.Run(`and else lazily evaluated`, func(t *testing.T) {
			t.Parallel()

			thenEvaluated := false
			elseEvaluated := false

			res := ternary.If[string](false).
				ThenF(func() string {
					thenEvaluated = true
					return thenString
				}).
				ElseF(func() string {
					elseEvaluated = true
					return elseString
				})

			assert.False(t, thenEvaluated)
			assert.True(t, elseEvaluated)
			assert.Equal(t, elseString, res)
		})
	})
}

func Test_Ternary_ConditionFalse_ConditionLazilyEvaluated(t *testing.T) {
	t.Run(`and then eagerly evaluated`, func(t *testing.T) {
		t.Parallel()

		t.Run(`and else eagerly evaluated`, func(t *testing.T) {
			t.Parallel()

			conditionEvaluated := false

			res := ternary.IfF[string](func() bool {
				conditionEvaluated = true
				return false
			}).
				Then(thenString).
				Else(elseString)

			assert.True(t, conditionEvaluated)
			assert.Equal(t, elseString, res)
		})

		t.Run(`and else lazily evaluated`, func(t *testing.T) {
			t.Parallel()

			conditionEvaluated := false
			elseEvaluated := false

			res := ternary.IfF[string](func() bool {
				conditionEvaluated = true
				return false
			}).
				Then(thenString).
				ElseF(func() string {
					elseEvaluated = true
					return elseString
				})

			assert.True(t, conditionEvaluated)
			assert.True(t, elseEvaluated)
			assert.Equal(t, elseString, res)
		})
	})

	t.Run(`and then lazily evaluated`, func(t *testing.T) {
		t.Parallel()

		t.Run(`and else eagerly evaluated`, func(t *testing.T) {
			t.Parallel()

			conditionEvaluated := false
			thenEvaluated := false

			res := ternary.IfF[string](func() bool {
				conditionEvaluated = true
				return false
			}).
				ThenF(func() string {
					thenEvaluated = true
					return thenString
				}).
				Else(elseString)

			assert.True(t, conditionEvaluated)
			assert.False(t, thenEvaluated)
			assert.Equal(t, elseString, res)
		})

		t.Run(`and else lazily evaluated`, func(t *testing.T) {
			t.Parallel()

			conditionEvaluated := false
			thenEvaluated := false
			elseEvaluated := false

			res := ternary.IfF[string](func() bool {
				conditionEvaluated = true
				return false
			}).
				ThenF(func() string {
					thenEvaluated = true
					return thenString
				}).
				ElseF(func() string {
					elseEvaluated = true
					return elseString
				})

			assert.True(t, conditionEvaluated)
			assert.False(t, thenEvaluated)
			assert.True(t, elseEvaluated)
			assert.Equal(t, elseString, res)
		})
	})
}
