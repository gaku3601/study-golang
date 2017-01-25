package sub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Sub(5, 4), 1, "引き算のテスト")
}
