// Package dhg provides a directed hypergraph implementation.
package dhg

import "strconv"

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

//  U+f0a9 Arrow circle right
//  U+f0a8 Arrow circle left
// 󱡓 U+f1853 Circle opacity
// 󰴣 = {( , )}
// 󱗜 = {(󰴷,󰴶)}

type Element interface {
	element()
}

type Edge struct {
	Value *Datum
	Rels  *Vertices
}

func (*Edge) element() {}

type Vertex struct {
	Value *Datum
	Rels  *Edges
}

func (*Vertex) element() {}

type Rel struct {
	Head *Element
	Tail *Element
}

type Rels []Rel

type Vertices []Vertex
type Edges []Edge

type Graph struct {
	Value    *Datum
	Vertices *Vertices
	Edges    *Edges
}

func (d *Graph) Fmt() string {
	return GRAPH_SYMBOL + " " + "TODO"
}

type Datum interface {
	datum()
	Fmt() string
}

type GraphDatum struct {
	Value *Graph
}

func (*GraphDatum) datum() {}

func (d *GraphDatum) Fmt() string {
	if d.Value == nil {
		return EMPTY_GRAPH_SYMBOL
	}
	return d.Value.Fmt()
}

type NilDatum struct{}

func (*NilDatum) datum() {}
func (*NilDatum) Fmt() string {
	return NO_DATA_SYMBOL
}

type BoolDatum struct {
	Value bool
}

func (*BoolDatum) datum() {}

func (d *BoolDatum) Fmt() string {
	if d.Value {
		return BOOLEAN_SYMBOL + "true"
	}
	return BOOLEAN_SYMBOL + "false"
}

type IntegerDatum struct {
	Value int64
}

func (*IntegerDatum) datum() {}

func (d *IntegerDatum) Fmt() string {
	return INTEGER_SYMBOL + " " + strconv.FormatInt(d.Value, 10)
}

type FloatDatum struct {
	Value float64
}

func (*FloatDatum) datum() {}

func (d *FloatDatum) Fmt() string {
	return REAL_SYMBOL + " " + strconv.FormatFloat(d.Value, 'f', -1, 64)
}

type StringDatum struct {
	Value string
}

func (*StringDatum) datum() {}

func (d *StringDatum) Fmt() string {
	return STRING_SYMBOL + " " + d.Value
}

type VertexDatum struct {
	Value *Vertex
}

func (*VertexDatum) datum() {}

func (*VertexDatum) Fmt() string {
	return VERTEX_SYMBOL + " " + "TODO"
}

type EdgeDatum struct {
	Value *Edge
}

func (*EdgeDatum) datum() {}

func (*EdgeDatum) Fmt() string {
	return EDGE_SYMBOL + " " + "TODO"
}

func Fmt[T Datum](d T) string {
	return d.Fmt()
}
