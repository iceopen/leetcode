package _87_first_unique_character_in_a_string

func firstUniqChar(s string) int {
	bytearr := make([]rune, 1024)
	for _, v := range s {
		bytearr[v]++
	}
	for k, v := range s {
		if bytearr[v] == 1 {
			return k
		}
	}
	return -1
}
