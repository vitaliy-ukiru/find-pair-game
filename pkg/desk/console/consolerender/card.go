package consolerender

import (
	"fmt"

	"github.com/vitaliy-ukiru/find-pair-game/domain/game"
	"github.com/vitaliy-ukiru/find-pair-game/pkg/visual"
)

type CardBracketsRenderer struct {
	cards visual.Cards
}

func NewCardBracketsRenderer(cards visual.Cards) *CardBracketsRenderer {
	return &CardBracketsRenderer{cards: cards}
}

func (c *CardBracketsRenderer) RenderCard(card visual.Card) string {
	switch card.Status {
	case game.CardGuess:
		return fmt.Sprintf("[%s]", c.cards[card.ItemId])
	case game.CardOpen:
		return fmt.Sprintf("(%s)", c.cards[card.ItemId])
	}
	return "[ ]"
}
