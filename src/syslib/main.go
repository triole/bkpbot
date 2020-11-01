package syslib

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"../rx"
)

// Find detects files in folder, even recursively
func Find(basedir string, rxFilter string, dirOrFile string, recursively bool) []string {
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

// MkdirAll is like "mkdir -p"
func MkdirAll(foldername string) {
	os.MkdirAll(foldername, os.ModePerm)
}

// Pabs makes sure a path is absolute
func Pabs(pathstring string) string {
	r, err := filepath.Abs(pathstring)
	if err != nil {
		fmt.Printf("Unable to make absolute path. %s\n", err)
		X(1)
	}
	return r
}

// Pj is kind of like Pytons path.join
func Pj(p1 string, p2 string) string {
	return path.Join(p1, p2)
}

// ReadFileToString reads content of a file into a string
func ReadFileToString(filename string) string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("An error occured: %q", err)
		X(1)
	}
	return string(f)
}

// X exits the program returning an exit code
func X(exitCode int) {
	os.Exit(exitCode)
}
