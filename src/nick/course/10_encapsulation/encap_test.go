package encapsulation_test

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCncapInit(t *testing.T) {
	e := Employee{"001", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 20}
	e2 := new(Employee)
	e2.Id = "002"
	e2.Age = 22
	e2.Name = "Rose"

	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)   // %T means type
	t.Logf("e2 is %T", e2) // pointer
}

// func (e Employee) formatStr() string {
// 	return fmt.Sprintf("Id: %s, Name: %s, Age: %d", e.Id, e.Name, e.Age)
// }

// (e *Employee) this line means bind function to type pointer

func (e *Employee) formatStr() string {
	// diff with above method: use pointer to operate the same obj
	return fmt.Sprintf("Id: %s; Name: %s; Age: %d", e.Id, e.Name, e.Age)
}

func TestEmployeeObj(t *testing.T) {
	e := Employee{"001", "Bob", 20}
	t.Log(e.formatStr()) // Id: 001, Name: Bob, Age: 20

}
