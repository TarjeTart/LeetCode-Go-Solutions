func reverse(x int) int {

	res := 0
	for ; x != 0; x /= 10 {
		digit := x % 10
		if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && digit > math.MaxInt32%10) {
			return 0
		}
		if res < math.MinInt32/10 || (res == math.MinInt32/10 && digit < math.MinInt32%10) {
			return 0
		}
		res = res*10 + digit
	}
	return res

}
