package console

import (
	"fmt"
	"os"

	"github.com/vitaliy-ukiru/find-pair-game/pkg/visual"
)

func (d *Desk) renderDefault() {
	d.render(true, "", nil)
}

func (d *Desk) renderMessage(msg string) {
	d.render(true, msg, nil)
}

func (d *Desk) render(clearDisplay bool, msg string, customRender *visual.DeskRender) {
	renderer := d.renderer
	if customRender != nil {
		renderer = customRender
	}

	if clearDisplay {
		clearConsole()
	}

	err := renderer.Render()
	if err != nil {
		fmt.Printf("%s: %v\n", d.texts.withError(d.texts.RenderError), err)
		os.Exit(-1)
	}
	fmt.Println()
	if msg != "" {
		fmt.Println(msg)
	}
}

//	func (d *Desk) render(items boardItems, clearDisplay bool, msg string) {
//		if clearDisplay {
//			clearConsole()
//		}
//		table := tablewriter.NewWriter(os.Stdout)
//		sizes := d.g.Sizes()
//		for row := 0; row < sizes.Height(); row++ {
//			itemsView := make([]string, 0, sizes.Width())
//			for col := 0; col < sizes.Width(); col++ {
//				cell := items[game.NewPoint(col, row)]
//				itemsView = append(itemsView, d.formatItem(cell))
//			}
//			table.Append(itemsView)
//		}
//		table.SetColWidth(7)
//		table.Render()
//		fmt.Println()
//		if msg != "" {
//			fmt.Println(msg)
//		}
//	}

// func (d *Desk) formatItem(item game.BoardItem) string {
//
// }
func clearConsole() {
	fmt.Print("\x1bc")
}
