package maybe_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tompaz3/fungo/maybe"
)

func Test_Maybe_IsDefined(t *testing.T) {
	t.Parallel()

	t.Run(`some value present`, func(t *testing.T) {
		t.Parallel()

		res := maybe.Some(1)

		assert.True(t, res.IsDefined())
	})

	t.Run(`no value present`, func(t *testing.T) {
		t.Parallel()

		res := maybe.None[int]()

		assert.False(t, res.IsDefined())
	})
}

func Test_Maybe_IsEmpty(t *testing.T) {
	t.Parallel()

	t.Run(`some value present`, func(t *testing.T) {
		t.Parallel()

		res := maybe.Some(1)

		assert.False(t, res.IsEmpty())
	})

	t.Run(`no value present`, func(t *testing.T) {
		t.Parallel()

		res := maybe.None[int]()

		assert.True(t, res.IsEmpty())
	})
}

func Test_Maybe_Get(t *testing.T) {
	t.Parallel()

	t.Run(`some value present`, func(t *testing.T) {
		t.Parallel()

		res, err := maybe.Some(1).Get()

		assert.Equal(t, 1, res)
		assert.NoError(t, err)
	})

	t.Run(`no value present`, func(t *testing.T) {
		t.Parallel()

		res, err := maybe.None[int]().Get()

		assert.Zero(t, res)
		assert.Error(t, err)
		assert.Equal(t, maybe.ErrEmptyMaybe, err)
	})
}

func Test_Maybe_OrZero(t *testing.T) {
	t.Parallel()

	t.Run(`some value present`, func(t *testing.T) {
		t.Parallel()

		res := maybe.Some(1).OrZero()

		assert.Equal(t, 1, res)
	})

	t.Run(`no value present`, func(t *testing.T) {
		t.Parallel()

		res := maybe.None[int]().OrZero()

		assert.Zero(t, res)
	})
}

func Test_Maybe_OrElse(t *testing.T) {
	t.Parallel()

	t.Run(`some value present`, func(t *testing.T) {
		t.Parallel()

		res := maybe.Some(1).OrElse(2)

		assert.Equal(t, 1, res)
	})

	t.Run(`no value present`, func(t *testing.T) {
		t.Parallel()

		res := maybe.None[int]().OrElse(2)

		assert.Equal(t, 2, res)
	})
}

func Test_Maybe_OrElseGet(t *testing.T) {
	t.Parallel()

	t.Run(`some value present`, func(t *testing.T) {
		t.Parallel()

		elseEvaluated := false

		res := maybe.Some(1).OrElseGet(func() int {
			elseEvaluated = false
			return 2
		})

		assert.Equal(t, 1, res)
		assert.False(t, elseEvaluated)
	})

	t.Run(`no value present`, func(t *testing.T) {
		t.Parallel()

		elseEvaluated := false

		res := maybe.None[int]().OrElseGet(func() int {
			elseEvaluated = true
			return 2
		})

		assert.Equal(t, 2, res)
		assert.True(t, elseEvaluated)
	})
}

func Test_Maybe_OrElseTryGet(t *testing.T) {
	t.Parallel()

	t.Run(`some value present`, func(t *testing.T) {
		t.Parallel()

		t.Run(`and orElseTryGet returns value`, func(t *testing.T) {
			t.Parallel()

			elseEvaluated := false

			res, err := maybe.Some(1).OrElseTryGet(func() (int, error) {
				elseEvaluated = false
				return 2, nil
			})

			assert.Equal(t, 1, res)
			assert.NoError(t, err)
			assert.False(t, elseEvaluated)
		})

		t.Run(`and orElseTryGet returns error`, func(t *testing.T) {
			t.Parallel()

			elseEvaluated := false

			res, err := maybe.Some(1).OrElseTryGet(func() (int, error) {
				elseEvaluated = false
				return 2, assert.AnError
			})

			assert.Equal(t, 1, res)
			assert.NoError(t, err)
			assert.False(t, elseEvaluated)
		})
	})

	t.Run(`no value present`, func(t *testing.T) {
		t.Parallel()

		t.Run(`and orElseTryGet returns value`, func(t *testing.T) {
			t.Parallel()

			elseEvaluated := false

			res, err := maybe.None[int]().OrElseTryGet(func() (int, error) {
				elseEvaluated = true
				return 2, nil
			})

			assert.Equal(t, 2, res)
			assert.NoError(t, err)
			assert.True(t, elseEvaluated)
		})

		t.Run(`and orElseTryGet returns error`, func(t *testing.T) {
			t.Parallel()

			elseEvaluated := false

			res, err := maybe.None[int]().OrElseTryGet(func() (int, error) {
				elseEvaluated = true
				return 2, assert.AnError
			})

			assert.Equal(t, 2, res)
			assert.Error(t, err)
			assert.Equal(t, assert.AnError, err)
			assert.True(t, elseEvaluated)
		})
	})
}

func Test_Maybe_ZeroValue(t *testing.T) {
	t.Parallel()

	t.Run(`zero value`, func(t *testing.T) {
		t.Parallel()

		var res maybe.Maybe[int]

		assert.False(t, res.IsDefined())
		assert.True(t, res.IsEmpty())
	})

	t.Run(`nil pointer`, func(t *testing.T) {
		t.Parallel()

		var res *maybe.Maybe[int]

		assert.False(t, res.IsDefined())
		assert.True(t, res.IsEmpty())
	})
}

func Test_Map(t *testing.T) {
	t.Parallel()

	t.Run(`some value present`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(i int) string {
			mapperEvaluated = true
			return strconv.Itoa(i)
		}
		mb := maybe.Some(1)

		res := maybe.Map(mb, mapper)

		assert.True(t, res.IsDefined())
		assert.True(t, mapperEvaluated)
		assert.Equal(t, "1", res.OrZero())
	})

	t.Run(`no value present`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(i int) string {
			mapperEvaluated = true
			return strconv.Itoa(i)
		}
		mb := maybe.None[int]()

		res := maybe.Map(mb, mapper)

		assert.True(t, res.IsEmpty())
		assert.False(t, mapperEvaluated)
	})
}

func Test_TryMap_Some(t *testing.T) {
	t.Parallel()

	t.Run(`mapping returns value`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(i int) (string, error) {
			mapperEvaluated = true
			return strconv.Itoa(i), nil
		}
		mb := maybe.Some(1)

		res, err := maybe.TryMap(mb, mapper)

		assert.True(t, res.IsDefined())
		assert.NoError(t, err)
		assert.True(t, mapperEvaluated)
	})

	t.Run(`mapping returns error`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(_ int) (string, error) {
			mapperEvaluated = true
			return "", assert.AnError
		}
		mb := maybe.Some(1)

		res, err := maybe.TryMap(mb, mapper)

		assert.True(t, res.IsEmpty())
		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.True(t, mapperEvaluated)
	})
}

func Test_TryMap_None(t *testing.T) {
	t.Parallel()

	t.Run(`mapping returns value`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(i int) (string, error) {
			mapperEvaluated = true
			return strconv.Itoa(i), nil
		}
		mb := maybe.None[int]()

		res, err := maybe.TryMap(mb, mapper)

		assert.True(t, res.IsEmpty())
		assert.NoError(t, err)
		assert.False(t, mapperEvaluated)
	})

	t.Run(`mapping returns error`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(_ int) (string, error) {
			mapperEvaluated = true
			return "", assert.AnError
		}
		mb := maybe.None[int]()

		res, err := maybe.TryMap(mb, mapper)

		assert.True(t, res.IsEmpty())
		assert.NoError(t, err)
		assert.False(t, mapperEvaluated)
	})
}

func Test_FlatMap_Some(t *testing.T) {
	t.Parallel()

	t.Run(`mapper returned some`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(i int) *maybe.Maybe[string] {
			mapperEvaluated = true
			return maybe.Some(strconv.Itoa(i))
		}
		mb := maybe.Some(1)

		res := maybe.FlatMap(mb, mapper)

		assert.True(t, res.IsDefined())
		assert.True(t, mapperEvaluated)
		assert.Equal(t, "1", res.OrZero())
	})

	t.Run(`mapper returned none`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(_ int) *maybe.Maybe[string] {
			mapperEvaluated = true
			return maybe.None[string]()
		}
		mb := maybe.Some(1)

		res := maybe.FlatMap(mb, mapper)

		assert.True(t, res.IsEmpty())
		assert.True(t, mapperEvaluated)
	})
}

func Test_FlatMap_None(t *testing.T) {
	t.Parallel()

	t.Run(`mapper returned some`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(i int) *maybe.Maybe[string] {
			mapperEvaluated = true
			return maybe.Some(strconv.Itoa(i))
		}
		mb := maybe.None[int]()

		res := maybe.FlatMap(mb, mapper)

		assert.True(t, res.IsEmpty())
		assert.False(t, mapperEvaluated)
	})

	t.Run(`mapper returned none`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(_ int) *maybe.Maybe[string] {
			mapperEvaluated = true
			return maybe.None[string]()
		}
		mb := maybe.None[int]()

		res := maybe.FlatMap(mb, mapper)

		assert.True(t, res.IsEmpty())
		assert.False(t, mapperEvaluated)
	})
}

func Test_FlatMap_MonadLaws(t *testing.T) {
	t.Parallel()
	t.Run("left identity", func(t *testing.T) {
		t.Parallel()

		f := func(x int) *maybe.Maybe[string] {
			return maybe.Some(fmt.Sprintf("number: %d", x))
		}
		value := 2
		mb := maybe.Some(value)

		lhs := maybe.FlatMap(mb, f)

		rhs := f(value)

		assert.Equal(t, lhs, rhs)
	})

	t.Run("right identity", func(t *testing.T) {
		t.Parallel()

		mb := maybe.Some("some: 2")

		lhs := maybe.FlatMap(mb, maybe.Some)

		rhs := mb

		assert.Equal(t, lhs, rhs)
	})

	t.Run("associativity", func(t *testing.T) {
		t.Parallel()

		f := func(x int) *maybe.Maybe[string] {
			return maybe.Some(fmt.Sprintf("some: %d", x))
		}
		g := func(s string) *maybe.Maybe[int] {
			return maybe.Some(len(s))
		}

		mb := maybe.Some(22)

		lhs := maybe.FlatMap(maybe.FlatMap(mb, f), g)

		rhs := maybe.FlatMap(mb, func(x int) *maybe.Maybe[int] {
			return maybe.FlatMap(f(x), g)
		})

		assert.Equal(t, lhs, rhs)
	})
}

func Test_TryFlatMap_Some(t *testing.T) {
	t.Parallel()

	t.Run(`mapper returned some`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(i int) (*maybe.Maybe[string], error) {
			mapperEvaluated = true
			return maybe.Some(strconv.Itoa(i)), nil
		}
		mb := maybe.Some(1)

		res, err := maybe.TryFlatMap(mb, mapper)

		assert.True(t, res.IsDefined())
		assert.NoError(t, err)
		assert.True(t, mapperEvaluated)
		assert.Equal(t, "1", res.OrZero())
	})

	t.Run(`mapper returned none`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(_ int) (*maybe.Maybe[string], error) {
			mapperEvaluated = true
			return maybe.None[string](), nil
		}
		mb := maybe.Some(1)

		res, err := maybe.TryFlatMap(mb, mapper)

		assert.True(t, res.IsEmpty())
		assert.NoError(t, err)
		assert.True(t, mapperEvaluated)
	})

	t.Run(`mapper returned error`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(_ int) (*maybe.Maybe[string], error) {
			mapperEvaluated = true
			return nil, assert.AnError
		}
		mb := maybe.Some(1)

		res, err := maybe.TryFlatMap(mb, mapper)

		assert.True(t, res.IsEmpty())
		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		assert.True(t, mapperEvaluated)
	})
}

func Test_TryFlatMap_None(t *testing.T) {
	t.Parallel()

	t.Run(`mapper returned some`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(i int) (*maybe.Maybe[string], error) {
			mapperEvaluated = true
			return maybe.Some(strconv.Itoa(i)), nil
		}
		mb := maybe.None[int]()

		res, err := maybe.TryFlatMap(mb, mapper)

		assert.True(t, res.IsEmpty())
		assert.NoError(t, err)
		assert.False(t, mapperEvaluated)
	})

	t.Run(`mapper returned none`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(_ int) (*maybe.Maybe[string], error) {
			mapperEvaluated = true
			return maybe.None[string](), nil
		}
		mb := maybe.None[int]()

		res, err := maybe.TryFlatMap(mb, mapper)

		assert.True(t, res.IsEmpty())
		assert.NoError(t, err)
		assert.False(t, mapperEvaluated)
	})

	t.Run(`mapper returned error`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(_ int) (*maybe.Maybe[string], error) {
			mapperEvaluated = true
			return nil, assert.AnError
		}
		mb := maybe.None[int]()

		res, err := maybe.TryFlatMap(mb, mapper)

		assert.True(t, res.IsEmpty())
		assert.NoError(t, err)
		assert.False(t, mapperEvaluated)
	})
}

func Test_Filter_Some(t *testing.T) {
	t.Parallel()

	t.Run(`predicate returns true`, func(t *testing.T) {
		t.Parallel()

		predicateEvaluated := false
		predicate := func(_ int) bool {
			predicateEvaluated = true
			return true
		}

		mb := maybe.Some(1)

		res := mb.Filter(predicate)

		assert.True(t, res.IsDefined())
		assert.True(t, predicateEvaluated)
	})

	t.Run(`predicate returns false`, func(t *testing.T) {
		t.Parallel()

		predicateEvaluated := false
		predicate := func(_ int) bool {
			predicateEvaluated = true
			return false
		}

		mb := maybe.Some(1)

		res := mb.Filter(predicate)

		assert.True(t, res.IsEmpty())
		assert.True(t, predicateEvaluated)
	})
}

func Test_Filter_None(t *testing.T) {
	t.Parallel()

	t.Run(`predicate returns true`, func(t *testing.T) {
		t.Parallel()

		predicateEvaluated := false
		predicate := func(_ int) bool {
			predicateEvaluated = true
			return true
		}

		mb := maybe.None[int]()

		res := mb.Filter(predicate)

		assert.True(t, res.IsEmpty())
		assert.False(t, predicateEvaluated)
	})

	t.Run(`predicate returns false`, func(t *testing.T) {
		t.Parallel()

		predicateEvaluated := false
		predicate := func(_ int) bool {
			predicateEvaluated = true
			return false
		}

		mb := maybe.None[int]()

		res := mb.Filter(predicate)

		assert.True(t, res.IsEmpty())
		assert.False(t, predicateEvaluated)
	})
}

func Test_Map_Receiver(t *testing.T) {
	t.Parallel()

	t.Run(`some value present`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(i int) int {
			mapperEvaluated = true
			return i * 2
		}
		mb := maybe.Some(1)

		res := mb.Map(mapper)

		assert.True(t, res.IsDefined())
		assert.True(t, mapperEvaluated)
		assert.Equal(t, 2, res.OrZero())
	})

	t.Run(`no value present`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(i int) int {
			mapperEvaluated = true
			return i * 2
		}
		mb := maybe.None[int]()

		res := mb.Map(mapper)

		assert.True(t, res.IsEmpty())
		assert.False(t, mapperEvaluated)
	})
}

func Test_FlatMap_Receiver_Some(t *testing.T) {
	t.Parallel()

	t.Run(`mapper returned some`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(i int) *maybe.Maybe[int] {
			mapperEvaluated = true
			return maybe.Some(i * 2)
		}
		mb := maybe.Some(1)

		res := mb.FlatMap(mapper)

		assert.True(t, res.IsDefined())
		assert.True(t, mapperEvaluated)
		assert.Equal(t, 2, res.OrZero())
	})

	t.Run(`mapper returned none`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(_ int) *maybe.Maybe[int] {
			mapperEvaluated = true
			return maybe.None[int]()
		}
		mb := maybe.Some(1)

		res := mb.FlatMap(mapper)

		assert.True(t, res.IsEmpty())
		assert.True(t, mapperEvaluated)
	})
}

func Test_FlatMap_Receiver_None(t *testing.T) {
	t.Parallel()

	t.Run(`mapper returned some`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(i int) *maybe.Maybe[int] {
			mapperEvaluated = true
			return maybe.Some(i * 2)
		}
		mb := maybe.None[int]()

		res := mb.FlatMap(mapper)

		assert.True(t, res.IsEmpty())
		assert.False(t, mapperEvaluated)
	})

	t.Run(`mapper returned none`, func(t *testing.T) {
		t.Parallel()

		mapperEvaluated := false
		mapper := func(_ int) *maybe.Maybe[int] {
			mapperEvaluated = true
			return maybe.None[int]()
		}
		mb := maybe.None[int]()

		res := mb.FlatMap(mapper)

		assert.True(t, res.IsEmpty())
		assert.False(t, mapperEvaluated)
	})
}

type maybeTestInterface interface {
	Do()
}

type maybeTestStruct struct{}

var _ maybeTestInterface = (*maybeTestStruct)(nil)

func (m *maybeTestStruct) Do() {}

func Test_OfNillable(t *testing.T) {
	t.Parallel()

	fTest := func(t *testing.T, name string, value any, expectedDefined bool) {
		t.Helper()
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res := maybe.OfNillable(value)
			assert.Equal(t, expectedDefined, res.IsDefined())
		})
	}

	fTest(t, `non-nil value`, 1, true)
	fTest(t, `nil value`, nil, false)
	fTest(t, `non-nil channel`, make(chan int), true)
	var ch chan int
	fTest(t, `nil channel`, ch, false)
	fTest(t, `non-nil function`, func() {}, true)
	var fn func()
	fTest(t, `nil function`, fn, false)
	var nonNilInt maybeTestInterface = &maybeTestStruct{}
	fTest(t, `non-nil interface`, nonNilInt, true)
	var nilInt maybeTestInterface
	fTest(t, `nil interface`, nilInt, false)
	var nilStruct *maybeTestStruct
	nilTypedInt := nilStruct
	fTest(t, `nil typed interface`, nilTypedInt, false)
	fTest(t, `non-nil map`, map[string]string{}, true)
	var nilMap map[string]string
	fTest(t, `nil map`, nilMap, false)
	nonNilValue := 1
	fTest(t, `non-nil pointer`, &nonNilValue, true)
	var nilPointer *int
	fTest(t, `nil pointer`, nilPointer, false)
	fTest(t, `non-nil slice`, make([]int, 0), true)
	var nilSlice []int
	fTest(t, `nil slice`, nilSlice, false)
}
