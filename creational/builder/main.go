package builder

import "fmt"

func Main() {
	humanBuilder := &HumanBuilder{}

	human := humanBuilder.SetName("John Doe").SetAge(22).SetGender("Male").Build()

	fmt.Println(human)
}
