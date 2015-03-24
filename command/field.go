package command

import (
	"github.com/codegangsta/cli"
	"github.com/denkhaus/cms/engine"
)

////////////////////////////////////////////////////////////////////////////////
func (c *Commander) NewTypeCommand() {
	c.Register(cli.Command{
		Name:  "field",
		Usage: "Manage type fields.",
		Subcommands: []cli.Command{
			{
				Name:  "create",
				Usage: "Create a new type.",
				Flags: []cli.Flag{
				//cli.StringFlag{"--name, -n", "", "Name of the type. Required", ""},
				},
				Action: func(ctx *cli.Context) {
					c.Execute(func(eng *engine.Engine) error {
						return eng.FieldCreate(ctx.Args().Get(0), ctx.Args().Get(1))
					}, ctx)
				},
			},
			{
				Name:  "list",
				Usage: "List all fields of a given type.",
				Action: func(ctx *cli.Context) {
					c.Execute(func(eng *engine.Engine) error {
						return eng.FieldList(ctx.Args().Get(0))
					}, ctx)
				},
			},
			{
				Name:  "rm",
				Usage: "Remove field from type.",
				Action: func(ctx *cli.Context) {
					c.Execute(func(eng *engine.Engine) error {
						return eng.FieldRemove(ctx.Args().Get(0))
					}, ctx)
				},
			},
		},
	})
}
