package gameState

import (
	"fmt"

	"github.com/alexunjm/money-game/core/game/option"
	"github.com/inancgumus/screen"
)

type gameSate struct {
	nombre, sexo          string
	dia                   int
	saldo, gananciaDiaria float32
	opcionElegida         option.Option
	constructorOpciones   option.OptionBuilder

	listaOpciones []option.Option
}

// ShowMenu implements GameState.
func (gs *gameSate) ShowMenu() {
	gs.opcionElegida = option.AskWithOptions("Elige una de las siguientes opciones", gs.listaOpciones)
	fmt.Printf("Opcion: `%s`\n", gs.opcionElegida.Text)
}

// ExecSelectedOption implements GameState.
func (gs *gameSate) ExecSelectedOption() *NextStatus {
	fmt.Printf("\tDia: %v\n", gs.dia)
	switch gs.opcionElegida.Id {
	case option.AVANZAR:
		gs.NextDay()
	case option.RESUMEN:
		screen.Clear()
		gs.ShowSummary()
	default:
		return &NextStatus{GameOver: true}
	}
	return &NextStatus{GameOver: false}
}

// Finish implements GameState.
func (gs *gameSate) Finish() {
	Save(gs)
	fmt.Println("Juego terminado")
}

// NextDay implements GameState.
func (gs *gameSate) NextDay() {
	gs.add(gs.gananciaDiaria)
	gs.dia++
	fmt.Println("")
	fmt.Printf("********************\n\n")
}

// ShowSummary implements GameState.
func (gs *gameSate) ShowSummary() {

	fmt.Println("")
	fmt.Printf("********************\n*\n")
	fmt.Printf("* Iniciando Dia: %v\n", gs.dia)
	fmt.Printf("*\tNombre: %v\n", gs.nombre)
	fmt.Printf("*\tSexo: %v\n", gs.sexo)
	fmt.Printf("*\tSaldo: %v\n", gs.saldo)
	fmt.Printf("*\tGanancia estimada del dia: %v\n", gs.gananciaDiaria)
	fmt.Printf("*\n********************\n\n\n")
}

func (gs *gameSate) init() {

	gs.nombre = option.Ask("Cual es tu nombre?")
	fmt.Printf("Hola, %s!\n", gs.nombre)

	gender := option.NewGenderOption()
	gs.sexo = option.AskWithOptions("Elige una de las siguientes opciones", gender.ToList()).Text
	fmt.Printf("Opcion: `%s`\n", gs.sexo)

	gs.constructorOpciones = option.NewOptionBuilder()
	gs.ReloadOptionsFromBuilder()

}

func (gs *gameSate) add(val float32) {
	fmt.Printf("\n\tAntes %v", gs.saldo)
	fmt.Printf("\n\tGanancia del dia %v", val)
	gs.saldo += val
	fmt.Printf("\n\tDespues %v\n\n", gs.saldo)
}

func (gs *gameSate) ReloadOptionsFromBuilder() {
	gs.listaOpciones = gs.constructorOpciones.Build().ToList()
}

func NewState() GameState {
	x, err := Load()
	if err == nil {
		return x
	}
	fmt.Println(err)

	y := &gameSate{
		dia:            1,
		saldo:          0,
		gananciaDiaria: 1,
	}
	y.init()
	return y
}
