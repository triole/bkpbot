package main

import (
	"olibs/rx"
	"olibs/syslib"

	"github.com/mholt/archiver/v3"
)

func archive(b BkpSet) {
	lg.Logf("Archive folder %q -> %q", b.Folder, b.TargetArchive)
	if *argsDebug == false {
		err := archiver.Archive([]string{b.Folder}, b.TargetArchive)
		if err != nil {
			lg.Logf("Error during compression %q -> %q", b.Folder, b.TargetArchive)
		}
	}
}

func targetArchiveName(b BkpSet) (s string) {
	s = b.OutputFolder
	if b.Subfolder != "" {
		s = syslib.Pj(b.OutputFolder, b.Subfolder)
	}
	s = syslib.Pj(s, b.Timestamp)
	s = syslib.Pj(s, rx.Find(rxLib.AfterLastSlash, b.Folder)+"."+b.Format)
	return
}
