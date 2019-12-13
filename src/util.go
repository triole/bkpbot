package main

import (
	"olibs/syslib"
)

func detectFolders(root string, detect bool) (fol []string) {
	fol = []string{root}
	if detect == true {
		fol = syslib.Find(root, ".*", "d", false)
	}
	return
}

func makeFolderSets() {

}
