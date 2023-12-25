package internal

import (
	"math/rand"

	"github.com/vitaliy-ukiru/find-pair-game/domain/entity"
)

func Generate(cards []entity.CardId, sizes entity.Sizes) map[entity.Point]*Cell {
	cells := make(map[entity.Point]*Cell)
	countCards := len(cards)
	height := sizes.Height()
	width := sizes.Width()

	usedCards := make(map[entity.CardId]int)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var card entity.CardId
			for {
				card = cards[rand.Intn(countCards)]

				if using := usedCards[card]; using < 2 {
					using++
					usedCards[card] = using
					break
				}
			}

			cells[entity.NewPoint(x, y)] = NewCell(card)

		}
	}
	return cells

}
