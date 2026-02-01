package converter

import (
	"testing"

	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
	"github.com/stretchr/testify/assert"
)

func TestConvertChildNodesToRichText_Breaks(t *testing.T) {
	paragraph := &ast.Paragraph{
		Container: ast.Container{
			Children: []ast.Node{
				&ast.Text{Leaf: ast.Leaf{Literal: []byte("Line 1")}},
				&ast.Hardbreak{},
				&ast.Text{Leaf: ast.Leaf{Literal: []byte("Line 2")}},
				&ast.Softbreak{},
				&ast.Text{Leaf: ast.Leaf{Literal: []byte("Line 3")}},
			},
		},
	}

	expected := []notion.RichText{
		{
			Type:      notion.RichTextTypeText,
			Text:      &notion.Text{Content: "Line 1"},
			PlainText: "Line 1",
		},
		{
			Type:      notion.RichTextTypeText,
			Text:      &notion.Text{Content: "\n"},
			PlainText: "\n",
		},
		{
			Type:      notion.RichTextTypeText,
			Text:      &notion.Text{Content: "Line 2"},
			PlainText: "Line 2",
		},
		{
			Type:      notion.RichTextTypeText,
			Text:      &notion.Text{Content: "\n"},
			PlainText: "\n",
		},
		{
			Type:      notion.RichTextTypeText,
			Text:      &notion.Text{Content: "Line 3"},
			PlainText: "Line 3",
		},
	}

	result := convertChildNodesToRichText(paragraph)
	assert.Equal(t, expected, result)
}
