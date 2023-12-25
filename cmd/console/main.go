package main

import (
	"bufio"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/vitaliy-ukiru/find-pair-game/domain/entity"
	"github.com/vitaliy-ukiru/find-pair-game/domain/game"
	"github.com/vitaliy-ukiru/find-pair-game/pkg/desk/console"
	"github.com/vitaliy-ukiru/find-pair-game/pkg/desk/console/consolerender"
	"github.com/vitaliy-ukiru/find-pair-game/pkg/visual"
)

var cards = visual.Cards{
	//"CAR":        "🚗",
	//"CHICKEN":    "🐤",
	//"BALL":       "🏐",
	"EIGHT_BALL": "🎱",
	"APPLE":      "🍏",
}

func main() {
	g := game.New(
		entity.NewSizes(2, 2),
		cards.Ids(),
	)
	if err := g.Init(); err != nil {
		panic(err)
	}

	time.Sleep(3 * time.Second)
	desk := console.New(
		g,
		bufio.NewReader(os.Stdin),
		visual.NewGameCardPlacer(g),
		consolerender.NewCardBracketsRenderer(cards),
		consolerender.NewTableRenderer(func() *tablewriter.Table {
			t := tablewriter.NewWriter(os.Stdout)
			t.SetCaption()
			t.SetColWidth(7)
			return t
		}),
		console.English,
	)
	desk.Play()
}
