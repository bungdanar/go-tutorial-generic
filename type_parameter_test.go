package tutorialgeneric

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	result := Length[string]("bung")
	assert.Equal(t, "bung", result)

	resultNumber := Length[int](100)
	assert.Equal(t, 100, resultNumber)
}
