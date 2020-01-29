package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"olibs/rx"
	"olibs/syslib"

	yaml "gopkg.in/yaml.v2"
)

func initConfig(configFile string) (rc RichConfig) {
	configFile = syslib.Pabs(configFile)
	configDir := rx.Find(rxLib.UpToLastSlash, configFile)
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

func makeRichConfig(config Config, configFileDir string) (richConfig RichConfig) {
	for _, e := range config {

		var toBackup []string
		for _, f := range e.ToBackup {
			t := expandEnv(f, configFileDir)
			if e.Detect == true {
				toBackup = append(toBackup, detectFolders(t)...)
			} else {
				toBackup = append(toBackup, t)
			}
		}
		if len(e.Exclusions) > 0 {
			toBackup = removeExclusions(toBackup, e.Exclusions)
		}
		r := RichFolder{
			ToBackup:     toBackup,
			OutputFolder: expandEnv(e.OutputFolder, configFileDir),
			Format:       e.Format,
		}
		richConfig = append(richConfig, r)
	}
	return
}

func saveRichConfig(rc RichConfig, targetfile string) {
	JSONData, _ := json.MarshalIndent(rc, "", "\t")
	_ = ioutil.WriteFile(targetfile, JSONData, 0644)
}
