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
