package main

import (
	"olibs/rx"
	"olibs/syslib"
	"olibs/times"
	"strings"
)

func main() {
	lg.Logf("Start %s", appName)

	argparse()

	timestamp := times.Ts()
	conf := initConfig(*argsConfigfile)

	if *argsDebug == true {
		tf := "/tmp/bkpbot_debug.json"
		lg.Logf("Debug rich config saved to %q", tf)
		saveRichConfig(conf, tf)
	}

	// make backups
	lg.Logf("Process config having %v entries %+v", len(conf), conf)
	for _, bkpSet := range conf {
		if len(bkpSet.ToBackup) > 0 {
			outputFolder := bkpSet.Output.Folder
			if *argsSubfolder != "" {
				outputFolder = syslib.Pj(bkpSet.Output.Folder, *argsSubfolder)
			}

			bs := BkpSet{
				Timestamp: timestamp,
				ToBackup:  bkpSet.ToBackup,
				Output: Output{
					Name:   bkpSet.Output.Name,
					Folder: outputFolder,
					Format: bkpSet.Output.Format,
				},
			}
			bs.TargetArchive = targetArchiveName(bs)

			archive(bs)
			if *argsKeepLast > 0 {
				cleanUp(outputFolder, *argsKeepLast)
			}
		} else {
			lg.Logf(
				"Skip set because empty. Check if possible detection works."+
					"Settings: ToBackup %q, Outfolder %q, Format: %q",
				bkpSet.ToBackup, bkpSet.Output.Folder, bkpSet.Output.Format,
			)
		}
	}
	if *argsDebug == true {
		lg.Log("Nothing happened. Just ran in debug.")
	}
}

func targetArchiveName(bs BkpSet) (s string) {
	s = bs.Output.Folder
	s = syslib.Pj(s, bs.Timestamp)

	shortname := bs.Output.Name
	if shortname == "" {
		shortname = strings.Replace(
			rx.Find(rxlib.AfterLastSlash, bs.ToBackup[0]), ".", "_", -1,
		)
	}

	s = syslib.Pj(s, shortname+"."+bs.Output.Format)
	return
}
