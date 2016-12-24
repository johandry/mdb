package commands

import (
	"flag"
	"fmt"
)

type ExportCmd struct {
	Command *flag.FlagSet
	Dest    *string
}

func init() {
	name := "export"
	exportCmd := flag.NewFlagSet(name, flag.ExitOnError)
	Commands[name] = ExportCmd{
		Command: exportCmd,
		Dest:    exportCmd.String("dest", "aws", "Location where to export the Movies Database"),
	}
}

func (export ExportCmd) Execute(args []string) error {
	export.Command.Parse(args)
	fmt.Printf("Exporting the Movies DB to %s\n", *export.Dest)
	return nil
}
