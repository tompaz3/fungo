package maybe_test

import (
	"errors"
	"strconv"

	"github.com/tompaz3/fungo/maybe"

	g "github.com/onsi/ginkgo/v2"
	o "github.com/onsi/gomega"
)

var errTestError = errors.New("test error")

var _ = g.Describe("Maybe", func() {
	g.Describe("IsDefined", func() {
		g.When("some value present", func() {
			g.It("should return true", func() {
				o.Expect(maybe.Some(1).IsDefined()).To(o.BeTrue())
			})
		})

		g.When("no value present", func() {
			g.It("should return false", func() {
				o.Expect(maybe.None[int]().IsDefined()).To(o.BeFalse())
			})
		})
	})

	g.Describe("IsEmpty", func() {
		g.When("some value present", func() {
			g.It("should return false", func() {
				o.Expect(maybe.Some(1).IsEmpty()).To(o.BeFalse())
			})
		})

		g.When("no value present", func() {
			g.It("should return true", func() {
				o.Expect(maybe.None[int]().IsEmpty()).To(o.BeTrue())
			})
		})
	})

	g.Describe("Get", func() {
		g.When("some value present", func() {
			g.It("should return the value", func() {
				value, err := maybe.Some(1).Get()
				o.Expect(value).To(o.Equal(1))
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})

		g.When("no value present", func() {
			g.It("should return an error", func() {
				value, err := maybe.None[int]().Get()
				o.Expect(value).To(o.Equal(0))
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(err).Should(o.Equal(maybe.ErrEmptyMaybe))
			})
		})
	})

	g.Describe("OrZero", func() {
		g.When("some value present", func() {
			g.It("should return the value", func() {
				value := maybe.Some(1).OrZero()
				o.Expect(value).To(o.Equal(1))
			})
		})

		g.When("when no value present", func() {
			g.It("should panic", func() {
				value := maybe.None[int]().OrZero()
				o.Expect(value).To(o.BeZero())
			})
		})
	})

	g.Describe("OrElse", func() {
		g.When("some value present", func() {
			g.It("should return the value", func() {
				value := maybe.Some(1).OrElse(2)
				o.Expect(value).To(o.Equal(1))
			})
		})

		g.When("when no value present", func() {
			g.It("should return the other value", func() {
				value := maybe.None[int]().OrElse(2)
				o.Expect(value).To(o.Equal(2))
			})
		})
	})

	g.Describe("OrElseGet", func() {
		g.When("some value present", func() {
			g.It("should return the value", func() {
				value := maybe.Some(1).OrElseGet(func() int { return 2 })
				o.Expect(value).To(o.Equal(1))
			})
		})

		g.When("no value present", func() {
			g.It("should return the other value", func() {
				value := maybe.None[int]().OrElseGet(func() int { return 2 })
				o.Expect(value).To(o.Equal(2))
			})
		})
	})

	g.Describe("OrElseTryGet", func() {
		g.When("some value present", func() {
			g.Context("and other function returns value", func() {
				g.It("should return the value", func() {
					value, err := maybe.Some(1).
						OrElseTryGet(func() (int, error) { return 2, nil })
					o.Expect(value).To(o.Equal(1))
					o.Expect(err).ShouldNot(o.HaveOccurred())
				})
			})

			g.Context("and other function returns error", func() {
				g.It("should return the value", func() {
					value, err := maybe.Some(1).
						OrElseTryGet(func() (int, error) { return 0, errTestError })
					o.Expect(value).To(o.Equal(1))
					o.Expect(err).ShouldNot(o.HaveOccurred())
				})
			})
		})

		g.When("no value present", func() {
			g.Context("and other function returns value", func() {
				g.It("should return the other value", func() {
					value, err := maybe.None[int]().
						OrElseTryGet(func() (int, error) { return 2, nil })
					o.Expect(value).To(o.Equal(2))
					o.Expect(err).ShouldNot(o.HaveOccurred())
				})
			})

			g.Context("and other function returns error", func() {
				g.It("should return the other error", func() {
					value, err := maybe.None[int]().
						OrElseTryGet(func() (int, error) { return 0, errTestError })
					o.Expect(value).To(o.BeZero())
					o.Expect(err).Should(o.HaveOccurred())
					o.Expect(err).To(o.Equal(errTestError))
				})
			})
		})
	})
})

var _ = g.Describe("MaybeZeroValue", func() {
	g.When("maybe is zero", func() {
		g.It("should follow None contract", func() {
			var mb maybe.Maybe[int]
			o.Expect(mb.IsEmpty()).To(o.BeTrue())
			o.Expect(mb.IsDefined()).To(o.BeFalse())
		})
	})
	g.When("maybe pointer is nil", func() {
		g.It("should follow None contract", func() {
			var mb *maybe.Maybe[int]
			o.Expect(mb.IsEmpty()).To(o.BeTrue())
			o.Expect(mb.IsDefined()).To(o.BeFalse())
		})
	})
})

var _ = g.Describe("Map", func() {
	g.When("some value present", func() {
		g.It("should return the mapped value", func() {
			mb := maybe.Some(1)
			mapped := maybe.Map(mb, strconv.Itoa)
			o.Expect(mapped.IsDefined()).To(o.BeTrue())
			o.Expect(mapped.OrZero()).To(o.Equal("1"))
		})
	})

	g.When("no value present", func() {
		g.It("should return None", func() {
			mb := maybe.None[int]()
			mapped := maybe.Map(mb, strconv.Itoa)
			o.Expect(mapped.IsEmpty()).To(o.BeTrue())
		})
	})
})

var _ = g.Describe("TryMap", func() {
	g.Describe("some value present", func() {
		g.When("mapping function returns value", func() {
			g.It("should return the mapped value", func() {
				mb := maybe.Some(1)
				mapped, err := maybe.TryMap(mb, func(it int) (string, error) {
					return strconv.Itoa(it), nil
				})
				o.Expect(mapped.IsDefined()).To(o.BeTrue())
				o.Expect(mapped.OrZero()).To(o.Equal("1"))
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})

		g.When("mapping function returns error", func() {
			g.It("should return the error", func() {
				mb := maybe.Some(1)
				mapped, err := maybe.TryMap(mb, func(_ int) (string, error) {
					return "", errTestError
				})
				o.Expect(mapped.IsEmpty()).To(o.BeTrue())
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(err).To(o.Equal(errTestError))
			})
		})
	})

	g.Describe("no value present", func() {
		g.When("mapping function returns value", func() {
			g.It("should return None", func() {
				mb := maybe.None[int]()
				mapped, err := maybe.TryMap(mb, func(it int) (string, error) {
					return strconv.Itoa(it), nil
				})
				o.Expect(mapped.IsEmpty()).To(o.BeTrue())
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})

		g.When("mapping function returns error", func() {
			g.It("should return None", func() {
				mb := maybe.None[int]()
				mapped, err := maybe.TryMap(mb, func(_ int) (string, error) {
					return "", errTestError
				})
				o.Expect(mapped.IsEmpty()).To(o.BeTrue())
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})
	})
})

var _ = g.Describe("FlatMap", func() {
	g.When("some value present", func() {
		g.It("should return the mapped value", func() {
			mb := maybe.Some(1)
			mapped := maybe.FlatMap(mb, func(it int) *maybe.Maybe[string] {
				return maybe.Some(strconv.Itoa(it))
			})
			o.Expect(mapped.IsDefined()).To(o.BeTrue())
			o.Expect(mapped.OrZero()).To(o.Equal("1"))
		})
	})

	g.When("no value present", func() {
		g.It("should return None", func() {
			mb := maybe.None[int]()
			mapped := maybe.FlatMap(mb, func(it int) *maybe.Maybe[string] {
				return maybe.Some(strconv.Itoa(it))
			})
			o.Expect(mapped.IsEmpty()).To(o.BeTrue())
		})
	})
})

var _ = g.Describe("TryFlatMap", func() {
	g.Describe("some value present", func() {
		g.When("mapping function returns value", func() {
			g.It("should return the mapped value", func() {
				mb := maybe.Some(1)
				mapped, err := maybe.TryFlatMap(mb, func(it int) (*maybe.Maybe[string], error) {
					return maybe.Some(strconv.Itoa(it)), nil
				})
				o.Expect(mapped.IsDefined()).To(o.BeTrue())
				o.Expect(mapped.OrZero()).To(o.Equal("1"))
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})

		g.When("mapping function returns error", func() {
			g.It("should return the error", func() {
				mb := maybe.Some(1)
				mapped, err := maybe.TryFlatMap(mb, func(_ int) (*maybe.Maybe[string], error) {
					return maybe.None[string](), errTestError
				})
				o.Expect(mapped.IsEmpty()).To(o.BeTrue())
				o.Expect(err).Should(o.HaveOccurred())
				o.Expect(err).To(o.Equal(errTestError))
			})
		})
	})

	g.Describe("no value present", func() {
		g.When("mapping function returns value", func() {
			g.It("should return None", func() {
				mb := maybe.None[int]()
				mapped, err := maybe.TryFlatMap(mb, func(it int) (*maybe.Maybe[string], error) {
					return maybe.Some(strconv.Itoa(it)), nil
				})
				o.Expect(mapped.IsEmpty()).To(o.BeTrue())
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})

		g.When("mapping function returns error", func() {
			g.It("should return None", func() {
				mb := maybe.None[int]()
				mapped, err := maybe.TryFlatMap(mb, func(_ int) (*maybe.Maybe[string], error) {
					return maybe.None[string](), errTestError
				})
				o.Expect(mapped.IsEmpty()).To(o.BeTrue())
				o.Expect(err).ShouldNot(o.HaveOccurred())
			})
		})
	})
})

var _ = g.Describe("Filter", func() {
	g.Describe("some value present", func() {
		g.When("predicate returns true", func() {
			g.It("should return Some", func() {
				mb := maybe.Some(1)
				filtered := mb.Filter(func(i int) bool {
					return i == 1
				})
				o.Expect(filtered.IsDefined()).To(o.BeTrue())
				o.Expect(filtered.OrZero()).To(o.Equal(1))
			})
		})

		g.When("predicate returns false", func() {
			g.It("should return None", func() {
				mb := maybe.Some(1)
				filtered := mb.Filter(func(i int) bool {
					return i == 2
				})
				o.Expect(filtered.IsEmpty()).To(o.BeTrue())
			})
		})
	})

	g.When("no value present", func() {
		g.It("should return None", func() {
			mb := maybe.None[int]()
			predicateExecuted := false
			filtered := mb.Filter(func(i int) bool {
				predicateExecuted = true
				return i == 1
			})
			o.Expect(filtered.IsEmpty()).To(o.BeTrue())
			o.Expect(predicateExecuted).To(o.BeFalse())
		})
	})
})

var _ = g.Describe("maybe.Map", func() {
	g.When("some value present", func() {
		g.It("should return the mapped value", func() {
			mb := maybe.Some(1)
			mapped := mb.Map(func(i int) int {
				return i * 2
			})
			o.Expect(mapped.IsDefined()).To(o.BeTrue())
			o.Expect(mapped.OrZero()).To(o.Equal(2))
		})
	})

	g.When("no value present", func() {
		g.It("should return None", func() {
			mb := maybe.None[int]()
			mapped := mb.Map(func(i int) int {
				return i * 2
			})
			o.Expect(mapped.IsEmpty()).To(o.BeTrue())
		})
	})
})

var _ = g.Describe("maybe.FlatMap", func() {
	g.When("some value present", func() {
		g.It("should return the mapped value", func() {
			mb := maybe.Some(1)
			mapped := mb.FlatMap(func(it int) *maybe.Maybe[int] {
				return maybe.Some(it * 2)
			})
			o.Expect(mapped.IsDefined()).To(o.BeTrue())
			o.Expect(mapped.OrZero()).To(o.Equal(2))
		})
	})

	g.When("no value present", func() {
		g.It("should return None", func() {
			mb := maybe.None[int]()
			mapped := mb.FlatMap(func(it int) *maybe.Maybe[int] {
				return maybe.Some(it * 2)
			})
			o.Expect(mapped.IsEmpty()).To(o.BeTrue())
		})
	})
})

type maybeTestInterface interface {
	Do()
}

type maybeTestStruct struct{}

var _ maybeTestInterface = (*maybeTestStruct)(nil)

func (m *maybeTestStruct) Do() {}

var _ = g.Describe("OfNillable", func() {
	g.When("non-nil value provided", func() {
		g.It("should return Some", func() {
			mb := maybe.OfNillable(1)
			o.Expect(mb.IsDefined()).To(o.BeTrue())
		})
	})

	g.When("nil value provided", func() {
		g.It("should return None", func() {
			mb := maybe.OfNillable[*int](nil)
			o.Expect(mb.IsEmpty()).To(o.BeTrue())
		})
	})

	g.When("nil channel", func() {
		g.It("should return None", func() {
			var ch chan int
			mb := maybe.OfNillable(ch)
			o.Expect(mb.IsEmpty()).To(o.BeTrue())
		})
	})

	g.When("non-nil channel", func() {
		g.It("should return Some", func() {
			ch := make(chan int)
			mb := maybe.OfNillable(ch)
			o.Expect(mb.IsDefined()).To(o.BeTrue())
		})
	})

	g.When("nil function", func() {
		g.It("should return None", func() {
			var fn func()
			mb := maybe.OfNillable(fn)
			o.Expect(mb.IsEmpty()).To(o.BeTrue())
		})
	})

	g.When("non-nil function", func() {
		g.It("should return Some", func() {
			fn := func() {}
			mb := maybe.OfNillable(fn)
			o.Expect(mb.IsDefined()).To(o.BeTrue())
		})
	})

	g.When("nil interface", func() {
		g.It("should return None", func() {
			var it maybeTestInterface
			mb := maybe.OfNillable(it)
			o.Expect(mb.IsEmpty()).To(o.BeTrue())
		})
	})

	g.When("nil typed interface", func() {
		g.It("should return None", func() {
			var str *maybeTestStruct
			var it maybeTestInterface = str
			mb := maybe.OfNillable(it)
			o.Expect(mb.IsEmpty()).To(o.BeTrue())
		})
	})

	g.When("non-nil interface", func() {
		g.It("should return Some", func() {
			var it maybeTestInterface = &maybeTestStruct{}
			mb := maybe.OfNillable(it)
			o.Expect(mb.IsDefined()).To(o.BeTrue())
		})
	})

	g.When("nil map", func() {
		g.It("should return None", func() {
			var m map[string]int
			mb := maybe.OfNillable(m)
			o.Expect(mb.IsEmpty()).To(o.BeTrue())
		})
	})

	g.When("non-nil map", func() {
		g.It("should return Some", func() {
			m := make(map[string]int)
			mb := maybe.OfNillable(m)
			o.Expect(mb.IsDefined()).To(o.BeTrue())
		})
	})

	g.When("nil pointer", func() {
		g.It("should return None", func() {
			var ptr *int
			mb := maybe.OfNillable(ptr)
			o.Expect(mb.IsEmpty()).To(o.BeTrue())
		})
	})

	g.When("non-nil pointer", func() {
		g.It("should return Some", func() {
			var i int
			mb := maybe.OfNillable(&i)
			o.Expect(mb.IsDefined()).To(o.BeTrue())
		})
	})

	g.When("nil slice", func() {
		g.It("should return None", func() {
			var sl []int
			mb := maybe.OfNillable(sl)
			o.Expect(mb.IsEmpty()).To(o.BeTrue())
		})
	})

	g.When("non-nil slice", func() {
		g.It("should return Some", func() {
			sl := make([]int, 0)
			mb := maybe.OfNillable(sl)
			o.Expect(mb.IsDefined()).To(o.BeTrue())
		})
	})
})

var _ = g.Describe("MonadLaws", func() {
	g.Describe("left identity", func() {
		g.It("should hold", func() {
			f := func(i int) *maybe.Maybe[string] {
				return maybe.Some(strconv.Itoa(i))
			}
			value := 1

			lhs := maybe.FlatMap(maybe.Some(value), f)
			rhs := f(value)

			o.Expect(lhs).To(o.Equal(rhs))
		})
	})

	g.Describe("right identity", func() {
		g.It("should hold", func() {
			mb := maybe.Some(1)

			lhs := maybe.FlatMap(mb, maybe.Some[int])
			rhs := mb

			o.Expect(lhs).To(o.Equal(rhs))
		})
	})

	g.Describe("associativity", func() {
		g.It("should hold", func() {
			f := func(i int) *maybe.Maybe[string] {
				return maybe.Some(strconv.Itoa(i))
			}
			g := func(i string) *maybe.Maybe[int] {
				return maybe.Some(len(i))
			}
			mb := maybe.Some(10)

			lhs := maybe.FlatMap(maybe.FlatMap(mb, f), g)
			rhs := maybe.FlatMap(mb, func(i int) *maybe.Maybe[int] {
				return maybe.FlatMap(f(i), g)
			})

			o.Expect(lhs).To(o.Equal(rhs))
		})
	})
})

var _ = g.Describe("maybe.MonadLaws", func() {
	g.Describe("left identity", func() {
		g.It("should hold", func() {
			f := func(i int) *maybe.Maybe[int] {
				return maybe.Some(i * 2)
			}
			value := 1

			lhs := maybe.Some(value).FlatMap(f)
			rhs := f(value)

			o.Expect(lhs).To(o.Equal(rhs))
		})
	})

	g.Describe("right identity", func() {
		g.It("should hold", func() {
			mb := maybe.Some(1)

			lhs := mb.FlatMap(maybe.Some[int])
			rhs := mb

			o.Expect(lhs).To(o.Equal(rhs))
		})
	})

	g.Describe("associativity", func() {
		g.It("should hold", func() {
			f := func(i int) *maybe.Maybe[int] {
				return maybe.Some(i * 2)
			}
			g := func(i int) *maybe.Maybe[int] {
				return maybe.Some(i + 2)
			}
			mb := maybe.Some(1)

			lhs := mb.FlatMap(f).FlatMap(g)
			rhs := mb.FlatMap(func(i int) *maybe.Maybe[int] {
				return f(i).FlatMap(g)
			})

			o.Expect(lhs).To(o.Equal(rhs))
		})
	})
})
