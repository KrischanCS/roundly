package iters

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// FromTo creates an iterator returning the values from start to end exclusive (Half open range).
//
// The iterator returns two values each call, the first one being the 'index' starting at 0,
// the second one is the value of the sequence,
// so this behaves the same as doing for range over a slice/array holding the same values.
func FromTo(start, endNotIncluded int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i := 0; i < endNotIncluded-start; i++ {
			if !yield(i, start+i) {
				return
			}
		}
	}
}

// FromToInclusive creates an iterator returning the values from start to end inclusive (Closed range).
//
// The iterator returns two values each call, the first one being the 'index' starting at 0,
// the second one is the value of the sequence,
// so this behaves the same as doing for range over a slice/array holding the same values.
func FromToInclusive(start, endIncluded int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i := range endIncluded - start + 1 {
			if !yield(i, start+i) {
				return
			}
		}
	}
}

type Number interface {
	constraints.Integer | constraints.Float
}

// FromStepTo creates an iterator returning the values from start to end,
// where the value is increased by step every call and end will not be included, even if it is met exactly.
//
// Panics in the following cases:
//   - step == 0
//   - start < end && step < 0
//   - start > end && step > 0
//
// The iterator returns two values each call, the first one being the 'index' starting at 0,
// the second one is the value of the sequence,
// so this behaves the same as doing for range over a slice/array holding the same values.
func FromStepTo[T Number](start, step, endNotIncluded T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		v := start
		for i := 0; v < endNotIncluded; i, v = i+1, v+step {
			if !yield(i, v) {
				return
			}
		}
	}
}
