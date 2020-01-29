package main

import (
	"olibs/times"
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
	for idx, e := range conf {
		if len(e.ToBackup) > 0 {
			lg.Logf("Run backup of set %v consisting of %v folders", idx, len(e.ToBackup))
			b := BkpSet{
				ToBackup:     e.ToBackup,
				OutputFolder: e.OutputFolder,
				Subfolder:    *argsSubfolder,
				Timestamp:    timestamp,
				Format:       e.Format,
			}
			b.TargetArchive = targetArchiveName(b)
			archive(b)
		} else {
			lg.Logf("Skip set because empty. Check if possible detection works. Settings: ToBackup %q, Outfolder %q, Format: %q", e.ToBackup, e.OutputFolder, e.Format)
		}
	}
	if *argsDebug == true {
		lg.Log("Nothing happened. Just ran in debug.")
	}
}
