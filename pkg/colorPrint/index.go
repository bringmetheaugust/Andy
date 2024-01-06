package colorPrint

import "fmt"

func print(str string, color string) {
	fmt.Printf("%v%v%v", color, str, string("\033[0m"))
}

func GreenPrint(str string) {
	print(str, "\033[32m")
}

func RedPrint(str string) {
	print(str, "\033[31m")
}

func BluePrint(str string) {
	print(str, "\033[36m")
}

func YellowPrint(str string) {
	print(str, "\033[33m")
}
