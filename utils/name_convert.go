package utils


//转驼峰命名【"camelCase"的样式】
func ToHumpCase(str string) string {
	str = startToLower(str)
	asciiArray := []rune(str)
	var result []rune
	toCapital := false
	for i := 0; i < len(asciiArray); i++ {
		if asciiArray[i] < 65 || asciiArray[i] > 122 || (asciiArray[i] < 97 && asciiArray[i] > 90) {
			toCapital = true
		} else {
			if toCapital {
				if asciiArray[i] > 96 && asciiArray[i] < 123 {
					asciiArray[i] -= 32
				}
				toCapital = false
			}
			result = append(result, asciiArray[i])
		}
	}
	return string(result)
}

//转Json名称【"snake_case"的样式】
func ToSnakeCase(str string) string {
	str = startToLower(str)
	return CapitalToLowRodLower(str)
}

//大写字母转横杠加小写字母
func CapitalToRodLower(str string) string {
	return capitalToRodLowerWithFlag(str, '-')
}

//大写字母转底杠加小写字母
func CapitalToLowRodLower(str string) string {
	return capitalToRodLowerWithFlag(str, '_')
}

//大写字母转指定标记加小写字母
func capitalToRodLowerWithFlag(str string, flag rune) string {
	if len(str) == 0 {
		return str
	}
	asciiArray := []rune(str)
	var result []rune
	for i := 0; i < len(asciiArray); i++ {
		if asciiArray[i] > 64 && asciiArray[i] < 91 {
			asciiArray[i] += 32
			result = append(result, flag)
		}
		result = append(result, asciiArray[i])
	}
	return string(result)
}

//首字母转大写
func startToCapital(str string) string {
	if len(str) == 0 {
		return str
	}
	c := rune(str[0])
	if c > 96 && c < 123 {
		word := c - 32
		if len(str) == 1 {
			return string(word)
		}
		return string(word) + str[1:]
	}
	return str
}

//首字母转小写
func startToLower(str string) string {
	if len(str) == 0 {
		return str
	}
	c := rune(str[0])
	if c > 64 && c < 91 {
		word := c + 32
		if len(str) == 1 {
			return string(word)
		}
		return string(word) + str[1:]
	}
	return str
}