package main

import (
	"encoding/json"
	"io/ioutil"
	"olibs/rx"
	"olibs/syslib"
	"os"

	"github.com/BurntSushi/toml"
)

func initConfig(configFile string) (rc tRichConfig) {
	configFile = syslib.Pabs(configFile)
	configDir := rx.Find(rxlib.UpToLastSlash, configFile)
	c := readTomlConfig(configFile)
	rc = makeRichConfig(c, configDir)
	return
}

func readTomlConfig(filename string) (c tConfig) {
	content := syslib.ReadFileToString(filename)
	if _, err := toml.Decode(string(content), &c); err != nil {
		lg.Logf("Exception reading config %q. %s\n", filename, err)
		syslib.X(1)
	}
	return
}

func makeRichConfig(config tConfig, configFileDir string) (richConfig tRichConfig) {
	richConfig = make(tRichConfig)
	for name, bs := range config {

		var toBackup []string
		for _, f := range bs.ToBackup {
			folder := expandEnv(f, configFileDir)

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
				lg.Logf("Folder or file does not exist. Not processing %q", folder)
			}
		}

		format := bs.OutputFormat
		if format == "" {
			format = "zip"
		}

		rc := tRichFolder{
			ToBackup:     toBackup,
			OutputName:   bs.OutputName,
			OutputFolder: expandEnv(bs.OutputFolder, configFileDir),
			OutputFormat: format,
		}
		richConfig[name] = rc
	}
	return
}

func saveRichConfig(rc tRichConfig, targetfile string) {
	JSONData, _ := json.MarshalIndent(rc, "", "\t")
	_ = ioutil.WriteFile(targetfile, JSONData, 0644)
}
