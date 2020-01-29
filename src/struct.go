package main

type Config []Folder

type RichConfig []RichFolder

type Folder struct {
	ToBackup     []string `yaml:"to_backup"`
	Detect       bool     `yaml:"detect"`
	Exclusions   []string `yaml:"exclusions"`
	OutputFolder string   `yaml:"output_folder"`
	Format       string   `yaml:"format"`
	KeepLast     int      `yaml:"keep_last"`
}

type RichFolder struct {
	ToBackup     []string
	OutputFolder string
	Format       string
	KeepLast     int
}

type BkpSet struct {
	ToBackup      []string
	OutputFolder  string
	Subfolder     string
	Timestamp     string
	TargetArchive string
	Format        string
	KeepLast      int
}
