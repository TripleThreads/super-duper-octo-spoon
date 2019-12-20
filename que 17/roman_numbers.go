package main

import (
	"math"
)

var toRoman = map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	50: "L", 100: "C", 500: "D", 1000: "M"}

var toNumeric = map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	"L": 50, "C": 100, "D": 500, "M": 1000}

func encode(num int) (string, bool) {
	if num >= 4000 || num <= 0 {
		return "", false
	}
	str := ""

	digit := int(math.Log10(float64(num)))

	for num > 0 {

		r := int(float64(num) / math.Pow(10, float64(digit)))

		or := int(math.Pow(10, float64(digit)))

		if val, ok := toRoman[or]; ok && r != 4 && r != 9 {
			if or*r <= 10 || or*r == 50 || or*r == 500 {
				str += toRoman[r*or]
			}
			if r > 5 {
				r -= 5
			}
			for i := 0; i < r && r != 5; i++ {
				str += val
			}
		} else {
			str += toRoman[or]
			str += toRoman[or+(r*or)]
		}
		num = num - (r * or)
		digit--
	}
	return str, len(str) != 0
}

func decode(str string) (int, bool) {
	converted := 0
	lastVal := -1

	for _, i := range str {
		if _, ok := toNumeric[string(i)]; !ok {
			return 0, false
		}
		current := toNumeric[string(i)]

		if lastVal != -1 && lastVal < current && (current-lastVal)%4 != 0 && (current-lastVal)%9 != 0 {
			return 0, false
		}
		if lastVal != -1 && lastVal < current && ((current-lastVal)%4 == 0 || (current-lastVal)%9 == 0) {
			converted -= lastVal
			converted += current - lastVal
		} else {
			converted += current
		}
		lastVal = current
	}
	return converted, true
}

func main() {
	val, ok := encode(3555)
	println(val, ok)

	dec, dok := decode("CCXVI")
	println(dec, dok)
}
