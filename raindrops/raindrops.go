package raindrops

import "strconv"

func Convert(num int) string {
	var ret string
	if num%3 == 0 {
		ret += "Pling"
	}
	if num%5 == 0 {
		ret += "Plang"
	}

	if num%7 == 0 {
		ret += "Plong"
	}

	if ret == "" {
		return strconv.Itoa(num)
	}
	return ret
}
