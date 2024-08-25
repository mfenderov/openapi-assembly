package writer

import (
	"encoding/json"
	"github.com/bdpiprava/scalar-go/model"
	"os"
	"path/filepath"
)

func WriteSpec(spec *model.Spec, outputDestination string) error {
	marshal, err := json.Marshal(spec)
	if err != nil {
		return err
	}
	specDir := filepath.Dir(outputDestination)
	err = os.MkdirAll(specDir, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(outputDestination, marshal, 0644)
	if err != nil {
		return err
	}
	return nil
}
