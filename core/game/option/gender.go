package option

type GenderOption interface {
	ToList() []Option
}

type genderOption struct {
	options []Option
}

// ToList implements GenderOption.
func (gender *genderOption) ToList() []Option {
	return gender.options
}

func NewGenderOption() OptionContainer {
	return &genderOption{options: []Option{
		{
			Id:   "A",
			Text: "Hombre",
		},
		{
			Id:   "B",
			Text: "Mujer",
		},
	}}
}
