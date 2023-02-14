func addBinary(a string, b string) string {
    var toReturn string = ""
	var carry = 0
	var first []int
	var second []int
	if len(b) > len(a) {
		var tmp = a
		a = b
		b = tmp
	}
	for i := len(a) - 1; i >= 0; i-- {
		j, err := strconv.Atoi(string(a[i]))
		if err != nil {
			panic(err)
		}
		first = append(first, j)
	}
	for i := len(b) - 1; i >= 0; i-- {
		j, err := strconv.Atoi(string(b[i]))
		if err != nil {
			panic(err)
		}
		second = append(second, j)
	}

	for i := 0; i < len(second); i++ {
		var tmp = first[i] + second[i] + carry
		if tmp == 0 || tmp == 1 {
			carry = 0
			toReturn += strconv.Itoa(tmp)
		} else if tmp == 2 {
			carry = 1
			toReturn += "0"
		} else {
			carry = 1
			toReturn += "1"
		}
	}

	for i := len(second); i < len(a); i++ {
		var tmp = first[i] + carry
		if tmp == 0 || tmp == 1 {
			carry = 0
			toReturn += strconv.Itoa(tmp)
		} else {
			carry = 1
			toReturn += "0"
		}
	}

	if carry == 1 {
		toReturn += "1"
	}

	var tmp = ""

	for i := len(toReturn) - 1; i >= 0; i-- {
		tmp += string(toReturn[i])
	}

	return tmp
}
