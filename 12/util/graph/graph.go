package graph

import (
	"advent-of-code-2022/12/util/queue"
	"advent-of-code-2022/helpers"
)

type Node[T comparable, K any] struct {
	Identifier T
	Value      K
	Neighbours map[T]*Node[T, K]
}

func (n *Node[T, K]) AddNeighbour(node *Node[T, K]) {
	n.Neighbours[node.Identifier] = node
}

func (n *Node[T, K]) TraverseGraph() []*Node[T, K] {
	var visited = helpers.Set[T]{}
	visited[(*n).Identifier] = struct{}{}

	var queue = queue.New[Node[T, K]]()
	queue.Append(*n)

	var result = []*Node[T, K]{n}

	for queue.Size() > 0 {
		element, err := queue.Pop()
		if err != nil {
			break
		}

		for _, neighbour := range element.Neighbours {
			if _, alreadyVisited := visited[(*neighbour).Identifier]; !alreadyVisited {
				visited[(*neighbour).Identifier] = struct{}{}

				queue.Append(*neighbour)
				result = append(result, neighbour)
			}
		}
	}

	return result
}

type Element[T comparable, K any] struct {
	Node  Node[T, K]
	Steps []Node[T, K]
}

func (n *Node[T, K]) ShortestPath(destination T) []Node[T, K] {
	var visited = helpers.Set[T]{}
	visited[(*n).Identifier] = struct{}{}

	var queue = queue.New[Element[T, K]]()
	queue.Append(Element[T, K]{
		Node:  *n,
		Steps: []Node[T, K]{},
	})

	for queue.Size() > 0 {
		currentItem, err := queue.Pop()
		if err != nil {
			break
		}

		for _, neighbour := range currentItem.Node.Neighbours {
			if _, alreadyVisited := visited[(*neighbour).Identifier]; !alreadyVisited {
				if (*neighbour).Identifier == destination {
					return append(currentItem.Steps, currentItem.Node)
				}
				visited[(*neighbour).Identifier] = struct{}{}

				queue.Append(Element[T, K]{
					Node:  *neighbour,
					Steps: append(currentItem.Steps, currentItem.Node),
				})
			}
		}
	}
	return nil
}
