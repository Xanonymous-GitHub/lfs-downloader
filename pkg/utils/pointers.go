package utils

type NilAble interface{}

// ToPointerType is a help method that returns the address of the `element`.
func ToPointerType[T any](element T) *T {
	return &element
}

// MayNilBy checks if `mayNil` is nil, if it is, return nil, otherwise, return the pointer of `element`.
func MayNilBy[T, M any](mayNil *M, element T) *T {
	if mayNil == nil {
		return nil
	}

	return ToPointerType(element)
}

// MayEmptyBy checks if `mayNil` is nil, if it is, return the empty type `T`, otherwise, return `element`.
func MayEmptyBy[T, M any](mayNil *M, element T) T {
	if mayNil == nil {
		return *new(T)
	}

	return element
}
