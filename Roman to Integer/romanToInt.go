func romanToInt(s string) int {
    var toReturn = 0
	m := make(map[string]int)

	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

    if len(s) == 1 {
		return m[s]
	}

	for i := 0; i < len(s)-1; i++ {
		if m[string(s[i+1])] > m[string(s[i])] {
			toReturn += m[string(s[i+1])] - m[string(s[i])]
			i++
		} else {
			toReturn += m[string(s[i])]
		}
		if i == len(s)-2 {
			toReturn += m[string(s[i+1])]
		}
	}

	return toReturn
}
