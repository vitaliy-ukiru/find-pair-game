package internal

import (
	"github.com/vitaliy-ukiru/find-pair-game/domain/entity"
)

type CellStatus byte

const (
	CellHided CellStatus = iota
	CellOpened
	CellGuessed
)

type Cell struct {
	ItemId entity.CardId
	Status CellStatus
}

func NewCell(item entity.CardId) *Cell {
	return &Cell{ItemId: item}
}

func (c *Cell) Open()  { c.Status = CellOpened }
func (c *Cell) Hide()  { c.Status = CellHided }
func (c *Cell) Guess() { c.Status = CellGuessed }

func (c *Cell) IsOpened() bool  { return c.Status != CellHided }
func (c *Cell) IsGuessed() bool { return c.Status == CellGuessed }

func (c *Cell) IsGuess(c2 *Cell) bool {
	return c.ItemId == c2.ItemId
}
