package main

import (
	"olibs/rx"
	"olibs/syslib"
)

func detectFolders(root string, detect bool) (fol []string) {
	fol = []string{root}
	if detect == true {
		fol = syslib.Find(root, ".*", "d", false)
	}
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
