package main

import (
	"flag"
	"log"
	"sort"

	"advent-of-code-2022/12/util"
	"advent-of-code-2022/12/util/graph"
	"advent-of-code-2022/helpers"
)

func init() {
	flag.Parse()
}

func main() {
	nodes, err := util.GetNodesFromFile(helpers.GetInputFilePath())

	if err != nil {
		log.Fatalln(err)
	}

	var possibleStarts []*graph.Node[util.Position, byte]
	var destination *graph.Node[util.Position, byte]
	for _, n := range nodes {
		if n.Value == 'S' || n.Value == 'a' {
			n.Value = 'a'
			possibleStarts = append(possibleStarts, n)
		} else if n.Value == 'E' {
			n.Value = 'z'
			destination = n
		}
	}

	var distances = map[util.Position]int{}

	for _, start := range possibleStarts {
		steps := util.FindShortestPath(nodes, start, destination)
		if steps != nil {
			distances[start.Identifier] = len(steps)
		}
	}

	startsWithPossibleWays := make([]util.Position, 0, len(distances))

	for startPosition := range distances {
		startsWithPossibleWays = append(startsWithPossibleWays, startPosition)
	}

	sort.SliceStable(startsWithPossibleWays, func(i, j int) bool {
		return distances[startsWithPossibleWays[i]] < distances[startsWithPossibleWays[j]]
	})

	log.Printf("searched shortest path from %d starting points, shortest path starting from %v with distance: %d", len(possibleStarts), startsWithPossibleWays[0], distances[startsWithPossibleWays[0]])
}
