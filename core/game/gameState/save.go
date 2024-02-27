package gameState

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"

	"github.com/alexunjm/money-game/core/game/option"
)

type GameToSave struct {
	Nombre, Sexo          string
	Dia                   int
	Saldo, GananciaDiaria float32

	ListaOpciones []option.Option
}

func (gs *GameToSave) save() {

	// Crear un buffer para mantener el flujo de bytes
	var buffer bytes.Buffer

	// Crear un encoder que escriba al buffer
	encoder := gob.NewEncoder(&buffer)

	// Serializar la estructura
	err := encoder.Encode(gs)
	if err != nil {
		fmt.Println("Error al serializar: ", err)
		panic(err)
	}

	// Ahora puedes guardar el buffer en un archivo
	err = os.WriteFile(SAVE_FILE_NAME, buffer.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error al escribir en el archivo: ", err)
		panic(err)
	}
}

func Save(gs *gameSate) {
	x := &GameToSave{
		Nombre:         gs.nombre,
		Sexo:           gs.sexo,
		Dia:            gs.dia,
		Saldo:          gs.saldo,
		GananciaDiaria: gs.gananciaDiaria,
		ListaOpciones:  gs.listaOpciones,
	}
	x.save()
}
