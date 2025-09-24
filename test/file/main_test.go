package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChangeLog(t *testing.T) {
	f, err := os.CreateTemp("", "TEST_CHANGELOG")
	assert.NoError(t, err)

	defer func() {
		_ = f.Close()
		_ = os.RemoveAll(f.Name())
	}()

	changeLogPath = f.Name()
	data := `
# Changelog
All notable changes to this project will be documented in this file.
`
	_, err = f.WriteString(data)
	assert.NoError(t, err)

	expected := ChangeLogSpec{
		Version:      "stable",
		BuildVersion: "nil",
		ChangeLog: `
# Changelog
All notable changes to this project will be documented in this file.
`,
	}

	actual, err := GetChangeLog()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
