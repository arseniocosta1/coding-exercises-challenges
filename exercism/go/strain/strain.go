package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](items []T, predicate func(T) bool) []T {
	if items == nil {
		return nil
	}
	kept := make([]T, 0, len(items))

	for _, item := range items {
		if predicate(item) {
			kept = append(kept, item)
		}
	}

	return kept
}

func Discard[T any](items []T, predicate func(T) bool) []T {
	return Keep(items, func(item T) bool { return !predicate(item) })
}
