package gameState

type NextStatus struct {
	GameOver bool
}
type GameState interface {
	Finish()
	ShowMenu()
	NextDay()
	ShowSummary()
	ExecSelectedOption() *NextStatus
}

var (
	SAVE_FILE_NAME string = "state-v1.gob"
)
