package tutorialgeneric

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, b := range bag {
		fmt.Println(b)
	}
}

func TestBag(t *testing.T) {
	numbers := Bag[int]{1, 2, 3, 4, 5}
	PrintBag[int](numbers)

	names := Bag[string]{"Bung", "Danar", "Sayaskumbang"}
	PrintBag[string](names)
}
