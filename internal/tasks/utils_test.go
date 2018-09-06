package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCliSliceMap(t *testing.T) {
	inSlice := []string{
		"FooFlag=Foo",
		"BarFlag=Bar",
		"BazFlag=Baz",
	}

	outMap := cliSliceMap(inSlice)

	assert.Len(t, outMap, 3)
	assert.Equal(t, "Foo", outMap["FooFlag"])
	assert.Equal(t, "Bar", outMap["BarFlag"])
	assert.Equal(t, "Baz", outMap["BazFlag"])
}
