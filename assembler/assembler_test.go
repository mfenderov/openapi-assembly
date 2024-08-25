package assembler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssemble(t *testing.T) {
	assembledSpec, err := Assemble("../test-resources/docs/api.yaml")
	assert.NoError(t, err)

	assert.NotNil(t, assembledSpec)
}
