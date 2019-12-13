package main

import (
	"olibs/times"
)

func main() {
	lg.Logf("Start %s", env.Name)

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
		lg.Logf("Run backup of set %v consisting of %v folders", idx, len(e.Folders))
		for _, folder := range e.Folders {
			b := BkpSet{
				Folder:       folder,
				OutputFolder: e.OutputFolder,
				Subfolder:    *argsSubfolder,
				Timestamp:    timestamp,
				Format:       e.Format,
			}
			b.TargetArchive = targetArchiveName(b)
			archive(b)
		}
	}
}
