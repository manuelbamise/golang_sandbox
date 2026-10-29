package main

import (
	"fmt"
)

func main() {

	fmt.Println(Decode("MMMCMXCI"))
}

func Decode(roman string) int {
	figureMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	fmt.Println(figureMap)
	var output int

	for i := 0; i < len(roman); i++ {
		if i+1 < len(roman) && figureMap[rune(roman[i])] < figureMap[rune(roman[i+1])] {
			output -= figureMap[rune(roman[i])]
		} else {
			output += figureMap[rune(roman[i])]
		}
	}

	// if len(roman)%2 == 0 {
	// 	for i := 0; i < len(roman); i += 2 {
	// 		fmt.Println(string(roman[i]))
	// 		// stringRoman := string(roman[i])
	// 		if figureMap[rune(roman[i])] >= figureMap[rune(roman[i+1])] {
	// 			output += figureMap[rune(roman[i])] + figureMap[rune(roman[i+1])]
	//
	// 		}
	//
	// 		if figureMap[rune(roman[i])] < figureMap[rune(roman[i+1])] {
	// 			output += figureMap[rune(roman[i+1])] - figureMap[rune(roman[i])]
	//
	// 		}
	//
	// 	}
	// }
	//
	// if len(roman)%2 != 0 {
	// 	output += figureMap[rune(roman[0])]
	// 	roman = roman[1:]
	// 	for i := 0; i < len(roman); i += 2 {
	// 		fmt.Println(string(roman[i]))
	// 		// stringRoman := string(roman[i])
	// 		if figureMap[rune(roman[i])] >= figureMap[rune(roman[i+1])] {
	// 			output += figureMap[rune(roman[i])] + figureMap[rune(roman[i+1])]
	//
	// 		}
	// 		if figureMap[rune(roman[i])] < figureMap[rune(roman[i+1])] {
	//
	// 			output += figureMap[rune(roman[i+1])] - figureMap[rune(roman[i])]
	//
	// 		}
	//
	// 	}
	// }
	//
	// for i := 0; i < len(roman); i++ {
	// 	if i < len(roman)-1 && figureMap[rune(roman[i])] < figureMap[rune(roman[i+1])] {
	// 		output -= figureMap[rune(roman[i])]
	// 	} else {
	// 		output += figureMap[rune(roman[i])]
	// 	}
	// }

	return output

}
