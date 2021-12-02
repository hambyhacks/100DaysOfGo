package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	for _, element := range slice {
		//Condition below checks 1.) if index is not a negative number, 2.) length of the slice is greater than index 3.) and if the element exists inside the slice in specified index.
		if !(index < 0) && len(slice) > index && element == slice[index] {
			return element, true
		}
	}
	return 0, false
	panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if !(index < 0) && len(slice) > index {
		slice[index] = value
		return slice
	} else {
		newSlice := append(slice, value)
		return newSlice
	}
	panic("Please implement the SetItem function")
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	var nullSlice []int
	if !(length < 0) {
		slice := make([]int, length)
		for i := range slice {
			slice[i] = value
		}
		return slice
	} else {
		return nullSlice
	}
	panic("Please implement the PrefilledSlice function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if !(index < 0) && !(cap(slice) < index) {
		copy(slice[index:], slice[index+1:])
		newSlice := slice[:len(slice)-1]
		return newSlice
	} else {
		return slice
	}
	panic("Please implement the RemoveItem function")
}
