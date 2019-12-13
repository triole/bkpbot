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
		root := expandEnv(e.Root, configFileDir)
		folders := detectFolders(root, e.Detect)
		folders = expandEnvMult(folders, configFileDir)
		folders = removeExclusions(folders, e.Exclusions)
		r := RichFolder{
			Root:         root,
			Detect:       e.Detect,
			Exclusions:   e.Exclusions,
			Folders:      folders,
			OutputFolder: e.OutputFolder,
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
