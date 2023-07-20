package tutorialgeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func GetFirst[T []E, E any](data T) E {
	return data[0]
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin(100, 200))
	assert.Equal(t, int64(100), FindMin(int64(100), int64(200)))
	assert.Equal(t, float64(100), FindMin(float64(100), float64(200)))
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Bung", "Danar",
	}

	first := GetFirst[[]string, string](names)
	assert.Equal(t, "Bung", first)
}
