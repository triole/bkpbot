package main

import (
	"compress/flate"
	"olibs/rx"
	"olibs/syslib"

	"github.com/mholt/archiver/v3"
)

func archive(b BkpSet) {
	lg.Logf("Archive folder %q -> %q", b.Folder, b.TargetArchive)
	if *argsDebug == false {

		var err error
		switch b.Format {
		case "txz":
			z := archiver.TarXz{}
			err = z.Archive([]string{b.Folder}, b.TargetArchive)
		default:
			z := archiver.Zip{
				CompressionLevel:       flate.BestCompression,
				MkdirAll:               true,
				SelectiveCompression:   true,
				ContinueOnError:        true,
				OverwriteExisting:      false,
				ImplicitTopLevelFolder: false,
			}
			err = z.Archive([]string{b.Folder}, b.TargetArchive)
		}
		if err != nil {
			lg.Logf("Error during compression %q -> %q: %s", b.Folder, b.TargetArchive, err)
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
