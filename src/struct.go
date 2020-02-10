package main

type Config []Folder

type RichConfig []RichFolder

type Folder struct {
	ToBackup   []string `yaml:"to_backup"`
	Detect     bool     `yaml:"detect"`
	Exclusions []string `yaml:"exclusions"`
	Output     Output   `yaml:"output"`
}

type Output struct {
	Name   string `yaml:"name"`
	Folder string `yaml:"folder"`
	Format string `yaml:"format"`
}

type RichFolder struct {
	ToBackup []string
	Output   Output
}

type BkpSet struct {
	ToBackup      []string
	Output        Output
	Subfolder     string
	Timestamp     string
	TargetArchive string
	KeepLast      int
}
