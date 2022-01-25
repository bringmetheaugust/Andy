package colorFmt

import "fmt"

const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
const colorBlue = "\033[34m"
const colorPurple = "\033[35m"
const colorCyan = "\033[36m"
const colorWhite = "\033[37m"

func GreenFmt(str string) {
	fmt.Println(string(colorGreen), str)
}

func RedFmt(str string) {
	fmt.Println(string(colorRed), str)
}

func BlueFmt(str string) {
	fmt.Println(string(colorBlue), str)
}
