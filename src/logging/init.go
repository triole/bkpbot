package logging

import (
	"bytes"
	"os"

	"github.com/triole/bkpbot/rx"
)

var (
	rxLib = rx.InitLib()
)

// Self contains all necessary settings
type Self struct {
	Logdir  string
	Logfile FileSet
}

type FileSet struct {
	Name     string
	IOWriter *bytes.Buffer
}

// Init is as the name says the init function, call it in var declaration
func Init(logfile string) (ls Self) {
	ls = Self{
		Logfile: FileSet{
			Name:     logfile,
			IOWriter: bytes.NewBufferString(logfile),
		},
		Logdir: rx.Find(rxLib.UpToLastSlash, logfile),
	}
	// make logdir if it does not exist
	m := os.FileMode(int(0755))
	os.MkdirAll(ls.Logdir, m)
	return
}

// InitDummy can be used in tests
func InitDummy() Self {
	return Self{}
}
