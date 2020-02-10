package templateutil

import (
	"fmt"
	"text/template/parse"
)

type Visitor interface {
	Visit(node parse.Node) (w Visitor)
}

func Walk(node parse.Node, f Visitor) {
	if f = f.Visit(node); f == nil {
		return
	}
	switch n := node.(type) {
	case *parse.ActionNode:
		Walk(n.Pipe, f)
	case *parse.BoolNode:
	case *parse.BranchNode:
		Walk(n.Pipe, f)
		Walk(n.List, f)
		if n.ElseList != nil {
			Walk(n.ElseList, f)
		}
	case *parse.ChainNode:
		Walk(n.Node, f)
	case *parse.CommandNode:
		for _, c := range n.Args {
			Walk(c, f)
		}
	case *parse.DotNode:
	case *parse.FieldNode:
	case *parse.IdentifierNode:
	case *parse.IfNode:
		Walk(n.Pipe, f)
		Walk(n.List, f)
		if n.ElseList != nil {
			Walk(n.ElseList, f)
		}
	case *parse.ListNode:
		for _, nn := range n.Nodes {
			Walk(nn, f)
		}
	case *parse.NilNode:
	case *parse.NumberNode:
	case *parse.PipeNode:
		for _, d := range n.Decl {
			Walk(d, f)
		}
		for _, c := range n.Cmds {
			Walk(c, f)
		}
	case *parse.RangeNode:
		Walk(n.Pipe, f)
		Walk(n.List, f)
		if n.ElseList != nil {
			Walk(n.ElseList, f)
		}
	case *parse.StringNode:
	case *parse.TemplateNode:
		if n.Pipe != nil {
			Walk(n.Pipe, f)
		}
	case *parse.TextNode:
	case *parse.VariableNode:
	case *parse.WithNode:
		Walk(n.Pipe, f)
		Walk(n.List, f)
		if n.ElseList != nil {
			Walk(n.ElseList, f)
		}
	default:
		panic(fmt.Sprintf("template.Walk: unexpected node type %T", n))
	}
	f.Visit(nil)
}

type inspector func(parse.Node) bool

func (f inspector) Visit(node parse.Node) Visitor {
	if f(node) {
		return f
	}
	return nil
}
func Inspect(node parse.Node, f func(parse.Node) bool) {
	Walk(node, inspector(f))
}
