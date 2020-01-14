package templateutil

import (
	"errors"
	"html/template"
	"testing"
)

func TestLine(t *testing.T) {
	tpl := `this is {{ println "template" }}
    hogehuga{{ println }}
`
	tmpl, err := template.New("Ident").Parse(tpl)
	if err != nil {
		t.Fatalf("parsing template %s", err)
	}
	l, err := Line(tmpl.Tree.Root, tmpl.Tree.Root.Nodes[len(tmpl.Tree.Root.Nodes)-2])
	if err != nil {
		t.Fatal(err)
	}
	expected := 1
	if l != expected {
		t.Errorf("result=%d != expected=%d", l, expected)
	}
}

func TestLine_NoLine(t *testing.T) {
	tpl := `this is {{ println "template" }}hogehuga{{ println }}`
	tmpl, err := template.New("Ident").Parse(tpl)
	if err != nil {
		t.Fatalf("parsing template %s", err)
	}
	l, err := Line(tmpl.Tree.Root, tmpl.Tree.Root.Nodes[len(tmpl.Tree.Root.Nodes)-2])
	if err != nil {
		t.Fatal(err)
	}
	expected := 0
	if l != expected {
		t.Errorf("result=%d != expected=%d", l, expected)
	}
}
func TestLine_TwoNewLine(t *testing.T) {
	tpl := `this is {{ println "template" }}

hogehuga{{ println }}`
	tmpl, err := template.New("Ident").Parse(tpl)
	if err != nil {
		t.Fatalf("parsing template %s", err)
	}
	l, err := Line(tmpl.Tree.Root, tmpl.Tree.Root.Nodes[len(tmpl.Tree.Root.Nodes)-2])
	if err != nil {
		t.Fatal(err)
	}
	expected := 2
	if l != expected {
		t.Errorf("result=%d != expected=%d", l, expected)
	}
}

func TestLine_NoText(t *testing.T) {
	tpl := ``
	tmpl, err := template.New("Ident").Parse(tpl)
	if err != nil {
		t.Fatalf("parsing template %s", err)
	}
	_, err = Line(tmpl.Tree.Root, nil)
	t.Log(err)
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
