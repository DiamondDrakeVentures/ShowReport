package main

import (
	"fmt"
	"os"

	"github.com/DiamondDrakeVentures/ShowReport/migrator/cmd"
	"github.com/alexflint/go-arg"
)

const Name = "ShowReport Migrator"
const Version = "0.1.0"

func main() {
	args := cmd.NewCmdMigrate(Name, Version)
	arg.MustParse(args)

	fmt.Println(args.Version())

	os.Exit(args.Execute())
}
