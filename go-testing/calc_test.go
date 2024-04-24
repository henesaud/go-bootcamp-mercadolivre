package cal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	result := Sub(10, 5)
	assert.Equal(t, 5, result, "should be equal")
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 5)
	assert.Nil(t, err)
	assert.Equal(t, 2, result)

	_, err = Divide(10, 0)
	assert.NotNil(t, err)
}
