package converter

import "github.com/gomarkdown/markdown/ast"

func isHardbreak(node ast.Node) bool {
	_, ok := node.(*ast.Hardbreak)
	return ok
}

func isSoftbreak(node ast.Node) bool {
	_, ok := node.(*ast.Softbreak)
	return ok
}

func isLineBreak(node ast.Node) bool {
	return isHardbreak(node) || isSoftbreak(node)
}
