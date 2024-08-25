package main

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestE2E(t *testing.T) {
	os.Args = []string{os.Args[0], "-i", "test-resources/docs/api.yaml", "-o", "tmp/api.json"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	Main()

	stat, err := os.Stat("tmp/api.json")
	assert.NoError(t, err)
	assert.False(t, stat.IsDir())
	assert.True(t, stat.Size() > 0)
}
