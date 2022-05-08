package entity

import (
	"fmt"
	"github.com/DamirJann/pretty-trie/pkg/drawing"
	"github.com/DamirJann/pretty-trie/pkg/entity"
)

type Ast interface {
	Visualize() string
	Root() Node
}

type ast struct {
	root Node
}

type Node interface {
	Token() *Token
	Label() string
	AddChildToEnd(...Node)
	AddChild(int, ...Node)
	AddChildToBegin(...Node)
	Delete(int)
	Child() []Node
	Replace(int, ...Node)
}

func NewAst(root Node) *ast {
	return &ast{
		root: root,
	}
}

func (n *node) AddChildToEnd(c ...Node) {
	n.child = append(n.child, c...)
}

func (n *node) AddChildToBegin(c ...Node) {
	n.child = append(c, n.child...)
}

func (n *node) AddChild(pos int, c ...Node) {
	n.child = append(append(n.child[:pos], c...), n.child[pos:]...)
}

func (n *node) Delete(pos int) {
	n.child = append(n.child[:pos], n.child[pos+1:]...)
}

func (n *node) Replace(pos int, new ...Node) {
	n.AddChild(pos, new...)
	n.Delete(pos + len(new))

}

func (n node) Child() []Node {
	return n.child
}

func (n node) Label() string {
	return n.label
}
func (n node) Token() *Token {
	return n.token
}

type node struct {
	label string
	token *Token
	child []Node
}

func NewNode(l string, t Token) Node {
	return &node{
		label: l,
		token: &t,
		child: nil,
	}
}

func (a *ast) Root() Node {
	return a.root
}

func (a *ast) Visualize() string {
	idSeq := new(int)
	*idSeq = 0
	res, _ := drawing.Visualize(a.traverse(a.root, idSeq))
	return res
}

func (a *ast) traverse(n Node, idSeq *int) (node []entity.Node, edge []entity.Edge) {
	label := n.Label()
	if label == "" {
		label = fmt.Sprintf("%v", n.Token().Value)
	}
	node = append(node, entity.Node{
		Id:    *idSeq,
		Label: label,
	})
	for _, child := range n.Child() {
		edge = append(edge, entity.Edge{
			From: node[0].Id,
			To:   *idSeq + 1,
		})
		*idSeq++
		newNode, newEdge := a.traverse(child, idSeq)
		node = append(node, newNode...)
		edge = append(edge, newEdge...)
	}
	return node, edge
}
