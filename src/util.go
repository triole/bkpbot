package main

import (
	"olibs/rx"
	"olibs/syslib"
	"os"
	"sort"
	"strings"
)

func cleanUp(folder string, keepLast int) {
	folders := detectFolders(folder)
	lg.Logf("Clean up. Keep only last %v archives of %v", *argsKeepLast, len(folders))
	sort.Strings(folders)
	for _, folder := range folders[:len(folders)-keepLast] {
		lg.Logf("Delete folder %q", folder)
		if *argsDebug == false {
			os.RemoveAll(folder)
		}
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

func detectFolders(root string) (fol []string) {
	fol = []string{root}
	fol = syslib.Find(root, ".*", "d", false)
	sort.Strings(fol)
	return
}

func removeExclusions(allFolders []string, exclusions []string) (folders []string) {
	for _, f := range allFolders {
		keep := true
		for _, e := range exclusions {
			if rx.Match(e, f) {
				keep = false
			}
		}
		if keep == true {
			folders = append(folders, f)
		}
	}
	return
}

func expandEnv(folder string, configFileDir string) (s string) {
	s = os.ExpandEnv(folder)
	s = strings.Replace(s, "<CURDIR>", env.Curdir, -1)
	s = strings.Replace(s, "<SCRIPTDIR>", env.Scriptdir, -1)
	s = strings.Replace(s, "<CONFIGDIR>", configFileDir, -1)
	s = syslib.Pabs(s)
	return
}
