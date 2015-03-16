package command

import (
	"fmt"
	"regexp"

	"github.com/codegangsta/cli"
	"github.com/denkhaus/cms/engine"
)

var types = map[string]struct{}{
	{"string", struct{}{}},
}

////////////////////////////////////////////////////////////////////////////////
func MatchFieldType(tf string) bool {
	_, ok := types[tf]
	return ok
}

////////////////////////////////////////////////////////////////////////////////
func MatchTypeField(tp string) bool {
	r, _ := regexp.Compile("[a-zA-Z0-9]+:[a-zA-Z0-9]+")
	return r.Match(tf)
}

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
						tf := ctx.Args().First()
						if !MatchTypeField(tf) {
							return fmt.Errorf("Fields::Type:Field combination not valid. Use e.g. 'picture:url'.")
						}
						tp := ctx.Args().Second()
						if !MatchFieldType(tp) {
							return fmt.Errorf("Fields::Field type '%s' not valid.", tp)
						}
						return eng.FieldCreate(tf, tp)
					}, ctx)
				},
			},
			{
				Name:  "list",
				Usage: "List all fields of a given type.",
				Action: func(ctx *cli.Context) {
					c.Execute(func(eng *engine.Engine) error {
						typeName := ctx.Args().First()
						if typeName == "" {
							return fmt.Errorf("Fields::Type Name is not set. Value is mandatory.")
						}
						return eng.FieldList(typeName)
					}, ctx)
				},
			},
			{
				Name:  "rm",
				Usage: "Remove field from type.",
				Action: func(ctx *cli.Context) {
					c.Execute(func(eng *engine.Engine) error {
						tf := ctx.Args().First()
						if !MatchTypeField(tf) {
							return fmt.Errorf("Fields::Type:Field combination not valid. Use e.g. 'picture:url'.")
						}
						return eng.FieldRemove(tf)
					}, ctx)
				},
			},
		},
	})
}
