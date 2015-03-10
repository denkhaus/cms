package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/denkhaus/cms/command"
	"github.com/denkhaus/cms/engine"
	"github.com/denkhaus/tcgl/applog"
)

func main() {
	app := cli.NewApp()
	app.Name = "cms"
	app.Version = AppVersion
	app.Usage = "A cli content management system."
	app.Flags = []cli.Flag{
	//		cli.StringFlag{"host, H", "localhost", "Host to connect to.", ""},
	//		cli.IntFlag{"port, P", 993, "Port number to connect to.", ""},
	//		cli.StringFlag{"user, u", "", "Your username at host", ""},
	//		cli.StringFlag{"pass, p", "", "Your IMAP password. For security reasons prefer IMAP_PASSWORD='yourpassword'", "IMAP_PASSWORD"},
	//		cli.BoolFlag{"reset, r", "Clear database before run", ""},
	}

	cnf, err := engine.NewConfig(".cmsrc")
	if err != nil {
		applog.Errorf("config::%s", err.Error())
		return
	}

	if cmdr, err := command.NewCommander(cnf, app); err != nil {
		applog.Errorf("ccommand::%s", err.Error())
		return
	} else {
		cmdr.Run(os.Args)
	}
}
