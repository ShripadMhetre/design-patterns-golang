package builder

// Builder Type

type HumanBuilder struct {
	name   string
	age    int
	gender string
}

func (builder *HumanBuilder) SetName(name string) *HumanBuilder {
	builder.name = name
	return builder
}

// Builder Methods

func (builder *HumanBuilder) SetAge(age int) *HumanBuilder {
	builder.age = age
	return builder
}

func (builder *HumanBuilder) SetGender(gender string) *HumanBuilder {
	builder.gender = gender
	return builder
}

// final "Build" method to build to Human object

func (builder *HumanBuilder) Build() Human {
	return Human{
		name:   builder.name,
		age:    builder.age,
		gender: builder.gender,
	}
}
