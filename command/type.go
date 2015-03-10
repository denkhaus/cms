package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/denkhaus/cms/engine"
)

////////////////////////////////////////////////////////////////////////////////
func (c *Commander) NewTypeCommand() {
	c.Register(cli.Command{
		Name:  "type",
		Usage: "Manage cms types.",
		Subcommands: []cli.Command{
			{
				Name:  "create",
				Usage: "Create a new type.",
				Flags: []cli.Flag{
				//cli.StringFlag{"--name, -n", "", "Name of the type. Required", ""},
				},
				Action: func(ctx *cli.Context) {
					c.Execute(func(eng *engine.Engine) error {
						name := ctx.Args().First()
						if name == "" {
							return fmt.Errorf("Types::Name is not set. Value is mandatory.")
						}

						return eng.TypeCreate(name)
					}, ctx)
				},
			},
			{
				Name:  "list",
				Usage: "List all types.",
				Action: func(ctx *cli.Context) {
					c.Execute(func(eng *engine.Engine) error {
						return eng.TypeList()
					}, ctx)
				},
			},
			{
				Name:  "rm",
				Usage: "Remove type.",
				Action: func(ctx *cli.Context) {
					c.Execute(func(eng *engine.Engine) error {
						name := ctx.Args().First()
						if name == "" {
							return fmt.Errorf("Types::Name is not set. Value is mandatory.")
						}
						return eng.TypeRemove(name)
					}, ctx)
				},
			},
		},
	})
}
