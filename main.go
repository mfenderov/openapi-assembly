package main

import (
	"github.com/mfenderov/openapi-assembly/assembler"
	paramparser "github.com/mfenderov/openapi-assembly/param-parser"
	"github.com/mfenderov/openapi-assembly/writer"
	"log"
	"os"
)

func main() {
	os.Exit(Main())
}

func Main() int {
	inputSpec, outputDestination := paramparser.Parse()
	assemble, err := assembler.Assemble(inputSpec)
	if err != nil {
		log.Fatal(err)
	}
	err = writer.WriteSpec(assemble, outputDestination)
	if err != nil {
		log.Fatal(err)
	}
	return 0
}
