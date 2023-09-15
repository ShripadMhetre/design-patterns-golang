package factory

import "fmt"

// OS interfaces
type OS interface {
	specs()
}

// Concrete OS Implementations - Windows, MacOS & Linux

type Windows struct{}

func (w Windows) specs() {
	fmt.Println("This is Windows OS...")
}

type MacOS struct{}

func (i MacOS) specs() {
	fmt.Println("This is Mac OS...")
}

type Linux struct{}

func (l Linux) specs() {
	fmt.Println("This is Linux...")
}

// OS Factory
func getOS(name string) (OS, error) {
	if name == "windows" {
		return Windows{}, nil
	}

	if name == "linux" {
		return Linux{}, nil
	}

	if name == "mac" {
		return MacOS{}, nil
	}

	return nil, fmt.Errorf("wrong os type")
}
