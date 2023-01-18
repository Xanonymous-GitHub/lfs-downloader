package utils

// Contains checks if the collection contains an element equal to [element].
//
// This operation will check each element in order for being equal to
// `element`, unless it has a more efficient way to find an element
// equal to `element`.
//
// The equality used to determine whether `element` is equal to an element of
// the iterable defaults to the `==` of the element.
func Contains[E comparable](collection []E, element E) bool {
	for _, v := range collection {
		if v == element {
			return true
		}
	}

	return false
}

// Map modifies all elements of this collection by `toElement`.
//
// Returns a new array with elements that are created by
// calling `toElement` on each element of this `collection` in
// iteration order.
func Map[T, E any](collection []E, toElement func(e E) T) []T {
	length := len(collection)
	mappedCollection := make([]T, length)

	for i := 0; i < length; i++ {
		mappedCollection[i] = toElement(collection[i])
	}

	return mappedCollection
}

// Any Checks whether any element of collection satisfies `test`.
//
// Checks every element in the collection, and returns `true` if
// any of them make `test` return `true`, otherwise returns false.
func Any[E any](collection []E, test func(element E) bool) bool {
	length := len(collection)

	for i := 0; i < length; i++ {
		if test(collection[i]) {
			return true
		}
	}

	return false
}

// Every check whether every element of this collection satisfies `test`.
//
// Checks every element in the collection, and returns `false` if
// any of them make `test` return `false`, otherwise returns `true`.
func Every[E any](collection []E, test func(element E) bool) bool {
	length := len(collection)

	for i := 0; i < length; i++ {
		if !test(collection[i]) {
			return false
		}
	}

	return true
}

func Count[T comparable](collection []T) map[T]int {
	length := len(collection)
	m := make(map[T]int, length)

	for i := 0; i < length; i++ {
		m[collection[i]] += 1
	}
	return m
}

// CountBy would get the property `K` from data struct `E`
// accumulate every property `K` from a slice of data `E`
func CountBy[E any, K comparable](collection []E, getter func(E) K) map[K]int {
	mappedCollection := Map(collection, getter)
	length := len(collection)
	counter := make(map[K]int, length)

	for i := 0; i < length; i++ {
		counter[mappedCollection[i]] += 1
	}
	return counter
}

// ToPointers creates an array of pointers for each element in the collection.
func ToPointers[E any](collection []E) []*E {
	return Map(collection, func(e E) *E {
		return &e
	})
}

// SkipWhile returns a collection that skips leading elements while [test] is satisfied.
func SkipWhile[E any](collection []E, test func(element E) bool) []E {
	length := len(collection)
	filtered := make([]E, 0)

	for i := 0; i < length; i++ {
		if !test(collection[i]) {
			filtered = append(filtered, collection[i])
		}
	}

	return filtered
}
