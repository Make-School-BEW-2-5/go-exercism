package sublist

import "fmt"

func Sublist(listOne []int, listTwo []int) string {
	if len(listOne) == 0 {
		if len(listTwo) == 0 {
			return "equal"
		}
		return "sublist"
	}
	if len(listTwo) == 0 {
		return "superlist"
	}
	fullyEqual := true
	counter := 0
	for i, num := range listOne {
		if num == listTwo[counter] {
			if counter == len(listTwo)-1 {
				if counter == i {
					return "equal"
				}
				fmt.Println(counter, i)
				return "sublist"
			}
			counter++
		} else {
			counter = 0
			fullyEqual = false
		}
	}
	if fullyEqual {
		return "superlist"
	}
	return "unequal"
}
