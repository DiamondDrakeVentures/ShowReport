package cmd

import (
	"fmt"

	"github.com/DiamondDrakeVentures/ShowReport/port/migrate"
)

type CmdMigrate struct {
	SrcFile    string `arg:"--input,-i" help:"source file"`
	SrcFormat  string `arg:"--input-format,-s" help:"format of source file"`
	DstFile    string `arg:"--output,-o" help:"target file"`
	DstFormat  string `arg:"--output-format,-f" default:"5x" help:"format of target file"`
	UpdateUser string `arg:"--user,-u" default:"migrator@diamonddrake.co" help:"user attributed for the migrated data"`

	verStr string
}

func NewCmdMigrate(name, version string) Command {
	return &CmdMigrate{
		verStr: verStr(name, version),
	}
}

func (c CmdMigrate) Execute() (status int) {
	var migrator migrate.Migrator

	if c.DstFormat == "5x" {
		migrator = migrate.To5x(c.SrcFile, c.DstFile, c.SrcFormat, c.UpdateUser)
	} else {
		fmt.Printf("Unknown output format: %s\n", c.DstFormat)
		return 1
	}

	fmt.Printf("Migrating %s (%s) to %s (%s)...\n",
		c.SrcFile, c.SrcFormat, c.DstFile, c.DstFormat)

	err := migrator.Execute()
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return
}

func (c CmdMigrate) Version() string {
	return c.verStr
}
