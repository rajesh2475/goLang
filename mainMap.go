package main

import "fmt"

type mapType map[string]string

func main() {

	// before adding new value to map it must be initilized('https://yourbasic.org/golang/gotcha-assignment-entry-nil-map/')
	// Method 1
	// colors := map[string]string{}

	// Method 2

	colors := make(mapType)

	// // Method 3
	// colors := map[string]string{
	// 	"red": "Rad",
	// }

	colors["red"] = "Red"
	colors["blue"] = "Blue"
	// fmt.Println(colors)
	// delete(colors, "red")
	// fmt.Println(colors)
	colors.println()
}

func (m mapType) println() {
	for key, value := range m {
		fmt.Println(value)
		fmt.Println(m[key])
	}
}
