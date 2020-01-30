package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"olibs/rx"
	"olibs/syslib"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func initConfig(configFile string) (rc RichConfig) {
	configFile = syslib.Pabs(configFile)
	configDir := rx.Find(rxlib.UpToLastSlash, configFile)
	c := readConfigYaml(configFile)
	rc = makeRichConfig(c, configDir)
	return
}

func readConfigYaml(filename string) (c Config) {
	lg.Logf("Read config %q", filename)
	yamlFile, err := ioutil.ReadFile(syslib.Pabs(filename))
	if err != nil {
		log.Fatalf("File read %q", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal %q", err)
	}
	return
}

// func appendConsideringExclusion()
func makeRichConfig(config Config, configFileDir string) (richConfig RichConfig) {
	for _, bs := range config {

		var toBackup [][]string
		for _, f := range bs.ToBackup {
			folder := expandEnv(f, configFileDir)

			// check if folder exists, only act if it does
			if _, err := os.Stat(folder); !os.IsNotExist(err) {
				if shouldExclude(folder, bs.Exclusions) == false {
					if bs.Detect == true {
						folders := detectFolders(folder, ".*")
						var foldersArr []string
						for _, folder := range folders {
							if shouldExclude(folder, bs.Exclusions) == false {
								foldersArr = append(foldersArr, folder)
							}
						}
						toBackup = append(toBackup, foldersArr)
					} else {
						toBackup = append(toBackup, []string{folder})
					}
				}
			} else {
				lg.Logf("Folder or file does not exist. Not processing %q", folder)
			}
		}

		format := bs.Output.Format
		if format == "" {
			format = "zip"
		}

		rc := RichFolder{
			ToBackup: toBackup,
			Output: Output{
				Name:   bs.Output.Name,
				Folder: expandEnv(bs.Output.Folder, configFileDir),
				Format: format,
			},
		}
		richConfig = append(richConfig, rc)
	}
	return
}

func saveRichConfig(rc RichConfig, targetfile string) {
	JSONData, _ := json.MarshalIndent(rc, "", "\t")
	_ = ioutil.WriteFile(targetfile, JSONData, 0644)
}
