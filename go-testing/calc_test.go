package cal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	result := Sub(10, 5)
	assert.Equal(t, 5, result, "should be equal")
}
