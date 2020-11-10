package main

import (
	"os"
	"sort"
	"strings"
	"time"

	"github.com/triole/bkpbot/rx"
	"github.com/triole/bkpbot/syslib"
)

func cleanUp(folder string, keepLast int) {
	folders := detectFolders(folder, `.*\/[0-9]{8}_[0-9]{6}$`)
	lg.Logf("Clean up. Keep last %v of %v archives", *argsKeepLast, len(folders))
	sort.Strings(folders)
	if len(folders) > keepLast && keepLast > 0 {
		for _, folder := range folders[:len(folders)-keepLast] {
			lg.Logf("Delete folder %q", folder)
			if *argsDebug == false {
				os.RemoveAll(folder)
			}
		}
	} else {
		lg.Logf("Keep all backups")
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func detectFolders(root, rx string) (fol []string) {
	fol = []string{root}
	fol = syslib.Find(root, rx, "d", false)
	sort.Strings(fol)
	return
}

func shouldExclude(folder string, exclusions []string) (exl bool) {
	exl = false
	for _, e := range exclusions {
		if rx.Match(e, folder) == true {
			exl = true
			break
		}
	}
	return
}

func expandVars(folder string, vars tVars) (s string) {
	s = folder
	for key, val := range vars {
		s = strings.Replace(s, "{{"+strings.ToUpper(key)+"}}", val, -1)
	}
	return
}

func getTimestamp() (t string) {
	n := time.Now()
	t = n.Format("20060102_150405")
	return
}
