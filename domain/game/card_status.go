package game

import "github.com/vitaliy-ukiru/find-pair-game/domain/game/internal"

type CardStatus int

const (
	CardInvalid CardStatus = -1
	CardHide               = CardStatus(internal.CellHided)
	CardOpen               = CardStatus(internal.CellOpened)
	CardGuess              = CardStatus(internal.CellGuessed)
)
