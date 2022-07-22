package utils

// Slice -> Slice of slices by maxItemsPerSlice.
func SplitSlice[T any](slice []T, maxItemsPerSlice int) [][]T {
	if slice == nil || maxItemsPerSlice < 1 {
		return nil
	}
	var items = make([][]T, 0)
	items = append(items, make([]T, 0))

	var delimIndex = 0
	for counter := range slice {
		if len(items[delimIndex]) > maxItemsPerSlice {
			items = append(items, make([]T, 0))
			delimIndex++
		}
		items[delimIndex] = append(items[delimIndex], slice[counter])
	}
	return items
}
