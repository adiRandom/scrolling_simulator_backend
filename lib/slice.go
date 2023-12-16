package lib

func SafeGet[T any](s []T, i int) (*T, error) {
	if i >= len(s) {
		return nil, Error{"Index out of bounds", ""}
	}

	return &s[i], nil
}

func SafeGetRange[T any](s []T, i, j int) ([]T, error) {
	if i > len(s) || j > len(s) {
		return nil, Error{"Index out of bounds", ""}
	}

	return s[i:j], nil
}

func Contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func PointerContains[T comparable](s []*T, e T) bool {
	for _, a := range s {
		if *a == e {
			return true
		}
	}
	return false
}

func ContainsFunc[T any](s []T, f func(T) bool) bool {
	for _, a := range s {
		if f(a) {
			return true
		}
	}
	return false
}
