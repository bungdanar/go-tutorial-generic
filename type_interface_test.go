package tutorialgeneric

import (
	"fmt"
	"testing"
)

type Employee interface {
	GetName() string
}

func GetEmployeeName[T Employee](employee T) string {
	return employee.GetName()
}

type Manager struct {
	Name string
}

func (m *Manager) GetName() string {
	return m.Name
}

type Engineer struct {
	Name string
}

func (e *Engineer) GetName() string {
	return e.Name
}

func TestTypeInterface(t *testing.T) {
	manager := Manager{Name: "Mr Dadang"}
	engineer := Engineer{Name: "Mr Ade"}

	fmt.Println(GetEmployeeName[Employee](&manager))
	fmt.Println(GetEmployeeName[Employee](&engineer))
}
