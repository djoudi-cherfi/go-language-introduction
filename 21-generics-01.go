package main

import "fmt"

// Generics 1.18+
// Types to be parameters
// Type parameters / Type inference / Type set

//  ---------------------------------- Type parameters

// func genericsOne() {
// 	fmt.Println(min[float32](0.5, 0.7))
// 	fmt.Println(min[int32](2, 1))
// }

// func min[T int32 | float32](x, y T) T {
// 	return x
// }

// func genericsOne() {
// 	fmt.Println(min[float32](0.5, 0.7))
// 	fmt.Println(min[int32](2, 1))
// }

// func min[T int32 | float32](x, y T) (T, T) {
// 	return x, y
// }

//  ---------------------------------- Type inference

// func genericsOne() {
// 	type f float64
// 	var foo f = 3.14
// 	fmt.Println(min(foo, 0.7))
// }

// func min[T ~float64](x, y T) T {
// 	return y
// }

// ---------------------------------- Type set
type Number interface {
	int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | ~float64
}

func genericsOne() {
	type f float64
	var foo f = 0.5
	fmt.Println(min(0.7, foo))
}

func min[T Number](x, y T) T {
	if x < y {
		return x
	}

	return y
}
