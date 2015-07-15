package functional_sets

import (
	"fmt"
)

type Set func(x int) bool
type pred func(x int) bool
type mapfn func(x int) int

// Checks if an element is part of the Set
func Contains(s Set, x int) bool {
	return s(x)
}

// Creates a Set with one element
func SingletonSet(x int) Set {
	return func(y int) bool {
		return (x == y)
	}
}

// Returns a Set with the union of the two Sets
func Union(s1 Set, s2 Set) Set {
	return func(y int) bool {
		return Contains(s1, y) || Contains(s2, y)
	}
}

// Returns an intersection of two Sets
func Intersection(s1 Set, s2 Set) Set {
	return func(y int) bool {
		return Contains(s1, y) && Contains(s2, y)
	}
}

// Returns a Set of elements in Set 1 and not in Set 2
func Diff(s1 Set, s2 Set) Set {
	return func(y int) bool {
		return Contains(s1, y) && !Contains(s2, y)
	}
}

// Return a Set of elements that satisfy a given predicate
func Filter(s Set, p pred) Set {
	return func(y int) bool {
		return Contains(s, y) && p(y)
	}
}

// We check a range of integers from -1000 to 1000
func Forall(s Set, p pred) bool {
	for x := -1000; x < 1000; x++ {
		if Contains(s, x) {
			if p(x) == false {
				return false
			}
		}
	}
	return true
}

// Define a map function. Since map is a keyword in go we cant use that.
func MapSet(s Set, m mapfn) Set {
	return func(y int) bool {
		p := func(elem int) bool { return m(elem) == y }
		return Exists(s, p)
	}
}

// Just a debug help tool
func PrintSet(s Set) {
	for x := -1000; x < 1000; x++ {
		if Contains(s, x) {
			fmt.Println(x)
		}
	}
}

// Does atleast one element exist that matches the predicate
func Exists(s Set, p pred) bool {
	negpred := func(y int) bool { return !p(y) }
	return !Forall(s, negpred)
}
