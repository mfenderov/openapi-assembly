package param_parser

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParserWithCustomValues(t *testing.T) {
	os.Args = []string{os.Args[0], "-i", "asd/fgh.txt", "-o", "zxc/vbn.exe"}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	inputSpec, outputDestination := Parse()
	assert.Equal(t, "asd/fgh.txt", inputSpec)
	assert.Equal(t, "zxc/vbn.exe", outputDestination)
}

func TestParserWithDefaultValues(t *testing.T) {
	os.Args = []string{os.Args[0]}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	inputSpec, outputDestination := Parse()
	assert.Equal(t, "docs/api.yaml", inputSpec)
	assert.Equal(t, "api.json", outputDestination)
}
