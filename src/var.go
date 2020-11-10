package main

import (
	"os"
	"time"

	"github.com/triole/bkpbot/env"
	"github.com/triole/bkpbot/logging"
	"github.com/triole/bkpbot/rx"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	BUILDTAGS      string
	appName        = "bkpbot"
	appMainVersion = "0.1"
	appDescription = "Backup bot archives folders and makes what it says"
	defaultLogFile = "/tmp/bkpbot.log"

	lg = logging.Init(defaultLogFile)
	ts = time.Now()

	app            = kingpin.New(appName, appDescription)
	argsConfigfile = app.Arg("config", "config file to read the setting from").Required().String()
	argsSubfolder  = app.Flag("subfol", "subfolder created in output directory, used for daily, weekly etc.").Short('s').Default("").String()
	argsKeepLast   = app.Flag("keep", "keep last n backups, zero keeps all").Short('k').Default("0").Int()
	argsLogfile    = app.Flag("logfile", "logfile which will be written").Short('l').Default(defaultLogFile).String()
	argsDebug      = app.Flag("debug", "debug mode, just print no action").Short('d').Default("false").Bool()

	rxLib = rx.InitLib()
)

func argparse() {
	e := env.Env{
		Name:        appName,
		MainVersion: appMainVersion,
		Description: appDescription,
	}
	app.Version(env.MakeInfoString(e, env.ParseBuildtags(BUILDTAGS)))
	app.HelpFlag.Short('h')
	app.VersionFlag.Short('V')
	kingpin.MustParse(app.Parse(os.Args[1:]))
	if *argsLogfile != e.Logfile {
		lg = logging.Init(*argsLogfile)
	}
}
