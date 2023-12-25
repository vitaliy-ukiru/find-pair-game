package visual

import (
	"github.com/vitaliy-ukiru/find-pair-game/domain/entity"
	"github.com/vitaliy-ukiru/find-pair-game/domain/game"
)

type Card struct {
	ItemId entity.CardId
	Status game.CardStatus
}

type (
	CardRenderer interface {
		RenderCard(card Card) string
	}
	BoardRenderer interface {
		RenderBoard(rows [][]string) error
	}
	CardPlacer interface {
		CardAt(point entity.Point) Card
	}
)

type DeskRender struct {
	sizes entity.Sizes

	cardRenderer  CardRenderer
	boardRenderer BoardRenderer
	cardPlacer    CardPlacer
}

func (d *DeskRender) CardPlacer() CardPlacer {
	return d.cardPlacer
}

func NewDeskRender(
	sizes entity.Sizes,
	cardRenderer CardRenderer,
	boardRenderer BoardRenderer,
	cardPlacer CardPlacer,
) *DeskRender {
	return &DeskRender{
		sizes:         sizes,
		cardRenderer:  cardRenderer,
		boardRenderer: boardRenderer,
		cardPlacer:    cardPlacer,
	}
}

func (d *DeskRender) WithCardRenderer(renderer CardRenderer) *DeskRender {
	dr := *d
	dr.cardRenderer = renderer
	return &dr
}

func (d *DeskRender) WithBoardRenderer(renderer BoardRenderer) *DeskRender {
	dr := *d
	dr.boardRenderer = renderer
	return &dr
}

func (d *DeskRender) WithCardPlacer(placer CardPlacer) *DeskRender {
	dr := *d
	dr.cardPlacer = placer
	return &dr
}

func (d *DeskRender) Render() error {
	height := d.sizes.Height()
	width := d.sizes.Width()
	rows := make([][]string, 0, height)
	for y := 0; y < height; y++ {
		row := make([]string, 0, height)
		for x := 0; x < width; x++ {
			p := entity.NewPoint(x, y)
			card := d.cardPlacer.CardAt(p)
			cardContent := d.cardRenderer.RenderCard(card)
			row = append(row, cardContent)
		}
		rows = append(rows, row)
	}
	return d.boardRenderer.RenderBoard(rows)
}
