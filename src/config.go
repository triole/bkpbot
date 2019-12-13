package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"olibs/syslib"

	yaml "gopkg.in/yaml.v2"
)

func initConfig(configFile string) (rc RichConfig) {
	c := readConfigYaml(configFile)
	rc = makeRichConfig(c)
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

func makeRichConfig(config Config) (richConfig RichConfig) {
	for _, e := range config {
		root := expandEnv(e.Root)
		folders := detectFolders(root, e.Detect)
		folders = expandEnvMult(folders)
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
