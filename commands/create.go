package commands

import (
	"flag"
	"fmt"
)

type CreateCmd struct {
	Command *flag.FlagSet
	Format  *string
}

func init() {
	name := "create"
	createCmd := flag.NewFlagSet(name, flag.ExitOnError)
	Commands[name] = CreateCmd{
		Command: createCmd,
		Format:  createCmd.String("fmt", "json", "Formating style to create the Movies Database"),
	}
}

func (create CreateCmd) Execute(args []string) error {
	create.Command.Parse(args)
	fmt.Printf("Creating the Movies DB to %s format\n", *create.Format)
	return nil
}
