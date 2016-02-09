// Functions are central in Go.

// Go has built-in support for multiple return values.
// This feature is used often in idiomatic Go, for example
// to return both result and error values from a function.

// Variadic functions can be called with any number of
// trailing arguments. For example, `fmt.Println` is a
// common variadic function.

// Go supports anonymous functions, which can form closures.
// Anonymous functions are useful when you want to define a
// function inline without having to name it.

// Go supports recursive functions

package main

import "fmt"

// Here's a function that takes two `int`s and returns
// their sum as an `int`.
func plus(a int, b int) int {

	// Go requires explicit returns, i.e. it won't
	// automatically return the value of the last
	// expression.
	return a + b
}

// When you have multiple consecutive parameters of
// the same type, you may omit the type name for the
// like-typed parameters up to the final parameter that
// declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

// The `(int, int)` in this function signature shows that
// the function returns 2 `int`s.
func vals() (int, int) {
	return 3, 7
}

// Here's a function that will take an arbitrary number
// of `ints` as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("total:", total)
}

// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function closes over the variable `i` to
// form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// This `fact` function calls itself until it reaches the
// base case of `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func functions() {

	// Call a function just as you'd expect, with
	// `name(args)`.
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)

	// Here we use the 2 different return values from the
	// call with _multiple assignment_.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// If you only want a subset of the returned values,
	// use the blank identifier `_`.
	_, c := vals()
	fmt.Println(c)

	// Variadic functions can be called in the usual way
	// with individual arguments.
	sum(1, 2)
	sum(1, 2, 3, 4, 5)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using
	// `func(slice...)` like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// We call `intSeq`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.
	nextInt := intSeq()

	// See the effect of the closure by calling `nextInt`
	// a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println(fact(7))
}
