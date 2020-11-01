package main

import (
	"sort"
	"testing"

	"./syslib"
)

func TestReadConfig(t *testing.T) {
	readTomlConfig("../testdata/case1.toml")
}

func TestInitConfig(t *testing.T) {
	basedir := "/tmp/bkpbot/testdata"
	assertInitConfig(
		"../testdata/case1.toml",
		[]string{syslib.Pj(basedir, "fol1")},
		t,
	)

	assertInitConfig(
		"../testdata/case2.toml",
		[]string{
			syslib.Pj(basedir, "fol2/a"),
			syslib.Pj(basedir, "fol2/b"),
			syslib.Pj(basedir, "fol2/c"),
		},
		t,
	)

	assertInitConfig(
		"../testdata/case3.toml",
		[]string{
			syslib.Pj(basedir, "fol2/b"),
			syslib.Pj(basedir, "fol2/c"),
		},
		t,
	)

	assertInitConfig(
		"../testdata/case4.toml",
		[]string{
			syslib.Pj(basedir, "fol2/b"),
			syslib.Pj(basedir, "fol2/c"),
		},
		t,
	)

	assertInitConfig(
		"../testdata/case5.toml",
		[]string{
			syslib.Pj(basedir, "fol2/a"),
		},
		t,
	)
}

func assertInitConfig(configFile string, assertFolders []string, t *testing.T) {
	conf := initConfig(configFile)
	b := conf["1"].ToBackup
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
