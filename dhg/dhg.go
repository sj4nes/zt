// Package dhg provides a directed hypergraph implementation.
// It includes data structures for vertices, edges, and graphs, as well as
// methods for creating, manipulating, and serializing these structures.
package dhg

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
	. "zt/quality"
)

const (
	// EDGE_SYMBOL        = "\u2ad8"     // ⫘
	// HEAD_SYMBOL        = "\u2b44"     // ⭄
	// TAIL_SYMBOL        = "\u297a"     // ⥺
	// 󰴣 U+f0d23 Perspective less
	// 󰴤 U+f0d24 Perspective more
	// 󰴶 U+f0d36 Skew less
	// 󰴷 U+f0d37 Skew greater
	// 󱗜 U+f15dc Circle box
	BOOLEAN_SYMBOL     = "\uea8f"     // 
	DATA_SYMBOL        = "\U000F0316" // 󰌖
	EDGE_SYMBOL        = "\U000f0d23" // 󰴣 = {(󱗜,󱗜)} ?
	EMPTY_GRAPH_SYMBOL = "\u29b2"     // ⦲
	EMPTY_SET_SYMBOL   = "\u2205"     // ∅
	GRAPH_SYMBOL       = "\ue662"     // 
	HEAD_SYMBOL        = "\U000f0d37" // 󰴷
	INTEGER_SYMBOL     = "\u2124"     // ℤ
	MIRROR_SYMBOL      = "\uf41a"     //  = 󰴤 = {(󱗜,󱗜)} ?
	NO_DATA_SYMBOL     = "\U000F0acc" // 󰫌
	POINTER_SYMBOL     = "\U000f1484" // 󱒄
	REAL_SYMBOL        = "\u211d"     // ℝ
	STRING_SYMBOL      = "\ueb8d"     // 
	TAIL_SYMBOL        = "\U000f0d36" // 󰴶
	VERTEX_SYMBOL      = "\U000f15dc" // 󱗜 (a point in an area) = {(󰴷,󰴶)} ?
)

// These are "private" globals that support NextId().
var pLAST_ID_MX sync.Mutex
var pLAST_SERIAL_ID int64 = 0
var pRNGSRC *rand.Rand

// init initializes the random number generator with a seed based on the
// current time or 42 during debugging/testing.
func init() {
	pRNGSRC = rand.New(rand.NewSource(time.Now().UnixNano()))
	if os.Getenv("DEBUG") == "1" {
		pRNGSRC = rand.New(rand.NewSource(42))
		x := pRNGSRC.Intn(30)
		Assert(x == 5, "static random seed is not set correctly: "+strconv.Itoa(x))
	}
}

// NextId generates a new unique ID for graph elements.
func NextId() string {
	pLAST_ID_MX.Lock()
	defer pLAST_ID_MX.Unlock()
	pLAST_SERIAL_ID++
	conv := strconv.FormatInt(pLAST_SERIAL_ID, 36)
	// Randomize part of the id to avoid collisions
	r := pRNGSRC.Intn(1e9)
	if os.Getenv("DEBUG") == "1" {
		r = 42 // For testing
	}
	nonce := strconv.FormatInt(int64(r), 36)
	return nonce + "_" + conv
}

//  U+f0a9 Arrow circle right
//  U+f0a8 Arrow circle left
// 󱡓 U+f1853 Circle opacity
// 󰴣 = {( , )}
// 󱗜 = {(󰴷,󰴶)}

// Element is a common interface for all elements in a graph
type Element interface {
	element()
}

type Edge struct {
	Id    string    `yaml:"id"`
	Value *Datum    `yaml:"value,omitempty"`
	Rels  *Vertices `yaml:"rels,omitempty"`
}

func NewEdge() *Edge {
	return &Edge{Id: NextId()}
}

func (*Edge) element() {}

type Vertex struct {
	Id    string `yaml:"id"`
	Value *Datum `yaml:"value,omitempty"`
	Rels  *Edges `yaml:"rels,omitempty"`
}

func NewVertex() *Vertex {
	return &Vertex{Id: NextId()}
}

func (*Vertex) element() {}

type Rel struct {
	Head *Element `yaml:"head,omitempty"`
	Tail *Element `yaml:"tail,omitempty"`
}

type Rels []Rel

type Vertices map[string]Vertex

// NewVertices creates a new Vertices collection.
func NewVertices() *Vertices {
	v := make(Vertices)
	return &v

}

// Add adds vertices to the collection.
func (v *Vertices) Add(vs ...Vertex) {
	for _, vertex := range vs {
		(*v)[vertex.Id] = vertex
	}
}

// Len returns the number of vertices in the collection.
func (v *Vertices) Len() int {
	return len(*v)
}

// Edges is a collection of edges.
type Edges map[string]Edge

// NewEdges creates a new Edges collection.
func NewEdges() *Edges {
	v := make(Edges)
	return &v
}

// Add adds edges to the collection.
func (e *Edges) Add(es ...Edge) {
	for _, edge := range es {
		(*e)[edge.Id] = edge
	}
}

// Len returns the number of edges in the collection.
func (e *Edges) Len() int {
	return len(*e)
}

type Graph struct {
	Id          string    `yaml:"id"`
	Vertices    *Vertices `yaml:"vertices,omitempty"`
	Edges       *Edges    `yaml:"edges,omitempty"`
	StringDatum `yaml:"dstr,omitempty"`
}

func NewGraph() *Graph {
	g := &Graph{Id: NextId()}
	label := NewStringDatum("TODO")
	g.Set(label)
	return g
}

func (g *Graph) Set(v StringDatum) {
	g.DStr = v.DStr
}

// ToYAML serializes the graph to YAML format.
func (g Graph) ToYAML() ([]byte, error) {
	return yaml.Marshal(g)
}

func (g *Graph) AddVertex(v *Vertex) {
	if g.Vertices == nil {
		g.Vertices = NewVertices()
	}
	g.Vertices.Add(*v)
}

func (g *Graph) AddEdge(e *Edge) {
	if g.Edges == nil {
		g.Edges = NewEdges()
	}
	g.Edges.Add(*e)
}

// idFmt formats the ID of a graph element for display. Used for graphs, vertices, and edges.
func idFmt(id string) string {
	return "(" + id + ")"
}

func (d *Graph) Fmt() string {
	return fmt.Sprintf("%s %s %s", GRAPH_SYMBOL, d.DStr, idFmt(d.Id))
}

type Datum interface {
	datum()
	Fmt() string
}

type GraphDatum struct {
	DGraph *Graph `yaml:"graph,omitempty"`
}

//go:noinline
func (GraphDatum) datum() {}

func (d *GraphDatum) Fmt() string {
	if d.DGraph == nil {
		return EMPTY_GRAPH_SYMBOL
	}
	return d.DGraph.Fmt()
}

type NilDatum struct {
	DNil bool `yaml:"nil" default:"true"`
}

//go:noinline
func (*NilDatum) datum() {}

func (*NilDatum) Fmt() string {
	return NO_DATA_SYMBOL
}

type BoolDatum struct {
	DBool bool `yaml:"bool" default:"false"`
}

//go:noinline
func (*BoolDatum) datum() {}

func (d *BoolDatum) Fmt() string {
	if d.DBool {
		return BOOLEAN_SYMBOL + "true"
	}
	return BOOLEAN_SYMBOL + "false"
}

type IntegerDatum struct {
	DInt int64
}

//go:noinline
func (*IntegerDatum) datum() {}

func (d *IntegerDatum) Fmt() string {
	return INTEGER_SYMBOL + " " + strconv.FormatInt(d.DInt, 10)
}

type FloatDatum struct {
	DFloat float64
}

//go:noinline
func (*FloatDatum) datum() {}

func (d *FloatDatum) Fmt() string {
	return REAL_SYMBOL + " " + strconv.FormatFloat(d.DFloat, 'f', -1, 64)
}

type StringDatum struct {
	DStr string `yaml:"dstr"`
}

func NewStringDatum(value string) StringDatum {
	return StringDatum{DStr: value}
}

//go:noinline
func (*StringDatum) datum() {}

func (d StringDatum) Fmt() string {
	return STRING_SYMBOL + " " + d.DStr
}

type VertexDatum struct {
	DVertex *Vertex
}

//go:noinline
func (*VertexDatum) datum() {}

func (*VertexDatum) Fmt() string {
	return VERTEX_SYMBOL + " " + "TODO"
}

type EdgeDatum struct {
	DEdge *Edge
}

//go:noinline
func (*EdgeDatum) datum() {}

func (*EdgeDatum) Fmt() string {
	return EDGE_SYMBOL + " " + "TODO"
}

type RuneDatum struct {
	Type  string `yaml:"type" default:"rune"`
	Value rune
}

//go:noinline
func (*RuneDatum) datum() {}

type TimestampDatum struct {
	Type  string `yaml:"type" default:"string"`
	Value time.Time
}

//go:noinline
func (*TimestampDatum) datum() {}

// Fmt returns a formatted string for the provided datum. Probably a ridululous
// generic since there is already an interface. :)
func Fmt[T Datum](d T) string {
	return d.Fmt()
}

// SaveYAML saves the graph to a YAML file.
func (g *Graph) SaveYAML(filename string) error {
	data, err := g.ToYAML()
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// LoadGraphYAML loads a graph from a YAML file.
func LoadGraphYAML(filename string) (*Graph, error) {
	y, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	g := &Graph{}
	err = yaml.Unmarshal(y, &g)
	if err != nil {
		return nil, err
	}
	return g, nil
}
