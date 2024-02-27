package option

type OptionContainer interface {
	ToList() []Option
}

type optionContainer struct {
	options []Option
}

// ToList implements DefaultOption.
func (dflt *optionContainer) ToList() []Option {
	return dflt.options
}

func newtOptionContainer(opts []Option) OptionContainer {
	return &optionContainer{options: opts}
}
func newDefaultOptions() OptionContainer {
	return &optionContainer{options: []Option{
		{
			Id:   SALIR,
			Text: "Salir",
		},
		{
			Id:   AVANZAR,
			Text: "Avanzar al siguiente dia",
		},
		{
			Id:   RESUMEN,
			Text: "Mostrar resumen",
		},
	}}
}
