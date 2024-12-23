package iters

import "fmt"

func ExampleFromTo() {
	for i, n := range FromTo(3, 7) {
		fmt.Printf("%d %d\n", i, n)
	}

	// Output:
	// 0 3
	// 1 4
	// 2 5
	// 3 6
}

// func ExampleFromTo_Reverse() {
//
//	for i, n := range FromTo(7, 3) {
//		fmt.Printf("%d %d\n", i, n)
//	}
//
//	// Output:
//	// 0 7
//	// 1 6
//	// 2 5
//	// 3 4
//}

func ExampleFromToInclusive() {
	for i, n := range FromToInclusive(3, 7) {
		fmt.Printf("%d %d\n", i, n)
	}

	// Output:
	// 0 3
	// 1 4
	// 2 5
	// 3 6
	// 4 7
}

func ExampleFromStepTo() {
	for i, n := range FromStepTo(1.0, 1.5, 7.0) {
		fmt.Printf("%d %.1f\n", i, n)
	}

	// Output:
	// 0 1.0
	// 1 2.5
	// 2 4.0
	// 3 5.5

}
