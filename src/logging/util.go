package logging

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/triole/bkpbot/rx"
)

// write a line into logfile
func (s Self) writeLogLine(line string) {
	// clean up string
	timestamp := timestamp()
	mclean := rx.Sub(rxLib.AnsiSequence, line, "")
	mclean = rx.Sub(rxLib.ControlCharacters, mclean, "")
	logmsg := timestamp + " " + mclean + "\n"
	// write it into the file
	addStringToFile(s.Logfile.Name, logmsg)
}

func timestamp() (t string) {
	n := time.Now()
	t = n.Format("2006-01-02 15:04:05.000")
	return
}

func addStringToFile(filename string, stringToWrite string) {
	rxLib := rx.InitLib()
	// create folder if it does not exist
	filefolder := rx.Find(rxLib.UpToLastSlash, filename)
	createFolderIfNotExists(filefolder)

	// create your file with desired read/write permissions
	fo, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error. Unable to write log: %q\n", err)
		os.Exit(1)
	}
	defer fo.Close()

	// write into it
	_, err = io.Copy(fo, strings.NewReader(stringToWrite))
	if err != nil {
		fmt.Printf("Error. Unable to write log: %q\n", err)
		os.Exit(1)
	}
}

func createFolderIfNotExists(foldername string) {
	if _, err := os.Stat(foldername); os.IsNotExist(err) {
		os.Mkdir(foldername, os.ModePerm)
	}
}
