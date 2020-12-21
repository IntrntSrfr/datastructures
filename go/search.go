package _go

func BinarySearch(l []int, t int) int {
	left, right, mid := 0, len(l)-1, 0

	for left < right {
		mid = (left + right) / 2
		if t > l[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if l[left] == t {
		return left
	}
	return -1
}
