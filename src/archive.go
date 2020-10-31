package main

import (
	"compress/flate"
	"olibs/rx"
	"olibs/syslib"

	"github.com/mholt/archiver"
)

func runBetween(cmd string, point string, prefix string) {
	_, err := syslib.RunCmdErrMsg(cmd)
	if err == "" {
		lg.Logf("%s Run %s %q, OK", prefix, point, cmd)
	} else {
		lg.Logf("%s Run %s %q, ERROR: %q", prefix, point, cmd, err)
	}
}

func archive(bs tBkpSet) {
	prefix := "[" + bs.Name + "]"
	lg.Logf("%s Start backup %q", prefix, bs.Name)
	if *argsDebug == false {

		// run before
		for _, cmd := range bs.RunBefore {
			runBetween(cmd, "before", prefix)
		}

		// make output folder although zip does automatically
		op := rx.Find(rxlib.UpToLastSlash, bs.TargetArchive)
		syslib.MkdirAll(op)

		var err error
		switch bs.OutputFormat {
		case "tar":
			z := archiver.Tar{}
			lg.Logf("%s Tar archive %q -> %q ", prefix, bs.ToBackup, bs.TargetArchive)
			err = z.Archive(bs.ToBackup, bs.TargetArchive)
			lg.Logf("%s Archiving finished", prefix)
		default:
			z := archiver.Zip{
				CompressionLevel:       flate.BestCompression,
				MkdirAll:               true,
				SelectiveCompression:   true,
				ContinueOnError:        true,
				OverwriteExisting:      false,
				ImplicitTopLevelFolder: false,
			}
			lg.Logf("%s Zip archive %q -> %q ", prefix, bs.ToBackup, bs.TargetArchive)
			err = z.Archive(bs.ToBackup, bs.TargetArchive)
			lg.Logf("%s Archiving finished", prefix)
		}
		if err != nil {
			lg.Logf("%s Error during compression %q -> %q: %s", prefix, bs.ToBackup, bs.TargetArchive, err)
		}

		// run after
		for _, cmd := range bs.RunAfter {
			runBetween(cmd, "after", prefix)
		}
	}
}
