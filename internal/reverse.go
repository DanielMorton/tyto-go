package internal

func Reverse[K any](array []K) {
	left, right := 0, len(array)-1
	for left < right {
		array[left], array[right] = array[right], array[left]
		left++
		right--
	}
}
