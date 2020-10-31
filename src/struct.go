package main

type tConfig struct {
	Vars tVars
	Jobs tJobs
}

type tVars map[string]string
type tJobs map[string]tFolder

type tRichConfig map[string]tRichFolder

type tFolder struct {
	ToBackup     []string `toml:"to_backup"`
	Detect       bool     `toml:"detect"`
	Exclusions   []string `toml:"exclusions"`
	OutputName   string   `toml:"output_name"`
	OutputFolder string   `toml:"output_folder"`
	OutputFormat string   `toml:"output_format"`
	RunBefore    []string `toml:"run_before"`
	RunAfter     []string `toml:"run_after"`
}

type tRichFolder struct {
	ToBackup     []string
	OutputName   string
	OutputFolder string
	OutputFormat string
	RunBefore    []string
	RunAfter     []string
}

type tBkpSet struct {
	ToBackup      []string
	OutputName    string
	OutputFolder  string
	OutputFormat  string
	Subfolder     string
	Timestamp     string
	TargetArchive string
	KeepLast      int
	Name          string
	RunBefore     []string
	RunAfter      []string
}
