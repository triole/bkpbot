package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"olibs/rx"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

func find(basedir string, rxFilter string, dirOrFile string, recursively bool) []string {
	filelist := []string{}
	// index files
	if recursively == true {
		filelist = listFolderContentRecursively(basedir)
	} else {
		filelist = listFolderContentSimple(basedir)
	}
	// loop over array to apply regex filter
	filelist = rx.FilterArrayByRegex(filelist, rxFilter)
	// gather file infos and filter by type
	newlist := []string{}
	for _, item := range filelist {
		fileinfo, err := os.Stat(item)
		if err != nil {
			log.Fatal(err)
		}
		if fileinfo.IsDir() == true && dirOrFile == "d" {
			newlist = append(newlist, item)
		} else if fileinfo.IsDir() == false && dirOrFile == "f" {
			newlist = append(newlist, item)
		}
	}
	sort.Strings(newlist)
	return newlist
}

func listFolderContentRecursively(basedir string) []string {
	filelist := []string{}
	err := filepath.Walk(basedir, func(path string, f os.FileInfo, err error) error {
		filelist = append(filelist, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return filelist
}

func listFolderContentSimple(basedir string) []string {
	if strings.HasSuffix(basedir, "/") == false {
		basedir += "/"
	}
	filelist, err := filepath.Glob(basedir + "*")
	if err != nil {
		log.Fatal(err)
	}
	return filelist
}

func mkdirAll(foldername string) {
	os.MkdirAll(foldername, os.ModePerm)
}

func pabs(pathstring string) string {
	r, err := filepath.Abs(pathstring)
	if err != nil {
		fmt.Printf("Unable to make absolute path. %s\n", err)
		x(1)
	}
	return r
}

func pj(p1 string, p2 string) string {
	return path.Join(p1, p2)
}

func readFileToString(filename string) string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("An error occured: %q", err)
		x(1)
	}
	return string(f)
}

func x(exitCode int) {
	os.Exit(exitCode)
}
