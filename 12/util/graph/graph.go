package graph

type Node[T comparable, K any] struct {
	Identifier T
	Value      K
	Neighbours map[T]*Node[T, K]
}

func (n *Node[T, K]) AddNeighbour(node *Node[T, K]) {
	n.Neighbours[node.Identifier] = node
}
