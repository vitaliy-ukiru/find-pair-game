package game

import "github.com/vitaliy-ukiru/find-pair-game/domain/game/internal"

type CardStatus int

const (
	CardClose CardStatus = iota
	CardOpen
	CardGuess
)

func cardStatusFromCell(status internal.CellStatus) CardStatus {
	return CardStatus(status)
}
