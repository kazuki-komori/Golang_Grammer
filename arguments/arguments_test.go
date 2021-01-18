package arguments

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIfArgument(t *testing.T) {
	res := IfArgument(1)
	assert.Equal(t, false, res)
}
