func strStr(haystack string, needle string) int {

	for i := 0; i < len(haystack); i++ {

		if haystack[i] == needle[0] {
			count := 1
			for count < len(needle) && i+count < len(haystack) && haystack[i+count] == needle[count] {
				count++
			}
			if count == len(needle) {
				return i
			}
		}

	}

	return -1

}
