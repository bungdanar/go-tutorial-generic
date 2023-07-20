package tutorialgeneric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (data *MyData[T]) GetValue() T {
	return data.Value
}

func (data *MyData[T]) SetValue(value T) {
	data.Value = value
}

func TestGenericInterface(t *testing.T) {
	data := MyData[string]{
		Value: "Bung",
	}

	assert.Equal(t, "Danar", ChangeValue[string](&data, "Danar"))
}
