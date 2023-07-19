package tutorialgeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (data *Data[T]) SayHello(name string) string {
	return "Hello " + name
}

func (data *Data[T]) ChangeFirst(value T) T {
	data.First = value
	return data.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Bung",
		Second: "Danar",
	}

	assert.Equal(t, "Kang", data.ChangeFirst("Kang"))
	assert.Equal(t, "Hello Kong", data.SayHello("Kong"))
}
