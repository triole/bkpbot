package main

type Config []Folder

type RichConfig []RichFolder

type Folder struct {
	Root       string   `yaml:"root"`
	Detect     bool     `yaml:"detect"`
	Exclusions []string `yaml:"exclusions"`
}

type RichFolder struct {
	Root       string
	Detect     bool
	Folders    []string
	Exclusions []string
}
