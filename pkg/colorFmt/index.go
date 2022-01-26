package colorFmt

import "fmt"

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorPurple = "\033[35m"
	colorBlue   = "\033[36m"
)

func GreenFmt(str string) {
	fmt.Println(string(colorGreen), str, string(colorReset))
}

func RedFmt(str string) {
	fmt.Println(string(colorRed), str, string(colorReset))
}

func BlueFmt(str string) {
	fmt.Println(string(colorBlue), str, string(colorReset))
}

func YellowFmt(str string) {
	fmt.Println(string(colorYellow), str, string(colorReset))
}
