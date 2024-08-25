package param_parser

import (
	"flag"
)

func Parse() (inputSpec string, outputDestination string) {
	if flag.Lookup("i") == nil {
		flag.StringVar(&inputSpec, "i", "docs/api.yaml", "input specification")
	}
	if flag.Lookup("o") == nil {
		flag.StringVar(&outputDestination, "o", "api.json", "output specification")
	}
	flag.Parse()

	inputSpec = flag.Lookup("i").Value.(flag.Getter).Get().(string)
	outputDestination = flag.Lookup("o").Value.(flag.Getter).Get().(string)

	return inputSpec, outputDestination
}
