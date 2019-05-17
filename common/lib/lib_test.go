package lib

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestConvertStringToInt(t *testing.T) {
	assert := assert.New(t)

	ans, err := ConvertStringToInt("4")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(4, ans)
	assert.True(true)
}
