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
	assert.Equal(t, g.Fmt(), " TODO ("+g.Id+")")
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
	g.AddVertex(NewVertex())
	g.AddVertex(NewVertex())
	g.AddVertex(NewVertex())
	assert.NotNil(t, g.Vertices)
	assert.Equal(t, 3, g.Vertices.Len())
}

func TestGraphSaveYAML(t *testing.T) {
	g := graphFixture()
	err := g.SaveYAML("test.yaml")
	assert.NoError(t, err)
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

func TestGraphDatumFmt(t *testing.T) {
	g := NewGraph()
	assert.Equal(t, g.Fmt(), " TODO ("+g.Id+")")
}

func TestVerticesLen(t *testing.T) {
	g := graphFixture()
	assert.Equal(t, g.Vertices.Len(), 3)
}

func TestEdgesLen(t *testing.T) {
	g := graphFixture()
	assert.Equal(t, g.Edges.Len(), 3)
}
