package option

type OptionBuilder interface {
	ResetDefaultOptions()
	WithOptions([]Option)
	Build() OptionContainer
}
type optionBuilder struct {
	options OptionContainer
}

// WithOptions implements OptionBuilder.
func (dfltOptBldr *optionBuilder) WithOptions(opts []Option) {
	dfltOptBldr.options = newtOptionContainer(opts)
}

// Build implements OptionBuilder.
func (dfltOptBldr *optionBuilder) Build() OptionContainer {
	return dfltOptBldr.options
}

// init implements OptionBuilder.
func (b *optionBuilder) ResetDefaultOptions() {
	b.options = newDefaultOptions()
}

func NewOptionBuilder() OptionBuilder {
	x := &optionBuilder{}
	x.ResetDefaultOptions()
	return x
}
