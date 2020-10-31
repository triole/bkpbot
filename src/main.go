package main

import (
	"olibs/rx"
	"olibs/times"
	"sort"
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
	lg.Logf("Process config having %v entries", len(conf))

	// make alpha iterator
	var keys []string
	for k, _ := range conf {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, name := range keys {
		bkpSet := conf[name]
		if len(bkpSet.ToBackup) > 0 {
			outputFolder := bkpSet.OutputFolder
			if *argsSubfolder != "" {
				outputFolder = pj(bkpSet.OutputFolder, *argsSubfolder)
			}

			bs := tBkpSet{
				Name:         name,
				Timestamp:    timestamp,
				ToBackup:     bkpSet.ToBackup,
				OutputName:   bkpSet.OutputName,
				OutputFolder: outputFolder,
				OutputFormat: bkpSet.OutputFormat,
				RunBefore:    bkpSet.RunBefore,
				RunAfter:     bkpSet.RunAfter,
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
				bkpSet.ToBackup, bkpSet.OutputFolder, bkpSet.OutputFormat,
			)
		}
	}
	if *argsDebug == true {
		lg.Log("Nothing happened. Just ran in debug.")
	}
}

func targetArchiveName(bs tBkpSet) (s string) {
	s = bs.OutputFolder
	s = pj(s, bs.Timestamp)

	shortname := bs.OutputName
	if shortname == "" {
		shortname = strings.Replace(
			rx.Find(rxlib.AfterLastSlash, bs.ToBackup[0]), ".", "_", -1,
		)
	}

	s = pj(s, shortname+"."+bs.OutputFormat)
	return
}
