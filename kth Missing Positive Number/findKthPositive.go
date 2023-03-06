func findKthPositive(arr []int, k int) int {

	if k < arr[0] {
		return k
	}

	k -= arr[0] - 1

	ll := 0
	lu := 1

	for lu < len(arr) && arr[lu]-arr[ll] <= k {
		k -= (arr[lu] - arr[ll]) - 1
		ll++
		lu++
	}

	if lu == len(arr) {
		return arr[len(arr)-1] + k
	}

	return arr[ll] + k

}
