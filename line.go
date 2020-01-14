package templateutil

import (
	"fmt"
	"strings"
	"text/template/parse"
)

var (
	NodeNilError = fmt.Errorf("node is nil")
)

func Line(root, target parse.Node) (int, error) {
	if root == nil {
		return 0, fmt.Errorf("root: %w", NodeNilError)
	}
	if target == nil {
		return 0, fmt.Errorf("target: %w", NodeNilError)
	}
	line := 0
	Inspect(root, func(n parse.Node) bool {
		if n == nil {
			return false
		}
		if target.Position() > n.Position() {
			return true
		}
		if tn, ok := n.(*parse.TextNode); ok {
			if strings.Contains(string(tn.Text), "\n") {
				line++
			}
		}
		if tn, ok := n.(*parse.StringNode); ok {
			if strings.Contains(tn.Text, "\n") {
				line++
			}
		}
		return true
	})
	return line, nil
}
