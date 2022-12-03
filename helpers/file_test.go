package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInputContent(t *testing.T) {
	scanner, file, err := GetInput("./testinputs/input.txt")
	defer file.Close()
	assert.Nil(t, err)

	var lines = []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	assert.Equal(t, []string{"5324", "1197", "", "thhPLzWzhzwPtLRL WrQlpPvvClcVcCppSvpl", "ggrLwFgWCBwbMWBbFwLMgNBZdmZHclJPllnJlNRPmS"}, lines)
}

func TestGetInputUnknownFile(t *testing.T) {
	_, _, err := GetInput("./unknownpath/unknown_file.txt")
	assert.NotNil(t, err)
}
