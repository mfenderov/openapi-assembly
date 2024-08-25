package assembler

import (
	"github.com/bdpiprava/scalar-go/loader"
	"github.com/bdpiprava/scalar-go/model"
	"path/filepath"
)

func Assemble(inputSpec string) (*model.Spec, error) {
	specFile := filepath.Base(inputSpec)
	specDir := filepath.Dir(inputSpec)

	spec, err := loader.LoadWithName(specDir, specFile)
	if err != nil {
		return nil, err
	}
	return spec, nil

}
