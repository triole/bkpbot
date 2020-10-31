package main

import (
	"compress/flate"
	"olibs/rx"
	"olibs/syslib"

	"github.com/mholt/archiver"
)

func archive(bs tBkpSet) {
	lg.Logf("Run backup %q: %q -> %q", bs.Name, bs.ToBackup, bs.TargetArchive)
	if *argsDebug == false {

		// make output folder although zip does automatically
		op := rx.Find(rxlib.UpToLastSlash, bs.TargetArchive)
		syslib.MkdirAll(op)

		var err error
		switch bs.OutputFormat {
		case "tar":
			z := archiver.Tar{}
			err = z.Archive(bs.ToBackup, bs.TargetArchive)
		default:
			z := archiver.Zip{
				CompressionLevel:       flate.BestCompression,
				MkdirAll:               true,
				SelectiveCompression:   true,
				ContinueOnError:        true,
				OverwriteExisting:      false,
				ImplicitTopLevelFolder: false,
			}
			err = z.Archive(bs.ToBackup, bs.TargetArchive)
		}
		if err != nil {
			lg.Logf("Error during compression %q -> %q: %s", bs.ToBackup, bs.TargetArchive, err)
		}
	}
}
