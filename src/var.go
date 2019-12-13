package main

import (
	"olibs/environment"
	"olibs/logging"
	"olibs/rx"
	"olibs/syslib"
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
	argsOutfolder  = app.Flag("outdir", "output directory").Short('o').Default(env.Curdir).String()
	argsFormat     = app.Flag("format", "archive format, currently only txz is supported").Short('f').Default("txz").String()
	argsConfigfile = app.Flag("config", "config file to read the setting from").Short('c').Default(syslib.Pj(env.Scriptdir, env.Name+".yml")).String()
	argsLogfile    = app.Flag("logfile", "logfile which will be written").Short('l').Default(env.Logfile).String()
	argsDebug      = app.Flag("debug", "debug mode, just print no action").Short('d').Default("false").Bool()

	rxLib = rx.InitLib()
)

func argparse() {
	app.Version(env.AppInfoString)
	app.HelpFlag.Short('h')
	app.VersionFlag.Short('v')
	kingpin.MustParse(app.Parse(os.Args[1:]))
	// maybe reinitialize logging if different logfile was set by arg call
	if *argsLogfile != env.Logfile {
		lg = logging.Init(*argsLogfile)
	}
}
