package builder

import "fmt"

// Concrete Human Type

type Human struct {
	name   string
	age    int
	gender string
}

func (h Human) ToString() {
	fmt.Println(h.name, h.age, h.gender)
}
