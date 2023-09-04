package main

func PalindromoA(str string) bool {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return str == string(byte_str)
}

func PalindromoB(str string) bool {
	aux := ""
	size := len(str)
	for i, v := range str {
		if i > size {
			break
		}

		aux = aux + string(v)
	}

	return str == string(aux)
}
