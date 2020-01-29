package main

import (
	"compress/flate"
	"olibs/rx"
	"olibs/syslib"
	"strings"

	"github.com/mholt/archiver"
)

func archive(b BkpSet) {
	lg.Logf("Archive folder(s) %q -> %q", b.ToBackup, b.TargetArchive)
	if *argsDebug == false {

		var err error
		switch b.Format {
		case "tar":
			z := archiver.Tar{}
			err = z.Archive(b.ToBackup, b.TargetArchive)
		default:
			z := archiver.Zip{
				CompressionLevel:       flate.BestCompression,
				MkdirAll:               true,
				SelectiveCompression:   true,
				ContinueOnError:        true,
				OverwriteExisting:      false,
				ImplicitTopLevelFolder: false,
			}
			err = z.Archive(b.ToBackup, b.TargetArchive)
		}
		if err != nil {
			lg.Logf("Error during compression %q -> %q: %s", b.ToBackup, b.TargetArchive, err)
		}
	}
}

func targetArchiveName(b BkpSet) (s string) {
	s = b.OutputFolder
	if b.Subfolder != "" {
		s = syslib.Pj(b.OutputFolder, b.Subfolder)
	}
	s = syslib.Pj(s, b.Timestamp)
	shortname := strings.Replace(
		rx.Find(rxLib.AfterLastSlash, b.ToBackup[0]), ".", "_", -1,
	)
	s = syslib.Pj(s, shortname+"."+b.Format)
	return
}
