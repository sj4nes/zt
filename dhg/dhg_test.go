package dhg

import "os"
import "io/ioutil"
import "testing"
import "github.com/gkampitakis/go-snaps/snaps"
import "github.com/stretchr/testify/assert"

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
	snaps.MatchYAML(t, y)
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
	e1 := NewEdge()
	e2 := NewEdge()
	e3 := NewEdge()
	g.AddEdge(e1)
	g.AddEdge(e2)
	g.AddEdge(e3)
	return g
}

func TestGraphAddVertex(t *testing.T) {
	g := NewGraph()
	v := NewVertex()
	g.AddVertex(v)
	g.AddVertex(v)
	g.AddVertex(v)
	if g.Vertices != nil && g.Vertices.Len() != 3 {
		t.Errorf("expected 1 vertex, got %v", len(*g.Vertices))
	}
	if g.Vertices != nil && (*g.Vertices)[0] != *v {
		t.Errorf("unexpected vertex: %v", (*g.Vertices)[0])
	}
}

func TestGraphSaveYAML(t *testing.T) {
	g := graphFixture()
	err := g.SaveYAML("test.yaml")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	defer func() {
		err := os.Remove("test.yaml")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}()
	y, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	snaps.MatchYAML(t, y)
}

func TestLoadGraphYAML(t *testing.T) {
	g := graphFixture()
	err := g.SaveYAML("test.yaml")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	defer func() {
		err := os.Remove("test.yaml")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}()
	g2, err := LoadGraphYAML("test.yaml")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	y1, err := g.ToYAML()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	y2, err := g2.ToYAML()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.Equal(t, y1, y2)
}
