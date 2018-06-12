package plugins

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHashOfFile(t *testing.T) {
	tests := []struct {
		input  string
		output string
		throws bool
	}{
		{
			input:  "testdata/hash/testfile.md",
			output: "6ce1f29e6b4c27210b8f6ab6c19655e582385ee6493524d670ecbed3c66868ac",
			throws: false,
		},
		{
			input:  "testdata/hash/notafile.md",
			output: "",
			throws: true,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput, err := getHashOfFile(
			test.input,
		)
		if test.throws {
			assert.NotNil(err)
		} else {
			assert.Nil(err)
			assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d:", i))
		}
	}
}
