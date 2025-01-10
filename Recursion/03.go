package recursion

// abc -> [”, a, ab, b, c, ac, abc, bc ]
func Subsequence(s string) []string {
	if len(s) == 0 {
		return []string{" "}
	}
	ch := s[0]
	str := Subsequence(s[1:])

	for ele := range str {
		str = append(str, string(ch)+str[ele])
	}

	return str
}

func KeyboardCombination(keyboard []string, nums string) []string {
	if len(nums) == 0 {
		return []string{"#"}
	}
	ch := nums[0]
	number := int(ch - '0')
	dial := KeyboardCombination(keyboard, nums[1:])
	resp := []string{}
	for ele := range dial {
		key := keyboard[number]
		for idx := range key {
			resp = append(resp, string(key[idx])+dial[ele])
		}
	}
	return resp
}
