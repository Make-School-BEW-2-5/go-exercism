package reverse

func Reverse(str string) string {
	var ret []byte
	for i := len(str); i >= 0; i -= 1 {
		ret = append(ret, str[i])
	}
	return string(ret)
}
