package cmd

import "fmt"

type Command interface {
	Execute() (status int)
	Version() string
}

type GenericCmd struct {
	verStr string
}

func NewGenericCmd(name, version string) Command {
	return &GenericCmd{
		verStr: verStr(name, version),
	}
}

func (c GenericCmd) Execute() (status int) { return }
func (c GenericCmd) Version() string       { return c.verStr }

func verStr(name, version string) string {
	return fmt.Sprintf("%s v%s\n", name, version)
}
