package decorator

// Component Interface
type IPizza interface {
	getPrice() int
}

// Concrete Component
type VeggeMania struct {
}

func (p *VeggeMania) getPrice() int {
	return 15
}
