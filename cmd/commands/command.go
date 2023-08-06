package commands

import (
	"os"
)

type ICommand interface {
	Name() string
	Handle()
}

func Register(command ICommand) {
	// fmt.Println(len(os.Args), command.Name(), os.Args[1])

	//if there are more argument/s (>1) after filename
	//and if tht next arg matches the command name,
	if len(os.Args) > 1 && command.Name() == os.Args[1] {
		command.Handle()
	}

}
