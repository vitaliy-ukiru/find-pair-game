package visual

import (
	"github.com/vitaliy-ukiru/find-pair-game/domain/entity"
	"github.com/vitaliy-ukiru/find-pair-game/domain/game"
)

type GameCardPlacer struct {
	g *game.Game
}

func NewGameCardPlacer(g *game.Game) *GameCardPlacer {
	return &GameCardPlacer{g: g}
}

func (g *GameCardPlacer) CardAt(point entity.Point) Card {
	item := g.g.ItemAt(point)
	return Card{
		ItemId: item.Card,
		Status: item.Status,
	}
}

type WrongCardPlacer struct {
	Current game.Card
	Prev    game.Card
	parent  CardPlacer
}

func NewWrongCardPlacer(current game.Card, prev game.Card, parent CardPlacer) *WrongCardPlacer {
	return &WrongCardPlacer{Current: current, Prev: prev, parent: parent}
}

func (w *WrongCardPlacer) CardAt(point entity.Point) Card {
	if point == w.Current.Point {
		return Card{
			ItemId: w.Current.Id,
			Status: game.CardOpen,
		}
	}
	if point == w.Prev.Point {
		return Card{
			ItemId: w.Prev.Id,
			Status: game.CardOpen,
		}
	}

	return w.parent.CardAt(point)
}
