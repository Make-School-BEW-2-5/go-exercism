package hamming

import "errors"

func Distance(a, b string) (int, error) {

	var distance int

	if len(a) != len(b) {
		return -1, errors.New("Inequal sequence lengeths")
	}
	for i, nucelotide := range a {
		if nucelotide != rune(b[i]) {
			distance += 1
		}
	}

	return distance, nil
}
