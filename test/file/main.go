package main

import "os"

var (
	version       = "stable"
	buildVersion  = "nil"
	changeLogPath = "ChangeLog.md"
)

type ChangeLogSpec struct {
	Version      string
	BuildVersion string
	ChangeLog    string
}

func GetChangeLog() (ChangeLogSpec, error) {
	data, err := os.ReadFile(changeLogPath)
	if err != nil {
		return ChangeLogSpec{}, err
	}

	return ChangeLogSpec{
		Version:      version,
		BuildVersion: buildVersion,
		ChangeLog:    string(data),
	}, nil
}
