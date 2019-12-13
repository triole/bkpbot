package main

type Config []Folder

type RichConfig []RichFolder

type Folder struct {
	Root         string   `yaml:"root"`
	Detect       bool     `yaml:"detect"`
	Exclusions   []string `yaml:"exclusions"`
	OutputFolder string   `yaml:"output_folder"`
	Format       string   `yaml:"format"`
}

type RichFolder struct {
	Root         string
	Detect       bool
	Folders      []string
	Exclusions   []string
	OutputFolder string
	Format       string
}

type BkpSet struct {
	Folder        string
	OutputFolder  string
	Subfolder     string
	Timestamp     string
	TargetArchive string
	Format        string
}
