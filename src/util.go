package main

import (
	"olibs/rx"
	"olibs/syslib"
	"os"
	"strings"
)

func detectFolders(root string) (fol []string) {
	fol = []string{root}
	fol = syslib.Find(root, ".*", "d", false)
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
