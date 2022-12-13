package util

import (
	"os"
	"strings"

	"advent-of-code-2022/12/util/graph"
	"advent-of-code-2022/12/util/queue"
	"advent-of-code-2022/helpers"
)

type Position struct {
	X int
	Y int
}

func GetNodesFromFile(path string) (map[Position]*graph.Node[Position, byte], error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var nodes = map[Position]*graph.Node[Position, byte]{}
	for row, line := range strings.Split(string(file), "\n") {
		for col := 0; col < len(line); col++ {
			nodes[Position{X: row, Y: col}] = &graph.Node[Position, byte]{Identifier: Position{X: row, Y: col}, Value: line[col], Neighbours: make(map[Position]*graph.Node[Position, byte])}

		}
	}
	return nodes, nil
}

func FindShortestPath(nodes map[Position]*graph.Node[Position, byte], root, destination *graph.Node[Position, byte]) []graph.Node[Position, byte] {
	var visited = helpers.Set[Position]{}
	visited[(*root).Identifier] = struct{}{}

	var queue = queue.New[graph.Element[Position, byte]]()
	queue.Append(graph.Element[Position, byte]{
		Node:  *root,
		Steps: []graph.Node[Position, byte]{},
	})

	for queue.Size() > 0 {
		currentItem, err := queue.Pop()
		if err != nil {
			break
		}

		for _, neighbour := range getNeighbours(currentItem.Node, nodes) {
			if _, alreadyVisited := visited[(*neighbour).Identifier]; !alreadyVisited {
				if (*neighbour).Identifier == destination.Identifier {
					return append(currentItem.Steps, currentItem.Node)
				}
				visited[(*neighbour).Identifier] = struct{}{}

				queue.Append(graph.Element[Position, byte]{
					Node:  *neighbour,
					Steps: append(currentItem.Steps, currentItem.Node),
				})
			}
		}
	}
	return []graph.Node[Position, byte]{}
}

func getNeighbours(node graph.Node[Position, byte], nodes map[Position]*graph.Node[Position, byte]) []*graph.Node[Position, byte] {
	var neighbourNodes = []*graph.Node[Position, byte]{}

	if nextNode, nodeExists := nodes[Position{X: node.Identifier.X + 1, Y: node.Identifier.Y}]; nodeExists && canClimb(node.Value, nextNode.Value) {
		neighbourNodes = append(neighbourNodes, nextNode)
	}
	if nextNode, nodeExists := nodes[Position{X: node.Identifier.X - 1, Y: node.Identifier.Y}]; nodeExists && canClimb(node.Value, nextNode.Value) {
		neighbourNodes = append(neighbourNodes, nextNode)
	}
	if nextNode, nodeExists := nodes[Position{X: node.Identifier.X, Y: node.Identifier.Y + 1}]; nodeExists && canClimb(node.Value, nextNode.Value) {
		neighbourNodes = append(neighbourNodes, nextNode)
	}
	if nextNode, nodeExists := nodes[Position{X: node.Identifier.X, Y: node.Identifier.Y - 1}]; nodeExists && canClimb(node.Value, nextNode.Value) {
		neighbourNodes = append(neighbourNodes, nextNode)
	}

	return neighbourNodes
}

func canClimb(source, destination byte) bool {
	if source > destination {
		return true
	}
	return ((destination-source) <= 1 || source == 'S' || source == 'E' || destination == 'S' || destination == 'E')
}
