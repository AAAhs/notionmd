package converter

import (
	"github.com/dstotijn/go-notion"
	"github.com/gomarkdown/markdown/ast"
)

func isHorizontalRule(node ast.Node) bool {
	_, ok := node.(*ast.HorizontalRule)
	return ok
}

func convertHorizontalRule() notion.Block {
	return notion.DividerBlock{}
}
