package option

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskWithOptions(question string, opts []Option) Option {
	fmt.Println(question)
	var optionMap map[string]int = make(map[string]int, len(opts))
	for indice, opcion := range opts {
		optionMap[strings.ToLower(opcion.Id)] = indice
	}
	userResponse := Ask(optsToString(opts))
	var indice, intento int
	var ok bool
	for intento = 1; intento <= 3; intento++ {
		indice, ok = optionMap[strings.ToLower(userResponse)]
		if ok {
			break
		}
		fmt.Println("Respuesta incorrecta, intenta de nuevo:")
		userResponse = Ask(optsToString(opts))
	}
	if intento > 3 {
		indice = 0
		fmt.Printf("Respuesta incorrecta\nElegiremos por ti una Option correcta: %s\n", opts[indice].Id)
	}
	return opts[indice]
}

func Ask(question string) string {
	fmt.Println(question)
	lector := bufio.NewReader(os.Stdin)
	respuesta, err := lector.ReadString('\n')

	for err != nil {
		fmt.Println("Hubo un error al leer tu respuesta:", err)
		return Ask(question)
	}

	// Quita el salto de línea del final
	return strings.TrimSpace(respuesta)

}
