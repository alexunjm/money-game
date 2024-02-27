package gameState

import (
	"bytes"
	"encoding/gob"
	"os"

	"github.com/alexunjm/money-game/core/game/option"
)

func Load() (GameState, error) {
	var state GameToSave
	// Leer el archivo
	datos, err := os.ReadFile(SAVE_FILE_NAME)
	if err != nil {
		return nil, err
	}

	// Crear un buffer a partir de los datos y un decoder
	buffer := bytes.NewBuffer(datos)
	decoder := gob.NewDecoder(buffer)

	// Deserializar en una instancia de GameToSave
	err = decoder.Decode(&state)
	if err != nil {
		return nil, err
	}

	optBuilder := option.NewOptionBuilder()
	optBuilder.WithOptions(state.ListaOpciones)
	x := &gameSate{
		nombre:         state.Nombre,
		sexo:           state.Sexo,
		dia:            state.Dia,
		saldo:          state.Saldo,
		gananciaDiaria: state.GananciaDiaria,

		constructorOpciones: optBuilder,
	}
	x.ReloadOptionsFromBuilder()
	return x, nil
}
