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

		for _, f := range e.Folders {
			doArchive(f, timestamp)
		}
	}

	// doBackup(conf.Folders)

	// // do backup
	// for idx, folders := range folders2Backup {
	// 	lg.Logf("Run backup of set %v consisting of %v folders", idx, len(folders))
	//
	// 	// iterate over folders
	// 	for _, fol := range folders {
	// 		backup(fol)
	// 	}
	// }
}
