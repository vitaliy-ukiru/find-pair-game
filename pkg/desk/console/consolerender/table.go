package consolerender

import (
	"github.com/olekukonko/tablewriter"
)

type TableRenderer struct {
	newTable func() *tablewriter.Table
}

func NewTableRenderer(newTable func() *tablewriter.Table) *TableRenderer {
	return &TableRenderer{newTable: newTable}
}

func (t TableRenderer) RenderBoard(rows [][]string) error {
	table := t.newTable()
	for _, row := range rows {
		table.Append(row)
	}
	table.Render()
	return nil
}
