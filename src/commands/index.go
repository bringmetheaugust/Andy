package commands

import "andy/src/utils"

type ICommand interface {
	Start(args []string)
	Run(args []string)
}

type Command struct {
	WithArgs      bool
	AvailableArgs []string
}

func (c Command) Start(args []string) {
	// check if command require arguments
	if c.WithArgs == true && len(args) == 0 {
		utils.LessArgument()

		return
	}

	// check existed available arguments
	for _, arg := range c.AvailableArgs {
		if arg == args[0] {
			c.Run(args)

			return
		}
	}

	utils.WrongArguments()
}
