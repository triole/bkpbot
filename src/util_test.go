package main

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/triole/bkpbot/syslib"
)

func TestCleanup(t *testing.T) {
	basedir := "/tmp/bkpbot_test/unittest"
	syslib.MkdirAll(syslib.Pj(basedir, "hello"))
	syslib.MkdirAll(syslib.Pj(basedir, "world"))
	for i := 1; i <= 10; i++ {
		syslib.MkdirAll(syslib.Pj(basedir, randomBkpFolder()))
	}
	cleanUp(basedir, 0)
	cleanUp(basedir, 5)
	if len(detectFolders(basedir, ".*")) != 7 {
		t.Errorf("CleanUp assertion failed.")
	}
}

func TestStringInSlice(t *testing.T) {
	if stringInSlice("hello", []string{"hello", "world"}) == false {
		t.Errorf("StringInSlice assertion error")
	}
	if stringInSlice("whatever", []string{"hello", "world"}) == true {
		t.Errorf("StringInSlice assertion error")
	}
}

func TestShouldExclude(t *testing.T) {
	assertTestShouldExclude("/hello/world", []string{".*/world$"}, true, t)
	assertTestShouldExclude("/hello/world", []string{"whatever", "string"}, false, t)
}

func assertTestShouldExclude(folder string, exclusions []string, exp bool, t *testing.T) {
	b := shouldExclude(folder, exclusions)
	if exp != b {
		t.Errorf("ShouldExclude assertion failed: %q, %q, %v", folder, exclusions, b)
	}
}

// test util
func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func randomBkpFolder() string {
	return strconv.Itoa(randomInt(10000000, 99999999)) +
		"_" + strconv.Itoa(randomInt(100000, 999999))
}
