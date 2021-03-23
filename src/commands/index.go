package commands

import "andy/src/utils"

type Command struct {
	WithArgs      bool
	AvailableArgs []string
}

func (c Command) Run(args []string) {
	if c.WithArgs == true && len(args) == 0 {
		utils.LessArgument()

		return
	}

	for _, arg := range c.AvailableArgs {
		if arg == args[0] {

			return
		}
	}

	utils.WrongArguments()
}
