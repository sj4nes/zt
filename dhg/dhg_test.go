package dhg

import "testing"
import "fmt"

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
	g := graphFixture()
	y, err := g.ToYAML()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expected := fmt.Sprintf(
		`id: "16_1"
vertices:
- id: "16_2"
- id: "16_3"
- id: "16_4"
`)

	if string(y) != expected {
		t.Errorf("unexpected yaml: \n%v\nvs:\n%v", string(y), expected)
	}
}

func TestGraphFmt(t *testing.T) {
	g := NewGraph()
	if g.Fmt() != " TODO" {
		t.Errorf("unexpected fmt: %v", g.Fmt())
	}
}

func graphFixture() *Graph {
	g := NewGraph()
	v1 := NewVertex()
	v2 := NewVertex()
	v3 := NewVertex()
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddVertex(v3)
	return g
}

func TestGraphAddVertex(t *testing.T) {
	g := NewGraph()
	v := NewVertex()
	g.AddVertex(v)
	g.AddVertex(v)
	g.AddVertex(v)
	if g.Vertices != nil && len(*g.Vertices) != 1 {
		t.Errorf("expected 1 vertex, got %v", len(*g.Vertices))
	}
	if g.Vertices != nil && (*g.Vertices)[0] != *v {
		t.Errorf("unexpected vertex: %v", (*g.Vertices)[0])
	}
}
