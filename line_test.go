package templateutil

import (
	"errors"
	"html/template"
	"testing"
)

func TestLine(t *testing.T) {
	testcases := []struct {
		title string
		tpl   string
		line  int
	}{
		{
			title: "normal",
			tpl: `this is {{ println "template" }}
    hogehuga{{ println }}
`,
			line: 1,
		},
		{
			title: "oneline",
			tpl:   `this is {{ println "template" }}hogehuga{{ println }}`,
			line:  0,
		},
		{
			title: "double newline",
			tpl: `this is {{ println "template" }}

    hogehuga{{ println }}`,
			line: 2,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.title, func(t *testing.T) {
			tmpl, err := template.New("Ident").Parse(tc.tpl)
			if err != nil {
				t.Fatalf("parsing template %s", err)
			}
			t.Log(tmpl.Tree.Root.Nodes[len(tmpl.Tree.Root.Nodes)-1])
			l, err := Line(tmpl.Tree.Root, tmpl.Tree.Root.Nodes[len(tmpl.Tree.Root.Nodes)-1])
			if err != nil {
				t.Fatal(err)
			}
			if l != tc.line {
				t.Errorf("result=%d != expected=%d", l, tc.line)
			}
		})
	}
}

func TestLine_NoText(t *testing.T) {
	tpl := ``
	tmpl, err := template.New("Ident").Parse(tpl)
	if err != nil {
		t.Fatalf("parsing template %s", err)
	}
	_, err = Line(tmpl.Tree.Root, nil)
	if err == nil {
		t.Fatal("expected error but no")
	}
	originalErr := errors.Unwrap(err)
	if originalErr == nil {
		t.Fatal("expected error but no")
	}
	if !errors.Is(originalErr, NodeNilError) {
		t.Fatal("unexpected error: ", err)
	}
}
