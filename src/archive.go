package main

import (
	"olibs/rx"
	"olibs/syslib"
)

func doArchive(foldername string, timestamp string) {
	ta := targetArchiveName(foldername, timestamp)
	lg.Logf("Archive folder %q -> %q", foldername, ta)
}

func targetArchiveName(folder string, timestamp string) (s string) {
	s = syslib.Pj(*argsOutfolder, timestamp)
	s = syslib.Pj(s, rx.Find(rxLib.AfterLastSlash, folder)+"."+*argsFormat)
	return
}
