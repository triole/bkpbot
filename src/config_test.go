package main

import (
	"olibs/syslib"
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
}

func assertInitConfig(configFile string, assertFolders []string, t *testing.T) {
	conf := initConfig(configFile)
	for idx, fol := range conf[0].Folders {
		a := assertFolders[idx]
		if a != fol {
			t.Errorf("Assertion failed: %q != %q", a, fol)
		}
	}
}
