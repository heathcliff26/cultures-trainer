package app

import (
	"os"
	"runtime"
	"testing"

	"fyne.io/fyne/v2/test"
	"github.com/go-openapi/testify/v2/assert"
)

func TestGetVersion(t *testing.T) {
	oldGitCommit := gitCommit
	defer func() { gitCommit = oldGitCommit }()

	gitCommit = "1234567890abcdef"

	a := test.NewApp()
	v := getVersion(a)

	assert := assert.New(t)

	if a.Metadata().Name != "" {
		assert.Contains(a.Metadata().Name, v.Name)
	} else {
		assert.Contains(os.Args[0], v.Name)
	}
	assert.Equal("v"+a.Metadata().Version, v.Version)
	assert.Equal("1234567", v.Commit, "commit hash should be truncated")
	assert.Equal(runtime.Version(), v.Go)
}

func TestInitGitCommit(t *testing.T) {
	oldGitCommit := gitCommit
	defer func() { gitCommit = oldGitCommit }()
	assert := assert.New(t)

	gitCommit = "1234567890abcdef"
	initGitCommit()
	assert.Equal("1234567890abcdef", gitCommit, "gitCommit should not be changed")

	gitCommit = "$Format:%H$"
	initGitCommit()
	assert.NotEqual("$Format:%H$", gitCommit, "gitCommit should be changed")
}

func TestMinimumLengthString(t *testing.T) {
	assert := assert.New(t)

	str := "Hello"
	minLen := 10
	result := minimumLengthString(str, minLen)
	assert.Equal(minLen, len(result), "resulting string should have the minimum length")
	assert.Equal("Hello     ", result, "string should be padded with spaces")

	str = "Hello, World!"
	minLen = 5
	result = minimumLengthString(str, minLen)
	assert.Equal(len(str), len(result), "resulting string should retain original length")
	assert.Equal(str, result, "string should remain unchanged")
}
