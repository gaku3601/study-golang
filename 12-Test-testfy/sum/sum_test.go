package sum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert := assert.New(t)
	// Equal
	assert.Equal(Sum(1, 2), 3, "Calc/Sumの加算テスト") // Ok
	//assert.Equal(Sum(1, 2), 4, "Calc/Sumの加算テスト") // false
}
