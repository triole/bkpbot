package main

import (
	"olibs/rx"
	"olibs/syslib"

	"github.com/mholt/archiver/v3"
)

func startBackup(foldername string, timestamp string) {
	targetArchive := targetArchiveName(foldername, timestamp)
	lg.Logf("Archive folder %q -> %q", foldername, targetArchive)
	if *argsDebug == false {
		archive(foldername, targetArchive)
	}
}

func targetArchiveName(folder string, timestamp string) (s string) {
	s = syslib.Pj(*argsOutfolder, timestamp)
	s = syslib.Pj(s, rx.Find(rxLib.AfterLastSlash, folder)+"."+*argsFormat)
	return
}

func archive(foldername string, targetArchive string) {
	err := archiver.Archive([]string{foldername}, targetArchive)
	if err != nil {
		lg.Logf("Error during compression %q -> %q", foldername, targetArchive)
	}
}
