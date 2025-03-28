package dhg

import "testing"

func TestZeroGraph(t *testing.T) {
	g := &Graph{}
	if g.Vertices != nil {
		t.Errorf("expected nil vertices, got %v", g.Vertices)
	}
	if g.Edges != nil {
		t.Errorf("expected nil edges, got %v", g.Edges)
	}
}

func TestGraphToYAML(t *testing.T) {
	g := NewGraph()
	y, err := g.ToYAML()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if string(y) != "id: \"16_1\"\nvalue: null\nvertices: null\nedges: null\n" {
		t.Errorf("unexpected yaml: %v", string(y))
	}
}

func TestGraphFmt(t *testing.T) {
	g := NewGraph()
	if g.Fmt() != " TODO" {
		t.Errorf("unexpected fmt: %v", g.Fmt())
	}
}

func TestGraphAddVertex(t *testing.T) {
	g := NewGraph()
	v := NewVertex()
	g.AddVertex(v)
	if g.Vertices != nil && len(*g.Vertices) != 1 {
		t.Errorf("expected 1 vertex, got %v", len(*g.Vertices))
	}
	if g.Vertices != nil && (*g.Vertices)[0] != *v {
		t.Errorf("unexpected vertex: %v", (*g.Vertices)[0])
	}
}
