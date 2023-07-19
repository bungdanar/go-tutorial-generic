package tutorialgeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1 T, value2 T) bool {
	return value1 == value2
}

func TestIsSame(t *testing.T) {
	assert.Equal(t, true, IsSame[string]("bung", "bung"))
	assert.Equal(t, false, IsSame[int](100, 110))
}
