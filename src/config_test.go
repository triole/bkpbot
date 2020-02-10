package main

import (
	"olibs/syslib"
	"sort"
	"testing"
)

func TestInitConfig(t *testing.T) {
	basedir := syslib.Pj(env.Curdir, "/../testdata/")
	assertInitConfig(
		"../testdata/testcase1.yaml",
		[]string{syslib.Pj(basedir, "fol1")},
		t,
	)

	assertInitConfig(
		"../testdata/testcase2.yaml",
		[]string{
			syslib.Pj(basedir, "fol2/a"),
			syslib.Pj(basedir, "fol2/b"),
			syslib.Pj(basedir, "fol2/c"),
		},
		t,
	)

	assertInitConfig(
		"../testdata/testcase3.yaml",
		[]string{
			syslib.Pj(basedir, "fol2/b"),
			syslib.Pj(basedir, "fol2/c"),
		},
		t,
	)

	assertInitConfig(
		"../testdata/testcase4.yaml",
		[]string{
			syslib.Pj(basedir, "fol2/b"),
			syslib.Pj(basedir, "fol2/c"),
		},
		t,
	)

	assertInitConfig(
		"../testdata/testcase5.yaml",
		[]string{
			syslib.Pj(basedir, "fol2/a"),
		},
		t,
	)
}

func assertInitConfig(configFile string, assertFolders []string, t *testing.T) {
	conf := initConfig(configFile)
	b := conf[0].ToBackup
	sort.Strings(assertFolders)
	sort.Strings(b)
	if len(assertFolders) != len(b) {
		printFail(assertFolders, b, configFile, t)
	} else {
		for idx, fol := range b {
			a := assertFolders[idx]
			if a != fol {
				printFail(assertFolders, b, configFile, t)
			}
		}
	}
}

func printFail(a, b []string, c string, t *testing.T) {
	t.Errorf("Assertion failed: %q != %q, Config file: %q", a, b, c)
}
