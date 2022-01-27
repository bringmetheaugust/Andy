package colorFmt

import "fmt"

func print(str string, color string) {
	fmt.Println(string(color), str, string("\033[0m"))
}

func GreenFmt(str string) {
	print(str, "\033[32m")
}

func RedFmt(str string) {
	print(str, "\033[31m")
}

func BlueFmt(str string) {
	print(str, "\033[36m")
}

func YellowFmt(str string) {
	print(str, "\033[33m")
}
