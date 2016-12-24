package commands

import (
	"flag"
	"os"
)

type Command interface {
	Execute(args []string) error
}

var Commands = make(map[string]Command)

func Run() {
	if len(os.Args) == 1 {
		flag.Parse()
	} else {
		for name, cmd := range Commands {
			if os.Args[1] == name {
				cmd.Execute(os.Args[2:])
			}
		}
	}
}
