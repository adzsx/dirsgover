package utils

func FilterChar(str string, char string, before bool) string {
	var final string

	for index, element := range str {
		if before {
			if string(element) != char {
				final += string(element)
			} else {
				return final
			}
		} else {

			if string(element) == char {
				final += str[index+1:]
			}
		}

	}

	return final
}