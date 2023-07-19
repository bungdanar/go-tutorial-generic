package tutorialgeneric

import (
	"fmt"
	"testing"
)

func MultipleParams[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParams(t *testing.T) {
	MultipleParams[string, int]("Bung", 100)
}
