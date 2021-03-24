package utils

import "fmt"

func PanicMessage() {
	panic("🆘Opps.. Something gonna wrong..")
}

func LessArgument() {
	fmt.Println("Not anought arguments for this command😠!")
}

func WrongArguments() {
	fmt.Println("Unknown arguments🧐!\nPlease, type --help for more details!")
}

func Greetings() {
	fmt.Println("Hi! I'm Andy🐼!\nRun --help or -h to watch my commands.\nSee You later🛀!")
}
