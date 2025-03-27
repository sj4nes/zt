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
	g := &Graph{}
	y, err := g.ToYAML()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if string(y) != "value: null\nvertices: null\nedges: null\n" {
		t.Errorf("unexpected yaml: %v", string(y))
	}
}
