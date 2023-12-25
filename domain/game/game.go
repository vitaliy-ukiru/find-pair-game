package game

import (
	"errors"
	"slices"

	"github.com/vitaliy-ukiru/find-pair-game/domain/entity"
	"github.com/vitaliy-ukiru/find-pair-game/domain/game/internal"
	"github.com/vitaliy-ukiru/find-pair-game/internal/container"
)

type Game struct {
	sizes     entity.Sizes
	cards     []entity.CardId
	cells     map[entity.Point]*internal.Cell
	lastPoint *entity.Point
	guessed   container.Hashset[entity.CardId]
}

func New(sizes entity.Sizes, cards []entity.CardId) *Game {
	return &Game{
		sizes: sizes,
		cards: cards,
	}
}

var ErrInvalidBoardSizes = errors.New("count of cells must be even")

var ErrInvalidCountOfCards = errors.New("count of items must be half of cells")

func (g *Game) Init() error {
	if err := g.validate(); err != nil {
		return err
	}

	g.cells = internal.Generate(g.cards, g.sizes)

	g.guessed = make(container.Hashset[entity.CardId])
	return nil
}

func (g *Game) validate() error {
	cellsCount := g.sizes.CellsCount()
	if cellsCount%2 != 0 {
		return ErrInvalidBoardSizes
	}
	if len(g.cards) != cellsCount/2 {
		return ErrInvalidCountOfCards
	}
	return nil
}

func (g *Game) Sizes() entity.Sizes           { return g.sizes }
func (g *Game) Cards() []entity.CardId        { return slices.Clone(g.cards) }
func (g *Game) GuessedCards() []entity.CardId { return g.guessed.ToSlice() }

type BoardItem struct {
	entity.Point
	Card   entity.CardId
	Status CardStatus
}

func (g *Game) VisibleItems() []BoardItem {
	height := g.sizes.Height()
	width := g.sizes.Width()

	opened := 0
	for _, cell := range g.cells {
		if cell.IsOpened() {
			opened++
		}
	}

	items := make([]BoardItem, 0, opened)

	// go as 2 loops, because want to have sort
	// but this way have less time complexity for sorting.
	// Sorting is needed because maps store items in random order.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			point := entity.NewPoint(x, y)

			cell := g.cells[point]
			if cell.IsOpened() {
				items = append(items, BoardItem{
					Point:  point,
					Status: cardStatusFromCell(cell.Status),
					Card:   cell.ItemId,
				})
			}
		}
	}
	return items
}

func (g *Game) ItemAt(p entity.Point) BoardItem {
	if !g.sizes.Contains(p) {
		return BoardItem{
			Status: CardInvalid,
		}
	}
	cell := g.cells[p]
	item := BoardItem{
		Point:  p,
		Status: cardStatusFromCell(cell.Status),
	}
	if cell.IsOpened() {
		item.Card = cell.ItemId
	}
	return item
}

func (g *Game) guess(id entity.CardId) {
	g.guessed.Add(id)
}
