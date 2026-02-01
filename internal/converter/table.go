package converter

import (
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

func isTable(node ast.Node) bool {
	_, ok := node.(*ast.Table)
	return ok
}

func convertTable(node *ast.Table) *notion.TableBlock {
	if node == nil {
		return nil
	}

	var rows [][][]notion.RichText
	hasHeader := false

	for _, child := range node.GetChildren() {
		switch section := child.(type) {
		case *ast.TableHeader:
			hasHeader = true
			rows = append(rows, convertTableSection(section)...)
		case *ast.TableBody:
			rows = append(rows, convertTableSection(section)...)
		case *ast.TableFooter:
			rows = append(rows, convertTableSection(section)...)
		case *ast.TableRow:
			rows = append(rows, convertTableRow(section))
		}
	}

	tableWidth := maxTableWidth(rows)
	if tableWidth == 0 {
		return nil
	}

	padTableRows(rows, tableWidth)

	var tableRows []notion.Block
	for _, row := range rows {
		tableRows = append(tableRows, notion.TableRowBlock{
			Cells: row,
		})
	}

	return &notion.TableBlock{
		TableWidth:      tableWidth,
		HasColumnHeader: hasHeader,
		HasRowHeader:    false,
		Children:        tableRows,
	}
}

func convertTableSection(node ast.Node) [][][]notion.RichText {
	if node == nil {
		return nil
	}

	var rows [][][]notion.RichText
	for _, child := range node.GetChildren() {
		row, ok := child.(*ast.TableRow)
		if !ok {
			continue
		}
		rows = append(rows, convertTableRow(row))
	}

	return rows
}

func convertTableRow(node *ast.TableRow) [][]notion.RichText {
	if node == nil {
		return nil
	}

	var cells [][]notion.RichText
	for _, child := range node.GetChildren() {
		cell, ok := child.(*ast.TableCell)
		if !ok {
			continue
		}

		cells = append(cells, convertTableCell(cell))
		if cell.ColSpan > 1 {
			for i := 1; i < cell.ColSpan; i++ {
				cells = append(cells, []notion.RichText{})
			}
		}
	}

	return cells
}

func convertTableCell(node *ast.TableCell) []notion.RichText {
	if node == nil {
		return []notion.RichText{}
	}

	richText := convertChildNodesToRichText(node)
	if richText == nil {
		return []notion.RichText{}
	}

	return richText
}

func maxTableWidth(rows [][][]notion.RichText) int {
	tableWidth := 0
	for _, row := range rows {
		if len(row) > tableWidth {
			tableWidth = len(row)
		}
	}

	return tableWidth
}

func padTableRows(rows [][][]notion.RichText, tableWidth int) {
	for i := range rows {
		if len(rows[i]) >= tableWidth {
			continue
		}
		for j := len(rows[i]); j < tableWidth; j++ {
			rows[i] = append(rows[i], []notion.RichText{})
		}
	}
}
