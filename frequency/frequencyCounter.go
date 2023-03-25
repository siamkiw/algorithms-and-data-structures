package frequency

func Same(arr1 []int, arr2 []int) bool {

	if len(arr1) != len(arr2) {
		return false
	}

	freq1 := make(map[int]int)
	freq2 := make(map[int]int)

	for _, v := range arr1 {
		freq1[v]++
	}

	for _, v := range arr2 {
		freq2[v]++
	}

	for k, v := range freq1 {
		if freq2[k*k] != v {
			return false
		}
	}

	return true

}

func Anagrams(arr1 string, arr2 string) bool {

	if len(arr1) != len(arr2) {
		return false
	}

	freq1 := make(map[string]int)
	freq2 := make(map[string]int)

	for _, v := range arr1 {
		freq1[string(v)]++
	}

	for _, v := range arr2 {
		freq2[string(v)]++
	}

	for k, v := range freq1 {
		if freq2[k] != v {
			return false
		}
	}

	return true
}

func containsDuplicate(nums []int) bool {
	data := map[int]bool{}

	for _, value := range nums {
		if data[value] {
			return false
		}

		data[value] = true
	}

	return true
}
