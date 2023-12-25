package game

import (
	"errors"

	"github.com/vitaliy-ukiru/find-pair-game/domain/entity"
	"github.com/vitaliy-ukiru/find-pair-game/domain/game/internal"
)

type ClickStatus int

const (
	Ignore ClickStatus = iota
	Open
	Wrong
	Guess
	Finish
)

type Card struct {
	Id    entity.CardId
	Point entity.Point
}

type ClickResult struct {
	Result  ClickStatus
	Current Card
	Prev    *Card
}

var ErrOutOfRange = errors.New("point in out of board")

func (g *Game) MakeClick(p entity.Point) (ClickResult, error) {
	if !g.sizes.Contains(p) {
		return ClickResult{}, ErrOutOfRange
	}
	if g.lastPoint == nil {
		return g.processFirstClick(p), nil
	}
	result := g.processSecondClick(p)
	if result.Result == Guess && len(g.guessed) == len(g.cards) {
		result.Result = Finish
	}
	return result, nil

}

func (g *Game) processFirstClick(p entity.Point) ClickResult {
	cell := g.cells[p]
	if cell.Status == internal.CellGuessed {
		return ClickResult{
			Result: Ignore,
		}
	}

	g.lastPoint = &p
	cell.Open()
	return ClickResult{
		Result: Open,
		Current: Card{
			Id:    cell.ItemId,
			Point: p,
		},
	}
}

func (g *Game) processSecondClick(p entity.Point) ClickResult {
	prevPoint := g.lastPoint
	lastPoint := *prevPoint

	prevCell := g.cells[lastPoint]
	currentCell := g.cells[p]

	prevCard := &Card{
		Point: lastPoint,
		Id:    prevCell.ItemId,
	}

	currentCard := Card{
		Id:    currentCell.ItemId,
		Point: p,
	}

	if currentCell.IsOpened() {
		return ClickResult{
			Result:  Ignore,
			Current: currentCard,
			Prev:    prevCard,
		}
	}

	g.lastPoint = nil
	if prevCell.IsGuess(currentCell) {
		prevCell.Guess()
		currentCell.Guess()
		g.guess(currentCard.Id)
		return ClickResult{
			Result:  Guess,
			Current: currentCard,
			Prev:    prevCard,
		}
	}

	prevCell.Hide()
	currentCell.Hide()
	return ClickResult{
		Result:  Wrong,
		Current: currentCard,
		Prev:    prevCard,
	}
}
