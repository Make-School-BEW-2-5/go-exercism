package grains

import "errors"

func Square(n int) (uint64, error) {
	var amount uint64 = 1
	if n < 1 || n > 64 {
		return 0, errors.New("FUCK YOU")
	}
	for i := 1; i < n; i++ {
		amount *= 2
	}
	return amount, nil
}

func Total() uint64 {
	total := uint64(0)
	for i := 1; i < 65; i++ {
		value, _ := Square(i)
		total += value
	}
	return total
}
