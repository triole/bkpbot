package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

func initConfig(configFile string) (rc tRichConfig) {
	configFile = pabs(configFile)
	c := readTomlConfig(configFile)
	rc = makeRichConfig(c)
	return
}

func readTomlConfig(filename string) (c tConfig) {
	content := readFileToString(filename)
	if _, err := toml.Decode(string(content), &c); err != nil {
		lg.Logf("Exception reading config %q. %s\n", filename, err)
		x(1)
	}
	return
}

func makeRichConfig(config tConfig) (richConfig tRichConfig) {
	richConfig = make(tRichConfig)
	for name, bs := range config.Jobs {

		var toBackup []string
		for _, f := range bs.ToBackup {
			folder := expandVars(f, config.Vars)

			// check if folder exists, only act if it does
			if _, err := os.Stat(folder); !os.IsNotExist(err) {
				if shouldExclude(folder, bs.Exclusions) == false {
					if bs.Detect == true {
						folders := detectFolders(folder, ".*")
						for _, folder := range folders {
							if shouldExclude(folder, bs.Exclusions) == false {
								toBackup = append(toBackup, folder)
							}
						}
					} else {
						toBackup = append(toBackup, folder)
					}
				}
			} else {
				lg.Logf("Error: File or folder does not exist. Not processing %q", folder)
			}
		}

		format := bs.OutputFormat
		if format == "" {
			format = "zip"
		}

		rc := tRichFolder{
			ToBackup:     toBackup,
			OutputName:   bs.OutputName,
			OutputFolder: expandVars(bs.OutputFolder, config.Vars),
			OutputFormat: format,
			RunBefore:    bs.RunBefore,
			RunAfter:     bs.RunAfter,
		}
		richConfig[name] = rc
	}
	return
}

func saveRichConfig(rc tRichConfig, targetfile string) {
	JSONData, _ := json.MarshalIndent(rc, "", "\t")
	_ = ioutil.WriteFile(targetfile, JSONData, 0644)
}
