package main

import (
	"olibs/environment"
	"olibs/logging"
	"olibs/rx"
	"os"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	BUILDTAGS      string
	appName        = "bkpbot"
	appMainversion = "0.1"
	appDescription = "Backup bot archives folders and makes what it says"
	env            = environment.Init(appName, appMainversion, appDescription, BUILDTAGS)

	lg = logging.Init(env.Logfile)
	ts = time.Now()

	app            = kingpin.New(appName, appDescription)
	argsConfigfile = app.Arg("config", "config file to read the setting from").Required().String()
	argsSubfolder  = app.Flag("subfol", "subfolder created in output directory, used for daily, weekly etc.").Short('s').Default("").String()
	argsKeepLast   = app.Flag("keep", "keep last n backups, zero keeps all").Short('k').Default("0").Int()
	argsLogfile    = app.Flag("logfile", "logfile which will be written").Short('l').Default(env.Logfile).String()
	argsDebug      = app.Flag("debug", "debug mode, just print no action").Short('d').Default("false").Bool()

	rxlib = rx.InitLib()
)

func argparse() {
	app.Version(env.AppInfoString)
	app.HelpFlag.Short('h')
	app.VersionFlag.Short('V')
	kingpin.MustParse(app.Parse(os.Args[1:]))
	if *argsLogfile != env.Logfile {
		lg = logging.Init(*argsLogfile)
	}
}
