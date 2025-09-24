package main

import (
	_ "embed"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/CHANGELOG.md
var changelog []byte

func TestGetChangeLog_by_embed(t *testing.T) {
	f, err := os.CreateTemp("", "TEST_CHANGELOG")
	assert.NoError(t, err)
	defer func() {
		_ = f.Close()
		_ = os.RemoveAll(f.Name())
	}()

	changeLogPath = f.Name()

	_, err = f.Write(changelog)
	assert.NoError(t, err)

	expected := ChangeLogSpec{
		Version:      "stable",
		BuildVersion: "nil",
		ChangeLog:    string(changelog),
	}

	actual, err := GetChangeLog()
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
