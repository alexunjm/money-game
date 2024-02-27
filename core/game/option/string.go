package option

import (
	"fmt"
	"strings"
)

func optsToString(opts []Option) string {
	var builder strings.Builder

	for _, opt := range opts {
		builder.WriteString(fmt.Sprintf(" %v. %v\n", opt.Id, opt.Text))
	}

	return builder.String() // Obtener el string resultante
}
